package biz

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"code.byted.org/larkim/oapi_demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReceiveEvent(c *gin.Context) {

	// get token and reply challenge
	var req = &ReceiveEventEncrypt{}
	var decryptStr string
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.WithError(err).Errorf("failed to read request")
		return
	}
	if strings.Contains(string(bytes), "encrypt") {
		err = json.Unmarshal(bytes, req)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
		}
		decryptStr, err = Decrypt(req.Encrypt, conf.Conf.EncryptKey)
		if err != nil {
			logrus.WithError(err).Errorf("decrypt error")
			return
		}
		logrus.Infof("decrypt event: %v", decryptStr)
		decryptToken := &DecryptToken{}
		err = json.Unmarshal([]byte(decryptStr), decryptToken)
		if err != nil {
			logrus.Errorf("Unmarshal failed again")
			return
		}
		if decryptToken.Challenge != "" {
			c.JSON(200, gin.H{
				"challenge": decryptToken.Challenge,
			})
			return
		}
	} else {
		logrus.Infof("without encrypt")
		decryptStr = string(bytes)
		logrus.Infof("decrypt event: %v", decryptStr)
	}

	event := &Event{}
	err = json.Unmarshal([]byte(decryptStr), event)
	if err != nil {
		logrus.Errorf("Unmarshal failed")
		return
	}
	logrus.Infof("receive event, event: %v", event)
	ctx := context.Background()
	eventType := event.Header.EventType
	logrus.Infof("header: %v", event.Header)
	logrus.Infof("eventType: %v", eventType)
	switch eventType {
	case "im.message.receive_v1":
		receiveMsgEvent := &ReceiveMessageEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveMsgEvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed, maybe Challenge")
			return
		}
		go func() {
			err = HandleReceiveMessageEvent(ctx, receiveMsgEvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	case "jenkins":
		receiveJenkinsEvent := &ReceiveJenkinsEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveJenkinsEvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
			return
		}
		go func() {
			err = HandleReceiveJenkinsEvent(ctx, receiveJenkinsEvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	case "MQ":
		receiveMQEvent := &ReceiveMQEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveMQEvent)
		if err != nil {
			logrus.Errorf("Unmarshal receiveMQEvent failed")
			return
		}

		data, err := HandleReceiveMQEvent(ctx, receiveMQEvent)
		if err != nil {
			logrus.WithError(err).Errorf("handle receive message event failed")
		}

		logrus.Infof("data out = %v", data)
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    data,
		})
	case "github":
		receiveGithubEvent := &ReceiveGithubEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveGithubEvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
			return
		}
		go func() {
			err = HandleReceiveGithubEvent(ctx, receiveGithubEvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	case "sendMessage":
		receiveSendMessageEvent := &ReceiveSendMessageEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveSendMessageEvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
			return
		}
		go func() {
			err = HandleReceiveSendMessageEvent(ctx, receiveSendMessageEvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	case "issue":
		receiveGithubIssueEvent := &ReceiveGithubIssueEvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveGithubIssueEvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
			return
		}
		go func() {
			err = HandleReceiveGithubIssueEvent(ctx, receiveGithubIssueEvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	case "pullrequest":
		receiveGithubPREvent := &ReceiveGithubPREvent{}
		err = json.Unmarshal([]byte(decryptStr), receiveGithubPREvent)
		if err != nil {
			logrus.Errorf("Unmarshal failed")
			return
		}
		go func() {
			err = HandleReceiveGithubPREvent(ctx, receiveGithubPREvent)
			if err != nil {
				logrus.WithError(err).Errorf("handle receive message event failed")
			}
		}()
		c.JSON(200, gin.H{
			"message": "ok",
		})
	default:
		logrus.Info("unhandled event")
		c.JSON(200, gin.H{
			"message": "error",
		})
	}

}

func Decrypt(encrypt string, key string) (string, error) {
	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
}
