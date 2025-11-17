package main

type User struct {
	Password string
}

const (
	usersFile = "users.json"

	DefaultAdminLogin = "admin"
	DefaultAdminHash  = "$2a$10$Cftvx4a248wcUgIIR3Dgq.5NU87VJTA/InNPl2TU/aqV59iwR3LRS"
)
