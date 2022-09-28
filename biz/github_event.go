package biz

import (
	"context"

	"code.byted.org/larkim/oapi_demo/conf"

	"github.com/sirupsen/logrus"
)

func HandleReceiveGithubEvent(ctx context.Context, event *ReceiveGithubEvent) error {
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
	switch msg.MessageType {
	case "interactive":
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: receiveID,
			Content:   content,
			MsgType:   "interactive",
		}
		resp, err := SendMessage(ctx, token, createMsgRequest)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send msg")
			return err
		}
		logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	case "text":
		logrus.Infof("text %v", content)
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: receiveID,
			Content:   content,
			MsgType:   "interactive",
		}
		resp, err := SendMessage(ctx, token, createMsgRequest)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send msg")
			return err
		}
		logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	default:
		logrus.Infof("unhandled message type, msg_type: %v", msg.MessageType)
	}
	return nil
}
