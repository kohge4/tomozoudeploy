package mainappimpl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) Login(c *gin.Context) {
	// context で 外から spotify か apple か判別
	// つまり ハンドラーを handler.Spotify handler.Apple とかで 使えるようにした方が良さげ
	u.Handler.Authenticator.SetAuthInfo(u.Handler.ClientID, u.Handler.SecretKey)
	c.JSON(200, Response{200, u.Handler.Authenticator.AuthURL(u.Handler.State)})
}

func (u *UserProfileApplicationImpl) Callback(c *gin.Context) {
	// Login が成功したら UserCase の domain.WebSeerviceAccount を更新する
	// => 更新してから RegistryUserを実行する
	accessToken, err := u.Handler.Authenticator.Token(u.Handler.State, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	u.Handler.Client = u.Handler.Authenticator.NewClient(accessToken)
	fmt.Println("accessToken")
	fmt.Println(accessToken.AccessToken)

	// ここで UseCase に切り替える
	u.UseCase.WebServiceAccount = u.Handler.ConvertWebServiceAccountImpl()

	existingUser, err := u.UseCase.CheckExistingUser()
	if err != nil {
		c.String(403, err.Error())
	}
	if existingUser != nil {
		// すでに そのサービスでログインしたことあるユーザーの場合
		c.Set("userid", existingUser.ID)
		c.Set("user_name", existingUser.UserName)

		err = u.UseCase.UpdateUser(existingUser.ID)
		if err != nil {
			c.String(403, err.Error())
		}
		u.AuthMiddleware.LoginHandler(c)
		return
	}

	user, err := u.UseCase.RegistryUser()
	if err != nil {
		c.String(403, err.Error())
	}
	c.Set("userid", user.ID)
	c.Set("user_name", user.UserName)
	u.AuthMiddleware.LoginHandler(c)
}
