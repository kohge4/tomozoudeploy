package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	appleLogin "github.com/kohge4/appleauth"
)

const TeamID = "4QLW4H766S"
const ClientID = "ongakulogin"
const RedirectURL = "https://ongakuconnection.com/apple/callback"
const KeyID = "S4FMLKM72S"

var a *appleLogin.AppleConfig

func (u *UserProfileApplicationImpl) LoginApple(c *gin.Context) {
	a = appleLogin.InitAppleConfig(TeamID, ClientID, RedirectURL, KeyID)

	//import cert
	err := a.LoadP8CertByFile("./settings/AuthKey_S4FMLKM72S.p8") //path to cert file
	//or you can load cert from a string
	//st := "MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQg9QveaHJ+/9kW8BPV4bFWxP2VztLhsMXYoVcYZ1OiIXOgCgYIKoZIzj0DAQehRANCAAShEOAmXBUIpVWtZEVcxEV+KNiiyckCSeww8OwlogrcFgpkDzegfpzOexMR+ftO+OU6VhIxammJD/K7aMnrDgaw"
	//err := a.LoadP8CertByByte([]byte(st))
	if err != nil {
		panic(err)
	}
	fmt.Println("certloadok")
	fmt.Println(a.AESCert)
	callbackURL := a.CreateCallbackURL("statehere", "")
	fmt.Println(callbackURL)
	c.JSON(200, gin.H{"url": callbackURL})
	//"https://appleid.apple.com/auth/authorize?" + u.Encode()
	// apple の画面, ここにアクセスして認証後に 設定した callbackURL が呼び出される
}

func (u *UserProfileApplicationImpl) CallbackApple(c *gin.Context) {
	// apple に設定した url から呼び出される
	// spa の場合 appleで設定した spa の url から axios　で呼ぶ出す

	// http://localhost:8000/apple/callback?code=undefined&state=undefined 多分ここらやる感じな気がする
	// apple サーバから送られるクエリを取得(response_mode)  その値をしようして ここの code で　用いて token を取得
	fmt.Println("checkCert")
	fmt.Println(a.AESCert)
	code := c.Query("code")
	token, err := a.GetAppleToken(code, 3600)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("accesstoken")
	fmt.Println(token.AccessToken)
	c.JSON(200, token.AccessToken)
}
