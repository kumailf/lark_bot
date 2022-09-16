package conf

type Config struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	EncryptKey string `json:"encrypt_key"`
}

var Conf *Config

// replace conf when in use
const (
	appID      = "cli_a234d72466f9d00b"
	secret     = "LFZzsDBKtVeKQibIzQ4eviR2SkhWd4e8"
	encryptKey = "bebKa2lcEIrcQpCBhEaOyezSwuolyaDf"
)

func init() {
	Conf = &Config{
		AppID:      appID,
		AppSecret:  secret,
		EncryptKey: encryptKey,
	}
}
