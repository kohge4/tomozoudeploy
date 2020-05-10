package mainappimpl

import (
	"fmt"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
)

func (u *UserProfileApplicationImpl) AddTrackComment(c *gin.Context) {
	var jsonBody interface{}
	c.BindJSON(&jsonBody)

	var commentIn domain.TrackComment
	fmt.Println(commentIn)
}

func (u *UserProfileApplicationImpl) GetTrackComment(c *gin.Context) {}
