package appleadapter

import (
	"tomozou/domain"

	applemusic "github.com/kohge4/go-apple-music-sdk"
	"github.com/rs/zerolog/log"
)

// 有効期限切れとかのデバッグ用
const APToken = "eyJhbGciOiJFUzI1NiIsImtpZCI6IkJRQzdMTFNOQ0IifQ.eyJleHAiOjE1OTM3NzY1MTgsImlhdCI6MTU5Mzc3NjUxOCwiaXNzIjoiNFFMVzRINzY2UyJ9.QuVeH4sKCtGxZUfxXToeqmr-3teVoHb_7ipLTAu7qJvR_PzsS0kKrQY6SN7Pnq_sKVamZYHRKAxZWimINM_XNA"

// apple api response から 必要なデータのみを取り出す
func createSearchObjByAppleAPIResponse(r *applemusic.Search, searchObj *domain.SearchObj) *domain.SearchObj {
	results := r.Results

	var searchResult *domain.SearchResult
	if searchObj.SearchKey == "track" {
		tracks := results.Songs.Data
		for _, t := range tracks {
			log.Info().Interface("SEARCHOBJ_CONVERT_TRACK", t).Msg("adapter/appleadapter/createSearchObjByAppleAPIResponse()")

			searchResult = newSearchResultByTrackKey(
				t.Attributes.ArtistName,
				t.Attributes.Name,
				t.Id,
				t.Attributes.URL,
				"tr",
			)
			accuracy := searchObj.GetAccuracy(searchResult)
			log.Info().Interface("SEARCHOBJ_CONVERT_SEARCHRESULT_TRACK", searchResult).Interface("SEARCHOBJ_CONVERT_ACCURACY", accuracy).Msg("adapter/appleadapter/createSearchObjByAppleAPIResponse()")

			if accuracy > 80 {
				searchResult.Accuracy = accuracy
				//searchObj.Results = []domain.SearchResult{*searchResult}
				searchObj.Results = append(searchObj.Results, *searchResult)
				break
			}
			if accuracy < 70 {
				continue
			}
		}
	} else if searchObj.SearchKey == "artist" {
		data := results.Artists.Data
		for _, t := range data {
			log.Info().Interface("SEARCHOBJ_CONVERT_ARTIST", t).Msg("adapter/appleadapter/createSearchObjByAppleAPIResponse()")

			searchResult = newSearchResultByArtistKey(
				t.Attributes.Name,
				t.Id,
				t.Attributes.URL,
				"ar",
			)
			accuracy := searchObj.GetAccuracy(searchResult)
			log.Info().Interface("SEARCHOBJ_CONVERT_SEARCHRESULT_ARTIST", searchResult).Interface("SEARCHOBJ_CONVERT_ACCURACY", accuracy).Msg("adapter/appleadapter/createSearchObjByAppleAPIResponse()")

			if accuracy > 80 {
				searchResult.Accuracy = accuracy
				searchObj.Results = []domain.SearchResult{*searchResult}
				break
			}
			if accuracy < 70 {
				continue
			}
			searchResult.Accuracy = accuracy
			searchObj.Results = append(searchObj.Results, *searchResult)
		}
	}
	return searchObj
}

func newSearchResultByTrackKey(artistName string, trackName string, socialID string, url string, resultKey string) *domain.SearchResult {
	return &domain.SearchResult{
		ResultKey:  resultKey,
		ArtistName: artistName,
		TrackName:  trackName,
		SocialID:   socialID,
		URL:        url,
	}
}

func newSearchResultByArtistKey(artistName string, socialID string, url string, resultKey string) *domain.SearchResult {
	return &domain.SearchResult{
		ResultKey:  resultKey,
		ArtistName: artistName,
		SocialID:   socialID,
		URL:        url,
	}
}

type SaveOptions struct {
	Accuracy float64
}
