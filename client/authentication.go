package main

import (
	"authenticationProject/protocols"
	"fmt"
	"net"
	"strconv"
)

const MinPasswordLength = 6

func registration(conn net.Conn) {
	login := enterLogin()
	password := enterPassword(true, "Придумайте пароль: ")
	req := protocols.RegisterRequest{Login: login, Password: password}
	if err := sendRequest(conn, "register", req); err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
		fmt.Println("Регистрация не была осуществлена")
		return
	}

	msg, err := readMessage(conn)
	if err != nil {
		fmt.Println("Ошибка при принятии ответа: ", err)
		fmt.Println("Регистрация не была осуществлена")
		return
	}
	resp, _ := decodeData[protocols.BasicResponse](msg)
	fmt.Println(resp.Message)
}

func authorization(conn net.Conn) (string, bool, bool) {
	fmt.Println("Введите данные для входа")
	login := enterLogin()
	password := enterPassword(false, "Введите пароль: ")

	req := protocols.AuthorizationRequest{Login: login, Password: password}
	if err := sendRequest(conn, "authorization", req); err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
		fmt.Println("Вход не был осуществлен")
		return "", false, false
	}

	msg, err := readMessage(conn)
	if err != nil {
		fmt.Println("Ошибка при принятии ответа: ", err)
		fmt.Println("Регистрация не была осуществлена")
		return "", false, false
	}
	resp, _ := decodeData[protocols.AuthorizationResponse](msg)
	if resp.Success {
		fmt.Println(resp.Message)
		return login, resp.Success, resp.IsAdmin
	} else {
		fmt.Println(resp.Message)
		return "", false, false
	}
}

func changePassword(conn net.Conn, login string) {
	password := enterPassword(false, "Введите текущий пароль: ")
	newPassword := enterPassword(true, "Введите новый пароль: ")

	req := protocols.ChangePasswordRequest{Login: login, Password: password, NewPassword: newPassword}
	if err := sendRequest(conn, "changePassword", req); err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
		fmt.Println("Пароль не был изменен")
		return
	}

	msg, err := readMessage(conn)
	if err != nil {
		fmt.Println("Ошибка при принятии ответа: ", err)
		fmt.Println("Пароль не был изменен")
		return
	}
	resp, _ := decodeData[protocols.BasicResponse](msg)
	if resp.Success {
		fmt.Println("Пароль изменен")
	} else {
		fmt.Println(resp.Message)
	}
}

func deleteUser(conn net.Conn, target string, askPassword bool) {
	password := ""
	if askPassword {
		password = enterPassword(false, "Введите пароль для удаления аккаунта: ")
	}
	req := protocols.DeleteUserRequest{Login: target, Password: password, AskPassword: askPassword}
	if err := sendRequest(conn, "deleteUser", req); err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
	}

	msg, err := readMessage(conn)
	if err != nil {
		fmt.Println("Ошибка при принятии ответа: ", err)
		fmt.Println("Удаление не было осуществлено")
		return
	}
	resp, _ := decodeData[protocols.BasicResponse](msg)
	fmt.Println(resp.Message)
}

func deleteUserByAdmin(conn net.Conn) {
	loginToDelete := enterLogin()
	deleteUser(conn, loginToDelete, false)
}

func getUsers(conn net.Conn) ([]string, bool) {
	if err := sendRequest(conn, "getAllUsers", struct{}{}); err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
		return nil, false
	}
	msg, err := readMessage(conn)
	if err != nil {
		fmt.Println("Ошибка при принятии ответа: ", err)
		fmt.Println("Регистрация не была осуществлена")
		return nil, false
	}
	resp, _ := decodeData[protocols.GetAllUsersResponse](msg)
	if resp.Success {
		return resp.Logins, true
	} else {
		fmt.Println(resp.Message)
		return nil, false
	}
}

func usersPrint(conn net.Conn) {
	if logins, ok := getUsers(conn); ok {
		var usersString string = "Список пользователей:\n"
		for index, value := range logins {
			usersString += strconv.Itoa(index+1) + ". " + value + "\n"
		}
		fmt.Println(usersString)
	} else {
		fmt.Println("Пользователи не будут выведены")
	}
}
