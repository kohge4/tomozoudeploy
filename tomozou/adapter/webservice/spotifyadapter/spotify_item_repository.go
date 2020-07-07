package spotifyadapter

import (
	"strconv"
	"tomozou/domain"

	"github.com/kohge4/spotify"
	"github.com/rs/zerolog/log"
)

func (h *SpotifyHandler) saveTopArtists(userID int) error {
	timerange := "long"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	// 日本語対応
	h.Client.Language = "ja"
	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	h.Client.Language = ""

	log.Info().Interface("SPOTIFY_API_RESPONSE", results).Msg("spotifyadapter/saveTopArtists")
	if err != nil {
		return err
	}
	for _, result := range results.Artists {
		var artist *domain.Artist

		artist, _ = h.ItemRepository.ReadArtistBySocialID(result.ID.String())
		if artist == nil {
			arti, err := h.Client.GetArtist(result.ID)
			if err != nil {
				log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
				return err
			}
			if result.Name == arti.Name {
				arti.Name = ""
			}
			artist = &domain.Artist{
				Name:       result.Name,
				SocialID:   result.ID.String(),
				Image:      result.Images[0].URL,
				Webservice: "spotify",
				NameOption: arti.Name,
			}
			artist.ID, err = h.ItemRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "top_artist",
			ArtistName: result.Name,
			URL:        result.ExternalURLs["Spotify"],
			Image:      result.Images[0].URL,
		}
		err = h.ItemRepository.SaveUserArtistTag(tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *SpotifyHandler) saveRecentlyFavoriteArtists(userID int) error {
	timerange := "short"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	// 日本語対応
	h.Client.Language = "ja"
	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	h.Client.Language = ""

	log.Info().Interface("SPOTIFY_API_RESPONSE", results).Msg("spotifyadapter/saveRecentFavoriteArtist")
	if err != nil {
		return err
	}
	for _, result := range results.Artists {
		var artist *domain.Artist

		artist, _ = h.ItemRepository.ReadArtistBySocialID(result.ID.String())
		if artist == nil {
			arti, err := h.Client.GetArtist(result.ID)
			if err != nil {
				log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
				return err
			}
			if result.Name == arti.Name {
				arti.Name = ""
			}
			artist = &domain.Artist{
				Name:       result.Name,
				SocialID:   result.ID.String(),
				Image:      result.Images[0].URL,
				Webservice: "spotify",
				NameOption: arti.Name,
			}
			artist.ID, err = h.ItemRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "recently_favorite_artist",
			ArtistName: result.Name,
			URL:        result.ExternalURLs["Spotify"],
			Image:      result.Images[0].URL,
		}
		err = h.ItemRepository.SaveUserArtistTag(tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *SpotifyHandler) deleteUserArtistInfo(userID int) error {
	//h.ItemRepository.DeleteAllArtitByUserID(userID)
	err := h.ItemRepository.DeleteAllUserArtistTagsByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

func (h *SpotifyHandler) saveTopTracks(userID int) error {
	timerange := "short"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	// 日本語対応
	h.Client.Language = "ja"
	results, err := h.Client.CurrentUsersTopTracksOpt(opt)
	h.Client.Language = ""

	log.Info().Interface("SPOTIFY_API_RESPONSE", results).Msg("spotifyadapter/saveTopTracks")
	if err != nil {
		return err
	}
	for _, result := range results.Tracks {
		var artistIn *domain.Artist
		var trackIn *domain.Track

		artists := result.Artists
		var ids string
		for i, a := range artists {
			artistIn, _ = h.ItemRepository.ReadArtistBySocialID(a.ID.String())
			if artistIn == nil {
				arti, err := h.Client.GetArtist(a.ID)
				if err != nil {
					log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
					return err
				}
				if a.Name == arti.Name {
					arti.Name = ""
				}
				artistIn = &domain.Artist{
					Name:       a.Name,
					SocialID:   a.ID.String(),
					Image:      arti.Images[0].URL,
					Webservice: "spotify",
					NameOption: arti.Name,
				}
				artistIn.ID, err = h.ItemRepository.SaveArtist(*artistIn)
				if err != nil {
					log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
					return err
				}
			}
			if i == len(artists)-1 {
				ids += (strconv.Itoa(artistIn.ID))
				continue
			}
			ids += (strconv.Itoa(artistIn.ID) + ",")
		}
		// track における mainのartist
		artistIn, err = h.getMainArtistOfTrack(&result)
		trackIn = &domain.Track{
			TrackName:     result.Name,
			SocialTrackID: result.SimpleTrack.ID.String(),
			ArtistName:    artistIn.Name,
			ArtistID:      artistIn.ID,
			ArtistIDs:     ids,
			AlbumName:     result.Album.Name,
			Webservice:    "spotify",
		}
		trackIn.ID, err = h.ItemRepository.SaveTrack(*trackIn)
		if err != nil {
			return err
		}
		userTrackTag := domain.NewUserTrackTag(trackIn, userID, "toptrack", 1)
		h.ItemRepository.SaveUserTrackTag(*userTrackTag)
	}
	return nil
}

func (h *SpotifyHandler) saveNowPlayingTrack(userID int) error {
	var track *spotify.FullTrack

	var artistIn *domain.Artist
	var trackIn *domain.Track

	h.Client.Language = "ja"
	nowPlaying, err := h.Client.PlayerCurrentlyPlaying()
	h.Client.Language = ""

	log.Info().Interface("SPOTIFY_API_RESPONSE", nowPlaying).Msg("spotifyadapter/saveNowPlayingTrack")
	if err != nil {
		return err
	}
	if nowPlaying.Item != nil {
		track = nowPlaying.Item
	} else {
		// ちょうど今再生しているものがない場合の処理 ==> 他のやり方あるかも
		recentlyPlaying, err := h.Client.PlayerRecentlyPlayed()
		if err != nil {
			return err
		}
		track, err = h.Client.GetTrack(recentlyPlaying[0].Track.ID)
		if err != nil {
			return err
		}
	}

	var ids string
	artists := track.Artists
	for i, artist := range artists {
		artistIn, _ = h.ItemRepository.ReadArtistBySocialID(artist.ID.String())
		if artistIn == nil {
			arti, err := h.Client.GetArtist(artist.ID)
			if artist.Name == arti.Name {
				arti.Name = ""
			}
			artistIn = &domain.Artist{
				Name:       artist.Name,
				SocialID:   artist.ID.String(),
				Image:      arti.Images[0].URL,
				Webservice: "spotify",
				NameOption: arti.Name,
			}
			artistIn.ID, err = h.ItemRepository.SaveArtist(*artistIn)
			if err != nil {
				return err
			}
		}
		if i == len(artists)-1 {
			ids += (strconv.Itoa(artistIn.ID))
			continue
		}
		ids += (strconv.Itoa(artistIn.ID) + ",")
	}

	trackIn, _ = h.ItemRepository.ReadTrackBySocialTrackID(track.ID.String())
	log.Info().Str("exTrack", trackIn.SocialTrackID).Str("nowplaying", track.ID.String()).Msg("CHECK DUPLICATE_TRACK")
	if trackIn.ID == 0 || trackIn == nil {
		artistIn, err = h.getMainArtistOfTrack(track)
		if err != nil {
			return err
		}
		trackIn = &domain.Track{
			TrackName:     track.Name,
			SocialTrackID: track.ID.String(),
			ArtistName:    artistIn.Name,
			ArtistID:      artistIn.ID,
			ArtistIDs:     ids,
			AlbumName:     track.Album.Name,
			Webservice:    "spotify",
		}
		trackIn.ID, err = h.ItemRepository.SaveTrack(*trackIn)
		if err != nil {
			return err
		}
	}

	log.Info().Interface("TrackCheck", trackIn).Msg("CHECK_BEFORE_SAVE_USERTRACKTAG")
	var userTrackTag *domain.UserTrackTag
	//lastTag, _ := h.ItemRepository.ReadUserTrackTagByUserIDANDTagName(userID, "nowplaying")
	lastTag, _ := h.ItemRepository.ReadUserTrackTagByUserIDANDTagNameANDTrackID(userID, "nowplaying", trackIn.ID)
	if len(lastTag) > 0 {
		userTrackTag = domain.NewUserTrackTag(trackIn, userID, "nowplaying", lastTag[0].UserTrackTag.Count+1)
		h.ItemRepository.DeleteUserTrackTag(lastTag[0].UserTrackTag)
		h.ItemRepository.SaveUserTrackTag(*userTrackTag)
		return nil
	}
	userTrackTag = domain.NewUserTrackTag(trackIn, userID, "nowplaying", 1)
	h.ItemRepository.SaveUserTrackTag(*userTrackTag)
	return nil
}

func (h *SpotifyHandler) checkDupulicateArtist(socialID string) bool {
	artist, _ := h.ItemRepository.ReadArtistBySocialID(socialID)
	if artist == nil {
		return false
	}
	return true
}

func (h *SpotifyHandler) updateNowplayingUserTrackTag(tag domain.UserTrackTag) {
	h.ItemRepository.DeleteUserTrackTag(tag)
	h.ItemRepository.SaveUserTrackTag(tag)
}
