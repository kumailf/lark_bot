package biz

import (
	"context"

	"code.byted.org/larkim/oapi_demo/conf"
	"github.com/sirupsen/logrus"
)

func HandleReceiveSendMessageEvent(ctx context.Context, event *ReceiveSendMessageEvent) error {
	msg := event.Event.Message
	content := msg.Content
	groupName, ok := conf.GroupMap[event.GroupName]
	if !ok {
		groupName = "机器人调试"
	}
	receiveID, err := GetGroupID(groupName)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get group id")
	}
	token, err := GetTenantAccessToken(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get tenant access token")
		return err
	}

	createMsgRequest := &CreateMessageRequest{
		ReceiveID: receiveID,
		Content:   content,
		MsgType:   msg.MessageType,
	}
	resp, err := SendMessage(ctx, token, createMsgRequest)
	if err != nil {
		logrus.WithError(err).Errorf("failed to send msg")
		return err
	}
	logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	return nil
}
