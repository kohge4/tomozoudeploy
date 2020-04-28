package mainappimpl

import (
	"tomozou/handler/common"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) TrackTimeLine(c *gin.Context) {
	// domain model に画像等を用意する必要はなくて response で作り直していく方針のほうがいいかも
	// user の画像とかを最適化することを考えると、DB モデルと domainモデルが一対一である必要はない
	length, _, err := common.GetQueryParamForItem(c)
	if err != nil {
		c.String(403, err.Error())
	}
	trackTags, err := u.UseCase.TrackTimeLine()
	if err != nil {
		c.String(403, err.Error())
	}
	l := len(trackTags)
	if l > *length {
		trackTags = trackTags[(l - 50):]
	}
	response := NewTrackTimeLineResponse(u, trackTags)
	c.JSON(200, response)
}
