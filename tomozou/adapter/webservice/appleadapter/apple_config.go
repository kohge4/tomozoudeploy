package appleadapter

import (
	"fmt"
	"io/ioutil"
	"os"

	applemusic "github.com/kohge4/go-apple-music-sdk"
	"github.com/minchao/go-apple-music/token"
	"github.com/rs/zerolog/log"
)

func NewWebServiceConfig() *WebServiceConfig {
	// デプロイ時ここが原因でエラー => ローカルでしか使えない
	// 本番環境の時どうするか
	secret, err := ioutil.ReadFile("./settings/AuthKey_BQC7LLSNCB.p8")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	gen := token.Generator{
		KeyId:  "BQC7LLSNCB",
		TeamId: "4QLW4H766S",
		//TTL:    ttl,
		Secret: secret,
	}
	appleToken, err := gen.Generate()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		//os.Exit(1)
	}
	log.Info().Interface("APPLE_TOKEN", appleToken).Msg("appleadapter/apple_token.go")
	apiToken := WebAPIToken{
		AccessToken: appleToken,
	}
	appleConfig := &WebServiceConfig{
		Token: &apiToken,
	}
	return appleConfig
}

func (h *AppleHandler) UpdateWebServiceConfig() {
	config := NewWebServiceConfig()
	h.Config = config
	tp := applemusic.Transport{Token: config.Token.AccessToken}
	client := applemusic.NewClient(tp.Client())
	h.Client = client
}
