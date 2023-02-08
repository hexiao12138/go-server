package utils

import (
	"fmt"
	"go-server/model"
	"go-server/model/chat/response"
)

func GroupMessage(list []model.MessageData, messageType string) []response.ChatMessage {
	returnData := make([]response.ChatMessage, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		if messageType == "my" {
			for j = i + 1; j < len(list) && list[i].ToId == list[j].ToId; j++ {
			}
			returnData = append(returnData, response.ChatMessage{MessageId: list[i].ToId, List: list[i:j]})
		} else {
			for j = i + 1; j < len(list) && list[i].ToId == list[j].ToId; j++ {
			}
		}
		i = j
	}
	fmt.Println(returnData)
	return returnData
}
