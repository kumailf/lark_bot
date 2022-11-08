package biz

import (
	"context"
	"fmt"
	"math/rand"

	"regexp"
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
		content := ""
		if strings.Contains(msg.Content, "吃") {
			var eatList = [...]string{
				"食堂",
				"麦当劳",
				"AI plaza",
				"便利店",
				"外卖",
				"减肥,不吃",
			}
			rand.Seed(time.Now().UnixNano())
			eat := eatList[rand.Intn(4)]
			content = "{\"text\":\"" + eat + "\"}"
		} else if strings.Contains(msg.Content, "jenkins") {
			re := regexp.MustCompile(`^{"text":"@_user_1 jenkins (.*?)"}$`)
			match := re.FindStringSubmatch(msg.Content)
			if len(match) != 0 {
				data := match[1]
				SendToMQ(data, "jenkins")
				logrus.Infof("send to mq susscss")
				content = "{\"text\":\"ok\"}"
			} else {
				content = "{\"text\":\"输入格式错误\"}"
			}
		} else if strings.Contains(msg.Content, "/help") {
			url := "https://aicarrier.feishu.cn/sheets/shtcndIKakmnLQ01dgY5e9At4Dh?table=tblGSRx6dsDMUuPr&view=vewwZW32nQ"
			f1 := "推送pr/issue/review等信息至指定群"
			f2 := "@机器人触发一次接口请求(需定制)"
			f3 := "响应群内消息/快速查询信息(需定制)"
			content := fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** 提需求 or 意见反馈: **%v\\n** 1: **%v\\n**2: **%v\\n** 3: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"New Issue\",\"tag\":\"plain_text\"}}}", url, f1, f2, f3)
			createMsgRequest := &CreateMessageRequest{
				ReceiveID: chatID,
				Content:   content,
				MsgType:   "interactive",
			}
			resp, err := SendMessage(ctx, token, createMsgRequest)
			if err != nil {
				logrus.WithError(err).Errorf("failed to send msg")
				return err
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
			return nil
		} else {
			return nil
		}
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: chatID,
			Content:   content,
			MsgType:   "text",
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
