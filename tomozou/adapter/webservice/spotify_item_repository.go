package webservice

import (
	"tomozou/domain"

	"github.com/kohge4/spotify"
)

func (h *SpotifyHandler) saveTopArtists(userID int) error {
	timerange := "long"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artist *domain.Artist

		artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
		if artist == nil {
			artist = &domain.Artist{
				Name:     result.Name,
				SocialID: result.ID,
				Image:    result.Images[0].URL,
			}
			artist.ID, err = h.SpotifyRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "top_artist",
			ArtistName: result.Name,
			URL:        result.ExternalUrls.Spotify,
			Image:      result.Images[0].URL,
		}
		err = h.SpotifyRepository.SaveUserArtistTag(tag)
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

	results, err := h.Client.CurrentUsersTopArtistsOpt(opt)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artist *domain.Artist

		artist, _ = h.SpotifyRepository.ReadArtistBySocialID(result.ID)
		if artist == nil {
			artist = &domain.Artist{
				Name:     result.Name,
				SocialID: result.ID,
				Image:    result.Images[0].URL,
			}
			artist.ID, err = h.SpotifyRepository.SaveArtist(*artist)
			if err != nil {
				return err
			}
		}
		tag := domain.UserArtistTag{
			UserID:     userID,
			ArtistID:   artist.ID,
			TagName:    "recently_favorite_artist",
			ArtistName: result.Name,
			URL:        result.ExternalUrls.Spotify,
			Image:      result.Images[0].URL,
		}
		err = h.SpotifyRepository.SaveUserArtistTag(tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *SpotifyHandler) deleteUserArtistInfo(userID int) error {
	//h.SpotifyRepository.DeleteAllArtitByUserID(userID)
	err := h.SpotifyRepository.DeleteAllUserArtistTagsByUserID(userID)
	if err != nil {
		return err
	}
	return nil
}

func (h *SpotifyHandler) saveTopTracks(userID int) error {
	/*
		trackのデータを取得
		artist データがなければ追加 => artist保存
		trackの保存データに変換
	*/
	timerange := "short"
	limit := 5
	opt := &spotify.Options{
		Timerange: &timerange,
		Limit:     &limit,
	}

	results, err := h.Client.GetUserTopTracks2Opt(opt)
	if err != nil {
		return err
	}
	for _, result := range results.Items {
		var artistIn *domain.Artist
		var trackIn *domain.Track

		artists := result.Album.Artists

		artistIn, _ = h.SpotifyRepository.ReadArtistBySocialID(artists[0].ID)
		if artistIn == nil {
			artistIn = &domain.Artist{
				Name:     result.Artists[0].Name,
				SocialID: result.Artists[0].ID,
				Image:    "",
			}
			artistIn.ID, err = h.SpotifyRepository.SaveArtist(*artistIn)
			if err != nil {
				return err
			}
		}

		trackIn = &domain.Track{
			TrackName: result.Name,
			// TrackURL ではなくsocialID で url作る方針
			SocialTrackID: result.ID,
			ArtistName:    artistIn.Name,
			ArtistID:      artistIn.ID,
		}
		trackIn.ID, err = h.SpotifyRepository.SaveTrack(*trackIn)
		if err != nil {
			return err
		}

		userTrackTag := domain.NewUserTrackTag(trackIn, userID, "toptrack")
		h.SpotifyRepository.SaveUserTrackTag(*userTrackTag)
		// 複数の arthist が 携わるトラックについてちゃんとやった方がいいかも
	}
	return nil
}

func (h *SpotifyHandler) saveNowPlayingTrack(userID int) error {
	var track *spotify.SimpleTrack

	var artistIn *domain.Artist
	var trackIn *domain.Track

	nowPlaying, err := h.Client.PlayerCurrentlyPlaying()
	if err != nil {
		return err
	}

	if nowPlaying.Item != nil {
		track = nowPlaying.Item.ToSimpleTrack()
	} else {
		recentlyPlaying, err := h.Client.PlayerRecentlyPlayed()
		if err != nil {
			return err
		}
		track = &recentlyPlaying[0].Track
	}
	artists := track.Artists

	artistIn, _ = h.SpotifyRepository.ReadArtistBySocialID(artists[0].ID.String())
	if artistIn == nil {
		artistIn = &domain.Artist{
			Name:     track.Artists[0].Name,
			SocialID: track.Artists[0].ID.String(),
			Image:    "",
		}
		artistIn.ID, err = h.SpotifyRepository.SaveArtist(*artistIn)
		if err != nil {
			return err
		}
	}
	trackIn = &domain.Track{
		TrackName: track.Name,
		// TrackURL ではなくsocialID で url作る方針
		SocialTrackID: track.ID.String(),
		ArtistName:    artistIn.Name,
		ArtistID:      artistIn.ID,
	}
	// Track 保存に関する処理
	// SimpleTrack を変換 => artist を保存 => tagとして track に持たせる
	trackIn.ID, err = h.SpotifyRepository.SaveTrack(*trackIn)
	if err != nil {
		return err
	}
	userTrackTag := domain.NewUserTrackTag(trackIn, userID, "nowplaying")
	lastTag, _ := h.SpotifyRepository.ReadUserTrackTagByUserIDANDTagName(userID, "nowplaying")
	if len(lastTag) < 1 {
		h.SpotifyRepository.SaveUserTrackTag(*userTrackTag)
		return nil
	}
	//if userTrackTag.TrackSocialID == lastTag[len(lastTag)-1].TrackSocialID {
	//return nil
	//}
	h.SpotifyRepository.SaveUserTrackTag(*userTrackTag)
	return nil
}

func (h *SpotifyHandler) checkDupulicateArtist(socialID string) bool {
	artist, _ := h.SpotifyRepository.ReadArtistBySocialID(socialID)
	if artist == nil {
		return false
	}
	return true
}
