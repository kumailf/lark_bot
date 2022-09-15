package biz

import (
	"context"

	"github.com/sirupsen/logrus"
)

func HandleReceiveJenkinsEvent(ctx context.Context, event *ReceiveJenkinsEvent) error {
	msg := event.Event.Message
	content := msg.Content
	token, err := GetTenantAccessToken(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get tenant access token")
		return err
	}
	switch msg.MessageType {
	case "interactive":
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: "oc_ab9c97a52c8baebc1110c6bb3a449bdc",
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
