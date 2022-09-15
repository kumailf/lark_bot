package conf

type Config struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	EncryptKey string `json:"encrypt_key"`
}

var Conf *Config

// replace conf when in use
const (
	appID      = "cli_a3aa9f78617c100c"
	secret     = "B72Y4SPSWf2GCoxH0hz7s2hYM5duQ7vJ"
	encryptKey = "bebKa2lcEIrcQpCBhEaOyezSwuolyaDf"
)

func init() {
	Conf = &Config{
		AppID:      appID,
		AppSecret:  secret,
		EncryptKey: encryptKey,
	}
}
