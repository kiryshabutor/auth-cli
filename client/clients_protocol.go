package main

import (
	"authenticationProject/protocols"
	"encoding/json"
	"net"
)

func sendRequest(conn net.Conn, msgType string, data any) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var msg protocols.Message
	msg.Type = msgType
	msg.Data = rawData
	return json.NewEncoder(conn).Encode(msg)

}

func readMessage(conn net.Conn) (protocols.Message, error) {
	var msg protocols.Message
	err := json.NewDecoder(conn).Decode(&msg)
	return msg, err
}

func decodeData[T any](msg protocols.Message) (T, error) {
	var out T
	err := json.Unmarshal(msg.Data, &out)
	return out, err
}
