package mainappimpl

import (
	"strconv"
	"tomozou/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (u *UserProfileApplicationImpl) ShowAppleMusic(c *gin.Context) {
	// 手順: DB を検索 => なかったらAPI を叩く => 表示 resp{track artist} 的に表示
	t, err := getTrackWithTrackWebServiceTagsOfApple(u, c)
	if err != nil {
		c.String(401, err.Error())
		return
	}
	// 同じ処理を書く(ArtistWebServiceTagについて)
	a, err := getArtistWithArtistWebServiceTagsOfApple(u, c, t)
	if err != nil {
		c.String(401, err.Error())
		return
	}
	log.Info().Interface("CHECK_ARTIST_WEBSERVICETAG", a).Msg("mainappimpl/ShowAppleMusic ")

	var response *ConnectedTrackResponse
	if t.WebServiceTags != nil && a.WebServiceTags != nil {
		response = NewConnectedTracksAndArtistResponse(t.WebServiceTags, a)
		c.JSON(200, response)
		return
	} else if t.WebServiceTags != nil {
		response = NewConnectedTracksResponse(t.WebServiceTags)
		c.JSON(200, response)
		return
	} else {
		c.JSON(200, t)
		return
	}
}

func getTrackWithTrackWebServiceTagsOfApple(u *UserProfileApplicationImpl, c *gin.Context) (*domain.TrackWithTrackWebServiceTags, error) {
	trackIDString := c.Param("trackID")
	trackID, _ := strconv.Atoi(trackIDString)
	t, err := u.UseCase.GetTrackWithTrackWebServiceTagByTrackID(trackID)

	log.Info().Interface("CHECK_TRACKWEBSERVICETAG", t).Msg("mainappimpl/ShowAppleMusic ")
	if t.WebServiceTags == nil {
		// artistNameOption を使用するために
		searchObj := u.UseCase.NewSearchObjByTrack(t)
		u.Connector.SearchWebServiceItemAndCreateItemTag(searchObj)
		log.Info().Interface("CHECK_SEARCHOBJ_RESULTS", searchObj.Results).Msg("mainappimpl/ShowAppleMusic ")

		t, err = u.UseCase.GetTrackWithTrackWebServiceTagByTrackID(trackID)
		log.Info().Interface("CHECK_TRACKWEBSERVICETAG_SECOND", t).Msg("mainappimpl/ShowAppleMusic ")
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		log.Info().Interface("[CTRL_MAIN]", t).Msg("mainappimpl/ShowAppleMusic ")
		c.String(401, err.Error())
	}
	log.Info().Interface("CHECK_TRACK_WEBSERVICETAG_BEFORE_RESPONSE", t).Msg("mainappimpl/ShowAppleMusic ")
	return t, nil
}

func getArtistWithArtistWebServiceTagsOfApple(u *UserProfileApplicationImpl, c *gin.Context, t *domain.TrackWithTrackWebServiceTags) (*domain.ArtistWithArtistWebServiceTags, error) {
	// 完全に分離させて考えたほうがいい
	// apple api のtarck response に artist 関連の ID とかはない
	artistID := t.Track.ArtistID
	a, err := u.UseCase.GetArtistWithArtistWebServiceTagByTrackID(artistID)
	if err != nil {
		return nil, err
	}
	if a.WebServiceTags == nil {
		// artistNameOption を使用するために
		searchObj := u.UseCase.NewSearchObjByArtist(a)
		u.Connector.SearchWebServiceItemAndCreateItemTag(searchObj)
		log.Info().Interface("CHECK_SEARCHOBJ_RESULTS", searchObj.Results).Msg("mainappimpl/ShowAppleMusic ")

		a, err = u.UseCase.GetArtistWithArtistWebServiceTagByTrackID(artistID)
		log.Info().Interface("CHECK_ARTIST_WEBSERVICETAG_SECOND", a).Msg("mainappimpl/ShowAppleMusic ")
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		log.Info().Interface("[CTRL_MAIN]", t).Msg("mainappimpl/ShowAppleMusic ")
		c.String(401, err.Error())
	}
	log.Info().Interface("CHECK_ARTIST_EBSERVICETAG_BEFORE_RESPONSE", a).Msg("mainappimpl/ShowAppleMusic ")
	return a, nil
}
