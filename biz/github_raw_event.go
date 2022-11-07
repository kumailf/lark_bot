package biz

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"code.byted.org/larkim/oapi_demo/conf"

	"github.com/sirupsen/logrus"
)

func HandleReceiveGithubIssueEvent(ctx context.Context, event *ReceiveGithubIssueEvent) error {
	logrus.Info("receive github raw event")
	ie := event.Event
	repo_fullname := ie.Repo.FullName
	eventType := ie.Action

	// Set Project
	if repo_fullname == "kumailf/MyPic" || repo_fullname == "open-mmlab/mmediting" {
		token := ""
		project_id := ""
		node_id := ie.Issue.NodeID
		if repo_fullname == "kumailf/MyPic" {
			token = os.Getenv("token_MyPic")
			project_id = os.Getenv("projectId_MyPic")
		}
		if repo_fullname == "open-mmlab/mmediting" {
			token = os.Getenv("token_mmediting")
			project_id = os.Getenv("projectId_mmediting")
		}
		url := "https://api.github.com/graphql"
		method := "POST"
		data := fmt.Sprintf(`{"query":"mutation {addProjectV2ItemById(input: {projectId: \"%v\" contentId: \"%v\"}) {item {id}}}"}`, project_id, node_id)
		payload := strings.NewReader(data)
		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			logrus.WithError(err).Errorf("failed to request")
			return err
		}
		auth := fmt.Sprintf("bearer %v", token)
		req.Header.Add("Authorization", auth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if err != nil {
			logrus.WithError(err).Errorf("failed to request")
			return err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logrus.WithError(err).Errorf("fail")
			return err
		}
		logrus.Infof(string(body))
	}
	// Send Message To Lark Group

	switch eventType {
	case "opened":
		issueTitle := ie.Issue.Title
		createBy := ie.Issue.User.Login
		issueUrl := ie.Issue.HTMLURL
		groupName, ok := conf.GroupMap[repo_fullname]
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
		}
		content := fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** Issue Title: **%v\\n** Created By: **%v\\n** Link: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"New Issue\",\"tag\":\"plain_text\"}}}", issueTitle, createBy, issueUrl)
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: receiveID,
			Content:   content,
			MsgType:   "interactive",
		}
		resp, err := SendMessage(ctx, token, createMsgRequest)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send msg")
		}
		logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	}

	return nil
}

func HandleReceiveGithubPREvent(ctx context.Context, event *ReceiveGithubPREvent) error {
	logrus.Info("receive github raw event")
	repo_fullname := event.Event.Repo.GetFullName()

	// Set Project
	go func() {
		if repo_fullname == "kumailf/MyPic" || repo_fullname == "open-mmlab/mmediting" {
			token := ""
			project_id := ""
			node_id := event.Event.PullRequest.GetNodeID()
			if repo_fullname == "kumailf/MyPic" {
				token = os.Getenv("token_MyPic")
				project_id = os.Getenv("projectId_MyPic")
			}
			if repo_fullname == "open-mmlab/mmediting" {
				token = os.Getenv("token_mmediting")
				project_id = os.Getenv("projectId_mmediting")
			}
			_ = SetProject(token, project_id, node_id)
		}
	}()

	return nil
}
