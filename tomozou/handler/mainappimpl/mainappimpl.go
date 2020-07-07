package mainappimpl

import (
	jwt "github.com/appleboy/gin-jwt/v2"

	"tomozou/adapter/webservice/appleadapter"
	"tomozou/adapter/webservice/spotifyadapter"
	"tomozou/usecase"
)

type UserProfileApplicationImpl struct {
	UseCase *usecase.UserProfileApplication

	Handler *spotifyadapter.SpotifyHandler
	//Connector      connectorappimpl.ConnectorApplicationImpl
	Connector      *appleadapter.AppleHandler
	AuthMiddleware *jwt.GinJWTMiddleware
}
