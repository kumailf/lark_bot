package biz

import (
	"context"
	"fmt"

	"code.byted.org/larkim/oapi_demo/conf"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	"github.com/sirupsen/logrus"
)

func CreateDocument(title string) (string, error) {
	client := lark.NewClient(conf.Conf.AppID, conf.Conf.AppSecret)
	resp, err := client.Docx.Document.Create(
		context.Background(),
		larkdocx.NewCreateDocumentReqBuilder().Body(
			larkdocx.NewCreateDocumentReqBodyBuilder().
				Title(title).
				Build()).
			Build())

	if err != nil {
		logrus.WithError(err).Errorf("fail to create document")
		return "", err
	}
	if !resp.Success() {
		logrus.WithError(err).Errorf("fail to create document")
		return "", err
	}
	document_id := resp.Data.Document.DocumentId
	return *document_id, err
}

func GetDocumentURL(document_id string) string {
	url := fmt.Sprintf("https://aicarrier.feishu.cn/docx/%v", document_id)
	return url
}
