package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	loginRegexp    = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	passwordRegexp = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*_]+$`)
)

func isValid(value string, isPassword bool) bool {
	if isPassword {
		return passwordRegexp.MatchString(value)
	}
	return loginRegexp.MatchString(value)
}

func enterLogin() (login string) {
	for {
		login = readInput("Введите логин: ")
		if isValid(login, false) {
			return
		}
		fmt.Println("Логин может состоять только из латинских букв и цифр. Попробуйте снова")
	}
}

func enterPassword(isRegistration bool, message string) (password string) {
	for {
		password = readInput(message)
		if len(password) < MinPasswordLength && isRegistration {
			fmt.Println("Длина пароля должна быть не меньше 6 символов. Попробуйте снова")
		} else if isValid(password, true) {
			return
		}
		fmt.Println("Пароль может состоять только из латинских букв, цифр и спец символов. Попробуйте снова")
	}
}

func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return ""
	}
	return strings.TrimSpace(text)
}

func readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	text = strings.TrimSpace(text)
	return strconv.Atoi(text)
}
