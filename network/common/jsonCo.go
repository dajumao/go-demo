package common

import (
	"encoding/json"
	"time"
)

type ChatInformation struct {
	Name 		string		`json:"name"`
	IP 			string		`json:"ip"`
	Information string		`json:"chatinformation"`
	Time 		time.Time 	`json:"time"`
}

func JsonCode(chatIn *ChatInformation) []byte {
	data,err := json.MarshalIndent(chatIn,""," ")
	if err != nil {
		return nil
	}
	return data
}

func JsonDecode(jsonChat []byte) ChatInformation {
	var chat  ChatInformation
	err := json.Unmarshal(jsonChat,&chat)
	if err != nil {
		return ChatInformation{}
	}
	return chat
}


