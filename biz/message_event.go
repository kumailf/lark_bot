package biz

import (
	"context"
	"strings"

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
				Content:   "{\"text\":\"问题已解决，辛苦了 \\n\"}",
				MsgType:   "text",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
			chatInfo, err := GetChatInfo(ctx, token, chatID)
			if err != nil {
				logrus.WithError(err).Errorf("failed to get chat")
				return err
			}
			updateChatRequest := &UpdateChatRequest{
				Name: "（已解决）" + chatInfo.Name,
			}
			if chatInfo.I18nNames != nil {
				i18nName := &I18nNames{
					ZhCn: "（已解决）" + chatInfo.I18nNames.ZhCn,
					EnUs: "(solved)" + chatInfo.I18nNames.EnUs,
					JaJp: "（解決済み）" + chatInfo.I18nNames.JaJp,
				}
				updateChatRequest.I18nNames = i18nName
			}
			logrus.Infof("succeed get chat info")
			_, err = UpdateChat(ctx, token, chatID, updateChatRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to update chat")
				return err
			}
			logrus.Infof("succeed update chat")
		}
	case "post":
		if strings.Contains(msg.Content, "/solve") {
			createMsgRequest := &CreateMessageRequest{
				ReceiveID: chatID,
				Content:   "{\"text\":\"问题已解决，辛苦了 \\n\"}",
				MsgType:   "text",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
			chatInfo, err := GetChatInfo(ctx, token, chatID)
			if err != nil {
				logrus.WithError(err).Errorf("failed to get chat")
				return err
			}
			updateChatRequest := &UpdateChatRequest{
				Name: "（已解决）" + chatInfo.Name,
			}
			if chatInfo.I18nNames != nil {
				i18nName := &I18nNames{
					ZhCn: "（已解决）" + chatInfo.I18nNames.ZhCn,
					EnUs: "(solved)" + chatInfo.I18nNames.EnUs,
					JaJp: "（解決済み）" + chatInfo.I18nNames.JaJp,
				}
				updateChatRequest.I18nNames = i18nName
			}
			logrus.Infof("succeed get chat info")
			_, err = UpdateChat(ctx, token, chatID, updateChatRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to update chat")
				return err
			}
			logrus.Infof("succeed update chat")
		} else {
			logrus.Infof("receive msg, msg: %v", msg.Content)
		}
	default:
		logrus.Infof("unhandled message type, msg_type: %v", msg.MessageType)
	}
	return nil
}



