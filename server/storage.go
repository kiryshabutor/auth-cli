package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var users = map[string]*User{}

func containsLogin(login string) bool {
	_, ok := users[login]
	return ok
}

func loadUsers() {
	file, err := os.Open(usersFile)
	if err != nil {
		if os.IsNotExist(err) {
			users = make(map[string]*User)
			return
		}
		fmt.Println("Ошибка при чтении файла: ", err)
		os.Exit(1)
	}

	defer file.Close()

	stat, _ := file.Stat()
	if stat.Size() == 0 {
		users = make(map[string]*User)
		return
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		fmt.Println("Ошибка при декодировании данных пользователей: ", err)
		os.Exit(1)
	}

}

func saveUsers() error {
	file, err := os.Create(usersFile)
	if err != nil {
		fmt.Println("Ошибка при открытии файла пользователей: ")
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(&users); err != nil {
		fmt.Println("Ошибка при кодировании данных пользователей")
		return err
	}
	return nil
}

func addUserToDB(login string, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	users[login] = &User{Password: hashedPassword}
	if errSave := saveUsers(); errSave != nil {
		return errSave
	}
	return nil
}

func deleteUser(login string) {
	delete(users, login)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func comparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
