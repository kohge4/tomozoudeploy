package appleadapter

import (
	"context"
	"strings"
	"tomozou/domain"

	applemusic "github.com/kohge4/go-apple-music-sdk"
	"github.com/rs/zerolog/log"
)

func (h *AppleHandler) SearchWebServiceItem(searchObj *domain.SearchObj) error {
	return nil
}

func (h *AppleHandler) SearchWebServiceItemAndCreateItemTag(searchObj *domain.SearchObj) error {
	// 関数で検索結果を判定
	//word := searchObj.SearchTrackName
	artistName := strings.Replace(searchObj.SearchArtistName, " ", "+", -1)
	trackName := strings.Replace(searchObj.SearchTrackName, " ", "+", -1)
	searchWord := trackName + "+" + artistName
	searchOpt := &applemusic.SearchOptions{
		Term: searchWord,
	}

	ctx := context.Background()
	r, resp, err := h.Client.Catalog.Search(ctx, "jp", searchOpt)
	log.Info().Interface("APPLE_API_RESP", resp.StatusCode).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	if resp.StatusCode != 200 {
		// token の有効期限が切れていた時に token を 更新
		// statuscode が 401 の時にしたほうがいいかも
		h.UpdateWebServiceConfig()
		r, resp, err = h.Client.Catalog.Search(ctx, "jp", searchOpt)
	}
	log.Info().Interface("SEARCH_KEYWORD", searchOpt).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	log.Info().Interface("API_RESPONSE", r).Interface("SEARCH_KEYWORD", searchOpt).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	if err != nil {
		log.Debug().Str("%v", err.Error()).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
		return err
	}

	// ここを変更することで、artistName とか TrackName にも対応する
	// この結果次第で artistname のみの結果を走らせたりもする
	searchObj = createSearchObjByAppleAPIResponse(r, searchObj)

	if len(searchObj.Results) == 0 {
		// NameOption 再試行: 日本語名でうまくいかなかった時
		artistName = strings.Replace(searchObj.ArtistNameOption, " ", "+", -1)
		trackName = strings.Replace(searchObj.TrackNameOption, " ", "+", -1)
		r, _, err = h.Client.Catalog.Search(ctx, "jp", searchOpt)
		if err != nil || len(searchObj.Results) == 0 {
			log.Debug().Interface("SEARCHOBJ_CONVERT_API_RESPONSE", err).Interface("RESPONSE_LEGTH", len(searchObj.Results)).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
			return nil
		}
	}
	log.Info().Str("[CONNECTOR_IMPL] searchResult", searchObj.Results[0].URL).Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	if searchObj.SearchKey == "track" {
		h.createTrackWebServiceTagBySearchObj(searchObj)
	} else if searchObj.SearchKey == "artist" {
		h.createArtistWebServiceTagBySearchObj(searchObj)
	}
	return nil
}

func (h *AppleHandler) createTrackWebServiceTagBySearchObj(searchObj *domain.SearchObj) {
	for _, s := range searchObj.Results {
		tag := &domain.TrackWebServiceTag{
			TrackID:        searchObj.ItemID,
			WebServiceID:   "ap",
			SocialTrackURL: s.URL,
			SocialTrackID:  s.SocialID,
			SearchBy:       searchObj.SearchKey,
		}
		h.ItemChildRepository.SaveTrackWebServiceTag(tag)
		log.Info().Str("[CONNECTOR_IMPL]SAVE_COMPLETE", "").Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	}
}

func (h *AppleHandler) createArtistWebServiceTagBySearchObj(searchObj *domain.SearchObj) {
	for _, s := range searchObj.Results {
		tag := &domain.ArtistWebServiceTag{
			ArtistID:        searchObj.ItemID,
			WebServiceID:    "ap",
			SocialArtistURL: s.URL,
			SocialArtistID:  s.SocialID,
		}
		h.ItemChildRepository.SaveArtistWebServiceTag(tag)
		log.Info().Str("[CONNECTOR_IMPL]SAVE_COMPLETE", "").Msg("adapter/appleadapter/SearchWebServiceItemAndCreateItemTag")
	}
}
