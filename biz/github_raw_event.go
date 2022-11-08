package biz

import (
	"context"
	"fmt"
	"os"

	"code.byted.org/larkim/oapi_demo/conf"

	"github.com/sirupsen/logrus"
)

func HandleReceiveGithubIssueEvent(ctx context.Context, event *ReceiveGithubIssueEvent) error {
	logrus.Info("receive github raw event")
	ie := event.Event
	repo := ie.Repo.Name
	repo_fullname := ie.Repo.FullName
	eventType := ie.Action
	// Set Project
	go func() {
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
			_ = SetProject(token, project_id, node_id)
		}
	}()

	// Send Message To Lark Group
	go func() {
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
				return
			}
			token, err := GetTenantAccessToken(ctx)
			if err != nil {
				logrus.WithError(err).Errorf("failed to get tenant access token")
				return
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
				return
			}
			logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
		default:
			return
		}
	}()

	// test func conf
	go func() {
		if repo != "MyPic" {
			return
		}
		if !(conf.FuncIsWork(repo, "testfunc")) {
			return
		}
		groupName := "机器人调试"
		issueTitle := ie.Issue.Title
		createBy := ie.Issue.User.Login
		issueUrl := ie.Issue.HTMLURL
		receiveID, err := GetGroupID(groupName)
		if err != nil {
			logrus.WithError(err).Errorf("failed to get group id")
			return
		}
		token, err := GetTenantAccessToken(ctx)
		if err != nil {
			logrus.WithError(err).Errorf("failed to get tenant access token")
			return
		}
		content := fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** Issue Title: **%v\\n** Created By: **%v\\n** Link: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"test func\",\"tag\":\"plain_text\"}}}", issueTitle, createBy, issueUrl)
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: receiveID,
			Content:   content,
			MsgType:   "interactive",
		}
		resp, err := SendMessage(ctx, token, createMsgRequest)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send msg")
			return
		}
		logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	}()

	return nil
}

func HandleReceiveGithubPREvent(ctx context.Context, event *ReceiveGithubPREvent) error {
	logrus.Info("receive github raw event")
	pr := event.Event
	repo_fullname := pr.Repo.GetFullName()
	// Set Project
	go func() {
		if repo_fullname == "kumailf/MyPic" || repo_fullname == "open-mmlab/mmediting" {
			token := ""
			project_id := ""
			node_id := pr.PullRequest.GetNodeID()
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

	// Send Message To Lark Group
	go func() {
		eventType := pr.GetAction()
		var content string
		switch eventType {
		case "opened":
			prTitle := pr.PullRequest.GetTitle()
			openBy := pr.PullRequest.User.GetLogin()
			prUrl := pr.PullRequest.GetHTMLURL()
			content = fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** PR Title: **%v\\n** Opened By: **%v\\n** Link: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"New PullRequest\",\"tag\":\"plain_text\"}}}", prTitle, openBy, prUrl)
		case "closed":
			merged := pr.PullRequest.GetMerged()
			if merged {
				prTitle := pr.PullRequest.GetTitle()
				openBy := pr.PullRequest.User.GetLogin()
				mergedBy := pr.PullRequest.MergedBy.GetLogin()
				prUrl := pr.PullRequest.GetHTMLURL()
				content = fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** PR Title: **%v\\n** Opened By: **%v\\n** Merged By: **%v\\n** Link: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"PullRequest Merged\",\"tag\":\"plain_text\"}}}", prTitle, openBy, mergedBy, prUrl)
			} else {
				return
			}
		case "review_requested":
			prTitle := pr.PullRequest.GetTitle()
			prUrl := pr.PullRequest.GetHTMLURL()
			prReviewer := pr.RequestedReviewer.GetLogin()
			content = fmt.Sprintf("{\"config\":{\"wide_screen_mode\":true},\"elements\":[{\"tag\":\"div\",\"text\":{\"content\":\"** PR Title: **%v\\n** Reviewer: **%v\\n** Link: **%v\",\"tag\":\"lark_md\"}}],\"header\":{\"template\":\"green\",\"title\":{\"content\":\"PullRequest Requeste Reviewer\",\"tag\":\"plain_text\"}}}", prTitle, prReviewer, prUrl)
		default:
			return
		}
		groupName, ok := conf.GroupMap[repo_fullname]
		if !ok {
			groupName = "机器人调试"
		}
		receiveID, err := GetGroupID(groupName)
		if err != nil {
			logrus.WithError(err).Errorf("failed to get group id")
			return
		}
		token, err := GetTenantAccessToken(ctx)
		if err != nil {
			logrus.WithError(err).Errorf("failed to get tenant access token")
			return
		}
		createMsgRequest := &CreateMessageRequest{
			ReceiveID: receiveID,
			Content:   content,
			MsgType:   "interactive",
		}
		resp, err := SendMessage(ctx, token, createMsgRequest)
		if err != nil {
			logrus.WithError(err).Errorf("failed to send msg")
			return
		}
		logrus.Infof("succeed send msg, msg_id: %v", resp.MessageID)
	}()

	return nil
}
