package mainappimpl

import (
	jwt "github.com/appleboy/gin-jwt/v2"

	"tomozou/adapter/webservice/spotifyadapter"
	"tomozou/usecase"
)

type UserProfileApplicationImpl struct {
	UseCase *usecase.UserProfileApplication

	Handler        *spotifyadapter.SpotifyHandler
	SecondHandler  *spotifyadapter.SpotifyHandler
	AuthMiddleware *jwt.GinJWTMiddleware
}
