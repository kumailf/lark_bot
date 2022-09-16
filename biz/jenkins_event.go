package biz

import (
	"context"

	"github.com/sirupsen/logrus"
)

func HandleReceiveJenkinsEvent(ctx context.Context, event *ReceiveJenkinsEvent) error {
	msg := event.Event.Message
	content := msg.Content
	receiveID, err := GetGroupID(event.GroupName)
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
	default:
		logrus.Infof("unhandled message type, msg_type: %v", msg.MessageType)
	}
	return nil
}
