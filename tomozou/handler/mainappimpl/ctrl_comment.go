package mainappimpl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (u *UserProfileApplicationImpl) AddTrackComment(c *gin.Context) {
	var jsonBody interface{}
	c.BindJSON(&jsonBody)

	var commentIn domain.TrackComment
	jsonByte, _ := json.Marshal(jsonBody)
	err := json.Unmarshal(jsonByte, &commentIn)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		log.Debug().Msg("ERROR: " + err.Error())
	}
	fmt.Println(commentIn)
	u.UseCase.AddTrackComment(&commentIn)
	c.JSON(200, commentIn)
}

func (u *UserProfileApplicationImpl) GetTrackCommentByTrackID(c *gin.Context) {}
