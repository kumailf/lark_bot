package biz

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"code.byted.org/larkim/oapi_demo/conf"
	"github.com/sirupsen/logrus"
)

// APIPath
var (
	TenantAccessTokenURL = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
)

// GetTenantAccessToken get tenant access token for app
// Refer to: https://open.feishu.cn/document/ukTMukTMukTM/ukDNz4SO0MjL5QzM/auth-v3/auth/tenant_access_token_internal
func GetTenantAccessToken(ctx context.Context) (string, error) {
	cli := &http.Client{}
	reqBody := TenantAccessTokenRequest{
		APPID:     conf.Conf.AppID,
		APPSecret: conf.Conf.AppSecret,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", TenantAccessTokenURL, strings.NewReader(string(reqBytes)))
	if err != nil {
		return "", err
	}
	resp, err := cli.Do(req)
	if err != nil {
		logrus.WithError(err).Errorf("failed to get token")
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	tokenResp := &TenantAccessTokenResponse{}
	err = json.Unmarshal(body, tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.TenantAccessToken, nil
}
