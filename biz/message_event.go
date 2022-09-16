package biz

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func HandleReceiveMessageEvent(ctx context.Context, event *ReceiveMessageEvent) error {
	msg := event.Event.Message
	chatID := msg.ChatID
	token, err := GetTenantAccessToken(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get tenant access token")
		return err
	}
	switch msg.MessageType {
	case "text":
		if strings.Contains(msg.Content, "/solve") {
			createMsgRequest := &CreateMessageRequest{
				ReceiveID: chatID,
				Content:   "{\"text\":\"开发中 \\n\"}",
				MsgType:   "text",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
		}
		if strings.Contains(msg.Content, "吃") {
			var eatList = [...]string{
				"食堂",
				"麦当劳",
				"AI palaza",
				"便利店",
				"不吃",
			}
			rand.Seed(time.Now().UnixNano())
			eat := eatList[rand.Intn(4)]
			createMsgRequest := &CreateMessageRequest{
				ReceiveID: chatID,
				Content:   "{\"text\":\"" + eat + "\\n\"}",
				MsgType:   "text",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
		}
	case "post":
		if strings.Contains(msg.Content, "/solve") {
			createMsgRequest := &CreateMessageRequest{
				ReceiveID: chatID,
				Content:   "{\"text\":\"开发中 \\n\"}",
				MsgType:   "text",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
		}
	default:
		logrus.Infof("unhandled message type, msg_type: %v", msg.MessageType)
	}
	return nil
}
