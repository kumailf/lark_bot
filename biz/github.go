package biz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func SetProject(token string, project_id string, node_id string) error {
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
	return nil
}
