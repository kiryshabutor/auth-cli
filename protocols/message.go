package protocols

import "encoding/json"

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"` 
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthorizationRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthorizationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	IsAdmin bool `json:"isAdmin"`
}

type ChangePasswordRequest struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type DeleteUserRequest struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	AskPassword bool   `json:"askPassword"`
}

type GetAllUsersResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Logins []string `json:"logins"`
}

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
