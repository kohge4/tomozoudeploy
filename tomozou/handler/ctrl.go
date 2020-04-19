package handler

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/usecase"
)

type UserProfileApplicationImpl struct {
	UseCase *usecase.UserProfileApplication

	Handler        *webservice.SpotifyHandler
	AuthMiddleware *jwt.GinJWTMiddleware
}

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
		c.Set("user_name", existingUser.Name)

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
	c.Set("user_name", user.Name)
	u.AuthMiddleware.LoginHandler(c)
}

func (u *UserProfileApplicationImpl) MyProfile(c *gin.Context) {
	userID, err := getIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	me, err := u.UseCase.Me(userID)
	if err != nil {
		c.String(403, err.Error())
	}
	tag, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		return
	}
	nowplaying, err := u.UseCase.MyNowPlayingUserTrackTag(userID)
	if err != nil {
		return
	}
	trackResp := NewTrackResponse(u, *nowplaying)
	// userTrackTag 型の trackID  を用いて trackurlを作成する処理
	//nowplaying, err := u.UseCase.

	response := MyProfileResponse{
		Me:      me,
		Artists: tag,
		Tracks:  trackResp,
		//NowPlayng: nowplayng,
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) MyChatList(c *gin.Context) {
	userID, err := getIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	tag, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		return
	}
	response := MyChatListResponse{
		Artists:     tag,
		ArtistsInfo: "",
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) MyArtist(c *gin.Context) {
	userID, err := getIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	myArtists, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, myArtists)
}

func (u *UserProfileApplicationImpl) SearchUsersByArtistID(c *gin.Context) {
	artistID := c.Param("artistID")
	id, _ := strconv.Atoi(artistID)

	users, err := u.UseCase.DisplayUsersByArtistID(id)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, users)
}

func (u *UserProfileApplicationImpl) SearchUsersByArtistName(c *gin.Context) {
	name := c.Query("name")
	users, err := u.UseCase.DisplayUsersByArtistName(name)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, users)
}

func (u *UserProfileApplicationImpl) MyTrack(c *gin.Context) {
	// nowplaying の表示用
	userID, err := getIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	if userID == 0 {
		userID = 1
	}
	tags, err := u.UseCase.MyUserTrackTag(userID)
	if err != nil {
		c.String(403, err.Error())
	}
	c.JSON(200, tags)
}

func (u *UserProfileApplicationImpl) NowPlaying(c *gin.Context) {
	userID, err := getIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	// Handler から直接取ってくる方がいいかも => streaming
	trackTag, err := u.UseCase.FetchNowPlayng(userID)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, trackTag)
}

func (u *UserProfileApplicationImpl) TrackTimeLine(c *gin.Context) {
	// ページング処理的なのしたい
	trackTags, err := u.UseCase.TrackTimeLine()
	if err != nil {
		c.String(403, err.Error())
	}
	// user の画像とかを最適化することを考えると、DB モデルと domainモデルが一対一である必要はない
	response := NewTrackTimeLineResponse(u, trackTags)
	c.JSON(200, response)
}
