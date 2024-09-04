package conf

type Config struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	EncryptKey string `json:"encrypt_key"`
}

var Conf *Config

// replace conf when in use
const (
	appID      = "**"
	secret     = "**"
	encryptKey = "**"
)

func init() {
	Conf = &Config{
		AppID:      appID,
		AppSecret:  secret,
		EncryptKey: encryptKey,
	}
}
