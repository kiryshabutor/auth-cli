package main

import (
	"authenticationProject/protocols"
)

func handleRegister(req protocols.RegisterRequest) protocols.BasicResponse {
	var resp protocols.BasicResponse
	if req.Login == "admin" {
		resp.Success = false
		resp.Message = "Данное имя пользователя зарезервировано"
	} else if containsLogin(req.Login) {
		resp.Success = false
		resp.Message = "Данное имя пользователя занято"
	} else if err := addUserToDB(req.Login, req.Password); err != nil {
		resp.Success = false
		resp.Message = "Ошибка при добавлении в Базу Данных"
	} else {
		resp.Success = true
		resp.Message = "Пользователь успешно зарегистрирован"
	}
	return resp
}

func handleAuthorization(req protocols.AuthorizationRequest) protocols.AuthorizationResponse {
	var resp protocols.AuthorizationResponse
	if containsLogin(req.Login) {
		if !comparePassword(req.Password, users[req.Login].Password) {
			resp.Success = false
			resp.Message = "Неверный пароль"
		} else {
			resp.Success = true
			resp.Message = "Вы успешно вошли"
			if req.Login == DefaultAdminLogin {
				resp.IsAdmin = true
			} else {
				resp.IsAdmin = false
			}
		}
	} else {
		resp.Success = false
		resp.Message = "Пользователь отсутствует"
	}
	return resp
}

func handleChangePassword(req protocols.ChangePasswordRequest) protocols.BasicResponse {
	var resp protocols.BasicResponse
	if !comparePassword(req.Password, users[req.Login].Password) {
		resp.Success = false
		resp.Message = "Неверный пароль"
	} else {
		newHashPassword, err := hashPassword(req.NewPassword)
		if err != nil {
			resp.Success = false
			resp.Message = "Ошибка на сервере"
		} else {
			users[req.Login].Password = newHashPassword
			resp.Success = true
			resp.Message = "Вы успешно сменили пароль"
		}
	}

	return resp
}

func handleGetAllUsers() protocols.GetAllUsersResponse {
	var resp protocols.GetAllUsersResponse
	logins := make([]string, 0, len(users))
	for loginValue := range users {
		logins = append(logins, loginValue)
	}

	resp.Success = true
	resp.Message = "Список пользователей успешно получен"
	resp.Logins = logins
	return resp
}

func handleDeleteUser(req protocols.DeleteUserRequest) protocols.BasicResponse {
	var resp protocols.BasicResponse
	_, ok := users[req.Login]
	if !ok {
		resp.Success = false
		resp.Message = "Данного пользователя не существует"
	} else if req.Login == DefaultAdminLogin {
		resp.Success = false
		resp.Message = "Данного пользователя нельзя удалить"
	} else if req.AskPassword && !comparePassword(req.Password, users[req.Login].Password) {
		resp.Success = false
		resp.Message = "Неверный пароль"
	} else {
		deleteUser(req.Login)
		resp.Success = true
		resp.Message = "Пользователь успешно удален"
	}
	return resp
}
