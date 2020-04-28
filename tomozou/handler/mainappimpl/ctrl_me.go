package mainappimpl

import (
	"tomozou/handler/chatappimpl"
	"tomozou/handler/common"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) MyProfile(c *gin.Context) {
	userID, err := common.GetIDFromContext(c)
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
	userID, err := common.GetIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	tag, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		return
	}
	response := chatappimpl.MyChatListResponse{
		Artists:     tag,
		ArtistsInfo: "",
	}
	c.JSON(200, response)
}

func (u *UserProfileApplicationImpl) MyArtist(c *gin.Context) {
	userID, err := common.GetIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	myArtists, err := u.UseCase.MyUserArtistTag(userID)
	if err != nil {
		c.JSON(403, err.Error())
	}
	c.JSON(200, myArtists)
}

func (u *UserProfileApplicationImpl) MyTrack(c *gin.Context) {
	// nowplaying の表示用
	userID, err := common.GetIDFromContext(c)
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
