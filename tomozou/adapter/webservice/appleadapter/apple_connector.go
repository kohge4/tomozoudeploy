package appleadapter

import (
	"context"
	"tomozou/domain"

	applemusic "github.com/kohge4/go-apple-music-sdk"
	"github.com/rs/zerolog/log"
)

func (h *AppleHandler) SearchWebServiceItem(searchObj *domain.SearchObj) error {
	return nil
}

func (h *AppleHandler) SearchWebServiceItemAndCreateItemTag(searchObj *domain.SearchObj) error {
	// 関数で検索結果を判定
	word := searchObj.SearchArtistName
	//word := searchObj.SearchArtistName + "+" + searchObj.SearchTrackName
	searchOpt := &applemusic.SearchOptions{
		Term: word,
	}

	ctx := context.Background()
	r, _, err := h.Client.Catalog.Search(ctx, "jp", searchOpt)
	if err != nil {
		log.Debug().Str("%v", err.Error()).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
		return err
	}
	searchObj = updateSearchObjByAppleAPIResponse(r, searchObj)
	//fmt.Println(r)
	log.Info().Str("[CONNECTOR_IMPL] searchResult", searchObj.Results[0].URL).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")

	for _, s := range searchObj.Results {
		//fmt.Println(s)
		tag := &domain.TrackWebServiceTag{
			TrackID:       searchObj.ItemID,
			WebServiceID:  "ap",
			SocialURL:     s.URL,
			SocialTrackID: s.SocialID,
		}
		h.ItemChildRepository.SaveTrackWebServiceTag(tag)
		log.Info().Str("[CONNECTOR_IMPL]SAVE_COMPLETE", "").Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	}
	return nil
}

func AppleTrackResponseToSearchResult() {}

func AppleArtistResponseToSearchResult() {}

func UpdateSearchObjByAppleTrackResponse() {}

func UpdateSearchObjByAppleArtistResponse() {}
