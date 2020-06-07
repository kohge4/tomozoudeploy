package connectorappimpl

import (
	"context"
	"fmt"
	"strconv"

	"tomozou/adapter/webservice/appleadapter"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
	applemusic "github.com/kohge4/go-apple-music-sdk"
)

type ConnectorApplicationImpl struct {
	AppleHandler *appleadapter.AppleHandler

	// *domain.ItemRepository を ポインタにすると失敗する
	ItemRepository domain.ItemRepository
}

func (app *ConnectorApplicationImpl) AppleConnectorDemo(c *gin.Context) {
	// searchObj を adapter に流し込む役割
	word := c.Param("word")
	searchOpt := &applemusic.SearchOptions{
		Term: word,
	}

	ctx := context.Background()
	storefronts, _, err := app.AppleHandler.Client.Catalog.Search(ctx, "jp", searchOpt)
	if err != nil {
		c.String(401, err.Error())
	}
	fmt.Printf("%T \n", storefronts.Results)
	fmt.Printf("%v \n", storefronts.Results)
	fmt.Printf("%T \n", storefronts.Results.Albums)
	resp := storefronts.Results.Albums.Data
	c.JSON(200, resp)
}

func (app *ConnectorApplicationImpl) AppleConnectorByTrack(c *gin.Context) {
	// searchObj を adapter に流し込む役割
	//word := c.Param("word")
	searchObj := &domain.SearchObj{
		SearchKey:        "track",
		SearchArtistName: "東京事変",
		SearchTrackName:  "キラーチューン",
		ItemType:         "",
		ItemID:           1,
	}
	err := app.AppleHandler.SearchWebServiceItemAndCreateItemTag(searchObj)
	if err != nil {
		c.String(401, err.Error())
	}
	c.JSON(200, searchObj)
}

func (app *ConnectorApplicationImpl) CreateAppleTrackWebServiceTagByTrackID(c *gin.Context) {
	// searchObj を adapter に流し込む役割
	//word := c.Param("word")
	param := c.Param("trackID")
	trackID, _ := strconv.Atoi(param)
	if trackID == 0 {
		trackID = 1
	}

	track, _ := app.ItemRepository.ReadTrackByTrackID(trackID)
	//track.ArtistName = "neveryoung+beach"
	searchObj := &domain.SearchObj{
		SearchKey:        "track",
		SearchArtistName: track.ArtistName,
		SearchTrackName:  track.TrackName,
		ItemType:         "",
		ItemID:           trackID,
	}

	err := app.AppleHandler.SearchWebServiceItemAndCreateItemTag(searchObj)
	if err != nil {
		c.String(401, err.Error())
	}
	c.JSON(200, searchObj)
}

func (app *ConnectorApplicationImpl) SearchAppleTrackAndCreateTrackwevServiceTag(trackID int) (*domain.SearchObj, error) {
	track, _ := app.ItemRepository.ReadTrackByTrackID(trackID)
	searchObj := &domain.SearchObj{
		SearchKey:        "track",
		SearchArtistName: track.ArtistName,
		SearchTrackName:  track.TrackName,
		ItemType:         "",
		ItemID:           trackID,
	}
	err := app.AppleHandler.SearchWebServiceItemAndCreateItemTag(searchObj)
	if err != nil {
		return nil, err
	}
	return searchObj, nil
}
