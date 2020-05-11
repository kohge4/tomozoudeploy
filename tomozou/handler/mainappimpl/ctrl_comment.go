package mainappimpl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (u *UserProfileApplicationImpl) GetTrackCommentWithUserByTrackID(c *gin.Context) {
	// query param 　に trackid を持たせる方針にしよう
	trackID := c.Param("trackID")
	id, _ := strconv.Atoi(trackID)
	commentList, err := u.UseCase.GetTrackCommentWithUserByTrackID(id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	//response, _ := chatResponse(chatList)
	c.JSON(200, commentList)
}
