package mainappimpl

import (
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) TrackTimeLine(c *gin.Context) {
	// domain model に画像等を用意する必要はなくて response で作り直していく方針のほうがいいかも
	// user の画像とかを最適化することを考えると、DB モデルと domainモデルが一対一である必要はない
	/*
		length, _, err := common.GetQueryParamForItem(c)
		if err != nil {
			c.String(403, err.Error())
		}
	*/
	lengthInt := 50
	length := &lengthInt
	trackTags, err := u.UseCase.TrackTimeLine()
	if err != nil {
		c.String(403, err.Error())
		log.Debug().Str("ERROR", err.Error()).Msg("mainappimpl/TrackTimeLine ")
		return
	}
	l := len(trackTags)
	if l > *length {
		trackTags = trackTags[(l - 50):]
	}
	response := NewTrackTimeLineResponse(u, trackTags)
	log.Info().Interface("[CTRL_MAIN]", response).Msg("mainappimpl/TrackTimeLine ")
	c.JSON(200, response)
}
