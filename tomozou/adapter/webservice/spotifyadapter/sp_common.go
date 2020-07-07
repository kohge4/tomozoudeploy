package spotifyadapter

import (
	"tomozou/domain"

	"github.com/kohge4/spotify"
	"github.com/rs/zerolog/log"
)

func (h *SpotifyHandler) getMainArtistOfTrack(track *spotify.FullTrack) (*domain.Artist, error) {
	//albumArtist := track.Album.Artists
	// alnumが完全に狭窄だった場合は popularityが高い方
	albumArtists := track.Album.Artists
	if len(albumArtists) == 1 {
		artist, err := h.ItemRepository.ReadArtistBySocialID(albumArtists[0].ID.String())
		if err != nil {
			return nil, err
		}
		log.Info().Interface("ALBUM_BY_A_ARTIST", artist.Name).Interface("SPOTIFY_API_RESPONSE_POP", artist.SocialID).Msg("spotifyadapter/saveNowPlayingTrack")
		return artist, nil
	} else {
		// 結局 クレジットの先に名前が来てる方
		artist, err := h.ItemRepository.ReadArtistBySocialID(albumArtists[0].ID.String())
		if err != nil {
			return nil, err
		}
		return artist, nil
		/*
			// popularity の高い方を表示する場合 (アルバムが複数Artistで作成されてる場合)
			artistPop := &spotify.FullArtist{}
			for _, a := range albumArtists {
				artistFull, err := h.Client.GetArtist(a.ID)
				log.Info().Interface("SPOTIFY_API_RESPONSE_NAME", artistFull.Name).Interface("SPOTIFY_API_RESPONSE_POP", artistFull.Popularity).Msg("spotifyadapter/saveNowPlayingTrack")
				if err != nil {
					return nil, err
				}
				if artistPop.Popularity < artistFull.Popularity {
					artistPop = artistFull
				}
			}
			artist, err := h.ItemRepository.ReadArtistBySocialID(artistPop.ID.String())
			if err != nil {
				return nil, err
			}
			return artist, nil
		*/
	}
}

/*
// track の artist のデータがあったらそのIDを使用, なかったら保存
func (h *SpotifyHandler) trackArtistIDs(artists []spotify.FullArtist) (string, error) {
	var ids string
	for i, a := range artists {
		artistIn, _ := h.SpotifyRepository.ReadArtistBySocialID(a.ID.String())
		if artistIn == nil {
			arti, err := h.Client.GetArtist(a.ID)
			if err != nil {
				log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
			}
			artistIn = &domain.Artist{
				Name:       a.Name,
				SocialID:   a.ID.String(),
				Image:      arti.Images[0].URL,
				Webservice: "spotify",
			}
			artistIn.ID, err = h.SpotifyRepository.SaveArtist(*artistIn)
			if err != nil {
				log.Debug().Interface("SPOTIFY_API_RESPONSE_ERROR", err.Error()).Msg("spotifyadapter/saveNowPlayingTrack")
			}
		}
		if i == len(artists)-1 {
			ids += (strconv.Itoa(artistIn.ID))
			continue
		}
		ids += (strconv.Itoa(artistIn.ID) + ",")
	}
	return ids, nil
}

*/
