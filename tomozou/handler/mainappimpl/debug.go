package mainappimpl

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) Debug(c *gin.Context) {
	a := u.Handler.DebugItem(1)
	c.JSON(200, a)
}

// デバッグ
func (u *UserProfileApplicationImpl) DebugTrackTag(c *gin.Context) {
	nowplaying, err := u.UseCase.MyNowPlayingUserTrackTag(5)
	if err != nil {
		return
	}
	trackResp := NewTrackResponse(u, *nowplaying)
	c.JSON(200, trackResp)
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
