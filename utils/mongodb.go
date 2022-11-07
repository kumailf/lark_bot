package utils

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:password@139.196.52.65:27017"

func MGDBFindOne(db string, collection string, key string, value string) map[string]interface{} {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithError(err).Errorf("failed to connect mongodb")
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			logrus.WithError(err).Errorf("failed to disconnect")
		}
	}()
	coll := client.Database(db).Collection(collection)
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{key, value}}).Decode(&result)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get key")
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		logrus.WithError(err).Errorf("failed to marshal")
	}
	result_map := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &result_map)
	if err != nil {
		logrus.WithError(err).Errorf("failed to unmarshal")
	}
	return result_map
}

func MGDBUpdateOne(db string, collection string, key string, value interface{}, update_key string, update_value interface{}) bool {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.WithError(err).Errorf("failed to connect mongodb")
		return false
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			logrus.WithError(err).Errorf("failed to disconnect")
		}
	}()
	coll := client.Database(db).Collection(collection)
	filter := bson.D{{key, value}}
	update := bson.D{{"$set", bson.D{{update_key, update_value}}}}
	_, err = coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logrus.WithError(err).Errorf("failed to update")
		return false
	}
	return true
}

// insert a document use mongosh

// mongosh mongodb://root:password@139.196.52.65:27017

// use lark_bot
// db.MyPic.insertOne(
//   {
// 		repo_fullname: "kumailf/MyPic",
// 		lark_group: "机器人调试",
// 		new_issue_notice: "off",
// 		new_pr_notice: "off",
// 		pr_merged_notice: "off",
//   }
// )

