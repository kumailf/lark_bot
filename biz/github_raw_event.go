package biz

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func HandleReceiveGithubIssueEvent(ctx context.Context, event *ReceiveGithubIssueEvent) error {
	logrus.Info("receive github raw event")

	repo_fullname := event.Event.Repo.FullName

	if repo_fullname == "kumailf/MyPic" || repo_fullname == "open-mmlab/mmediting" {
		token := ""
		project_id := ""
		node_id := event.Event.Issue.NodeID
		if repo_fullname == "kumailf/MyPic" {
			token = os.Getenv("token_MyPic")
			project_id = os.Getenv("projectId_MyPic")
		}
		if repo_fullname == "open-mmlab/mmediting" {
			token = os.Getenv("token_mmediting")
			project_id = os.Getenv("projectId_mmediting")
		}
		err := SetProject(token, project_id, node_id)
		url := "https://api.github.com/graphql"
		method := "POST"
		data := fmt.Sprintf(`{"query":"mutation {addProjectV2ItemById(input: {projectId: \"%v\" contentId: \"%v\"}) {item {id}}}"}`, project_id, node_id)
		payload := strings.NewReader(data)
		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		auth := fmt.Sprintf("bearer %v", token)
		req.Header.Add("Authorization", auth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(string(body))
	}
	return nil
}

// func HandleReceivePREvent(ctx context.Context, event *ReceiveGithubPREvent) error {
// 	logrus.Info("receive github raw event")

// 	repo_fullname := event.Event
// 	go func () {}
// 	if repo_fullname == "kumailf/MyPic" || repo_fullname == "open-mmlab/mmediting" {
// 		token := ""
// 		project_id := ""
// 		node_id := event.Event.Issue.NodeID
// 		if repo_fullname == "kumailf/MyPic" {
// 			token = os.Getenv("token_MyPic")
// 			project_id = os.Getenv("projectId_MyPic")
// 		}
// 		if repo_fullname == "open-mmlab/mmediting" {
// 			token = os.Getenv("token_mmediting")
// 			project_id = os.Getenv("projectId_mmediting")
// 		}
// 		err := SetProject(token, project_id, node_id)
// 		url := "https://api.github.com/graphql"
// 		method := "POST"
// 		data := fmt.Sprintf(`{"query":"mutation {addProjectV2ItemById(input: {projectId: \"%v\" contentId: \"%v\"}) {item {id}}}"}`, project_id, node_id)
// 		payload := strings.NewReader(data)
// 		client := &http.Client{}
// 		req, err := http.NewRequest(method, url, payload)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}
// 		auth := fmt.Sprintf("bearer %v", token)
// 		req.Header.Add("Authorization", auth)
// 		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}
// 		defer res.Body.Close()

// 		body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		}
// 		fmt.Println(string(body))
// 	}
// 	return nil
// }
