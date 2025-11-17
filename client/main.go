package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		return
	}
	defer conn.Close()
	baseMenu(conn)
}

func baseMenu(conn net.Conn) {
Exit:
	for {
		fmt.Println("Меню:\n1 - Зарегистрироваться\n2 - Войти\n0 - Выход")
		if choose, err := readInt("Выберите номер пункта меню: "); err == nil {
			switch choose {
			case 1:
				registration(conn)
			case 2:
				if login, ok, isAdmin := authorization(conn); ok {
					if isAdmin {
						adminPanel(conn)
					} else {
						userPanel(conn, login)
					}
				}
			case 0:
				fmt.Println("Осуществляем выход из программы...")
				break Exit
			default:
				fmt.Println("\nОтсутствует такой вариант ввода. Попробуйте еще раз!")
			}
		} else {
			fmt.Println("Ошибка ввода! Введите номер меню")
		}
	}
}

func adminPanel(conn net.Conn) {
Exit:
	for {
		fmt.Println("|ADMIN PANEL|")
		fmt.Println("Меню:\n1 - вывести всех пользователей\n2 - удалить пользователя\n0 -  выйти из панели администратора")
		if adminChoose, err := readInt("Выберите номер пункта меню: "); err == nil {
			switch adminChoose {
			case 1:
				usersPrint(conn)
			case 2:
				deleteUserByAdmin(conn)
			case 0:
				fmt.Println("Осуществляем выход из учетной записи...")
				break Exit
			default:
				fmt.Println("\nОтсутствует такой вариант ввода. Попробуйте еще раз!")
			}
		} else {
			fmt.Println("Ошибка ввода! Введите номер меню")
		}
	}
}

func userPanel(conn net.Conn, login string) {
Exit:
	for {
		fmt.Println("\nДоступные функции:\n1 - изменить пароль\n2 - удалить пользователя\n0 - выйти из аккаунта")
		if systemChoose, err := readInt("Выберите номер пункта меню: "); err == nil {
			switch systemChoose {
			case 1:
				changePassword(conn, login)
			case 2:
				deleteUser(conn, login, true)
				break Exit
			case 0:
				fmt.Println("Осуществляем выход из учетной записи...")
				break Exit
			default:
				fmt.Println("\nОтсутствует такой вариант ввода. Попробуйте еще раз!")
			}
		} else {
			fmt.Println("Ошибка ввода! Введите номер меню")
		}
	}
}
