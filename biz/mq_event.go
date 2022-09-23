package biz

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()
var rdb *redis.Client

func HandleReceiveMQEvent(ctx context.Context, event *ReceiveMQEvent) (string, error) {
	ops := event.Event.Ops
	mqType := event.Event.MQType
	data := event.Event.Data
	switch ops {
	case "send":
		SendToMQ(data, mqType)
	case "receive":
		return ReceiveFromMQ(mqType), nil
	default:
		logrus.Infof("error ops")
	}

	return "", nil
}

func SendToMQ(data string, mqType string) {
	err := initClient()
	if err != nil {
		logrus.WithError(err).Errorf("redis client error")
		return
	}
	rdb.LPush(ctx, mqType, data)
}

func ReceiveFromMQ(mqType string) string {
	err := initClient()
	if err != nil {
		logrus.WithError(err).Errorf("redis client error")
		return ""
	}
	ret, err := rdb.LPop(ctx, mqType).Result()
	if err != nil {
		logrus.WithError(err).Errorf("get error")
		return ""
	}
	return ret
}

func initClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
