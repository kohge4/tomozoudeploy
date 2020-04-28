package mainappimpl

import (
	"tomozou/handler/common"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) NowPlaying(c *gin.Context) {
	userID, err := common.GetIDFromContext(c)
	if err != nil {
		c.String(403, err.Error())
	}
	// Handler から直接取ってくる方がいいかも => streaming
	nowplayingTrackTag, err := u.UseCase.CallNowPlayng(userID)
	if err != nil {
		c.JSON(403, err.Error())
	}
	trackResp := NewTrackResponse(u, *nowplayingTrackTag)
	c.JSON(200, trackResp)
}
