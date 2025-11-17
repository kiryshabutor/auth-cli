package main

import (
	"authenticationProject/protocols"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func runServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Сервер слушает на порту 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)
	for {
		var msg protocols.Message
		if err := dec.Decode(&msg); err != nil {
			fmt.Println("Ошибка разбора обёртки Message:", err)
			return
		}

		switch msg.Type {
		case "register":
			var (
				recv protocols.BasicResponse
				req  protocols.RegisterRequest
			)
			_ = json.Unmarshal(msg.Data, &req)
			recv = handleRegister(req)
			rawData, _ := json.Marshal(recv)
			msg := protocols.Message{
				Type: "register",
				Data: rawData,
			}
			json.NewEncoder(conn).Encode(msg)
		case "authorization":
			var (
				recv protocols.AuthorizationResponse
				req  protocols.AuthorizationRequest
			)
			_ = json.Unmarshal(msg.Data, &req)
			recv = handleAuthorization(req)
			rawData, _ := json.Marshal(recv)
			msg := protocols.Message{
				Type: "authorization",
				Data: rawData,
			}
			json.NewEncoder(conn).Encode(msg)
		case "changePassword":
			var (
				recv protocols.BasicResponse
				req  protocols.ChangePasswordRequest
			)
			_ = json.Unmarshal(msg.Data, &req)
			recv = handleChangePassword(req)
			rawData, _ := json.Marshal(recv)
			msg := protocols.Message{
				Type: "changePassword",
				Data: rawData,
			}
			json.NewEncoder(conn).Encode(msg)
		case "getAllUsers":
			var (
				recv protocols.GetAllUsersResponse
			)
			recv = handleGetAllUsers()
			rawData, _ := json.Marshal(recv)
			msg := protocols.Message{
				Type: "getAllUsers",
				Data: rawData,
			}
			json.NewEncoder(conn).Encode(msg)
		case "deleteUser":
			var (
				recv protocols.BasicResponse
				req  protocols.DeleteUserRequest
			)
			_ = json.Unmarshal(msg.Data, &req)
			recv = handleDeleteUser(req)
			rawData, _ := json.Marshal(recv)
			msg := protocols.Message{
				Type: "deleteUser",
				Data: rawData,
			}
			json.NewEncoder(conn).Encode(msg)
		default:
			enc.Encode(protocols.BasicResponse{
				Success: false,
				Message: fmt.Sprintf("unknown message type %q", msg.Type),
			})
		}
	}
}
