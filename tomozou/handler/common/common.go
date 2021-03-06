package common

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIDFromContext(c *gin.Context) (int, error) {
	id, _ := c.Get("userid")
	userID, ok := id.(float64)
	if ok == false {
		c.String(403, "Authentication is failed")
		return 0, fmt.Errorf("Authentication is failed")
	}
	if userID == 0 {
		c.String(403, "Authentication is failed")
		return 0, fmt.Errorf("Authentication is failed")
	}
	return int(userID), nil
}

func GetQueryParamForItem(c *gin.Context) (*int, *int, error) {
	lengthString := c.Query("length")
	if lengthString == "" {
		lengthString = "50"
	}
	length, err := strconv.Atoi(lengthString)
	if err != nil {
		return nil, nil, err
	}

	offsetString := c.Query("offset")
	if offsetString == "" {
		offsetString = "0"
	}
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return nil, nil, err
	}
	return &length, &offset, nil
}

func embedTrackURLFromSopotifyID() {}

func embedArtistURLFromSopotifyID() {}

func embedTrackURLFromAppleID() {}

func embedArtistURLFromAppleID() {}
