package appleadapter

import (
	"fmt"
	"tomozou/domain"

	applemusic "github.com/kohge4/go-apple-music-sdk"
)

func updateSearchObjByAppleAPIResponse(r *applemusic.Search, searchObj *domain.SearchObj) *domain.SearchObj {
	results := r.Results

	var searchResult *domain.SearchResult
	if searchObj.SearchKey == "track" {
		tracks := results.Songs.Data
		for _, t := range tracks {
			searchResult = newSearchResultByTrackKey(
				t.Attributes.ArtistName,
				t.Attributes.Name,
				t.Id,
				t.Attributes.URL,
			)
			accuracy := searchResult.UpdateAccuracy()
			if accuracy < 0.8 {
				break
			}
			searchObj.Results = append(searchObj.Results, *searchResult)
		}
	}
	fmt.Println(searchResult)
	return searchObj
}

func newSearchResultByTrackKey(artistName string, trackName string, socialID string, url string) *domain.SearchResult {
	return &domain.SearchResult{
		ResultKey:  "tarck",
		ArtistName: artistName,
		TrackName:  trackName,
		SocialID:   socialID,
		URL:        url,
	}
}

type SaveOptions struct {
	Accuracy float64
}
