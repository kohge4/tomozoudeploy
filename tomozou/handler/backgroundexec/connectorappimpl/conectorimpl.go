package connectorappimpl

import (
	"context"
	"fmt"

	"tomozou/adapter/webservice/appleadapter"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
	applemusic "github.com/kohge4/go-apple-music-sdk"
)

type ConnectorApplicationImpl struct {
	AppleHandler *appleadapter.AppleHandler
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
		SearchArtistName: "cero",
		SearchTrackName:  "orphans",
		ItemType:         "",
		ItemID:           1,
	}
	err := app.AppleHandler.SearchWebServiceItemAndCreateItemTag(searchObj)
	if err != nil {
		c.String(401, err.Error())
	}
	c.JSON(200, searchObj)
}
