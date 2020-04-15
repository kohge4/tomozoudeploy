package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/domain"
	"tomozou/handler"
	"tomozou/infra/datastore"
	"tomozou/middleware/auth"
	"tomozou/usecase"
)

func main() {
	DRIVER := "mysql"
	//DSN := "root:@(db:3306)/tomozou?charset=utf8&parseTime=True"
	DSN := "root:@unix(/cloudsql/ongakuconnection:asia-northeast1:ongkdb)/tomozoudb?charset=utf8&parseTime=True"
	//"ユーザー名:パスワード@unix(/cloudsql/インスタンス接続名)/DB名"

	gormConn, _ := datastore.GormConn(DRIVER, DSN)
	userRepo := datastore.NewUserDBRepository(gormConn)
	itemRepo := datastore.NewSpotifyItemDBRepository(gormConn)

	useCase := usecase.NewUserProfileApplication(userRepo, itemRepo)

	spotifyHandler := webservice.NewSpotifyHandler(userRepo, itemRepo, gormConn)
	authMiddleware := auth.AuthUser()

	userProfileAppImpl := handler.UserProfileApplicationImpl{
		UseCase: useCase,

		Handler:        spotifyHandler,
		AuthMiddleware: authMiddleware,
	}

	r := gin.Default()

	crs := cors.DefaultConfig()
	crs.AllowOrigins = []string{"http://localhost:8080", "https://tomozoufront.firebaseapp.com", "https://ongakuconnection.com", "https://ongakuuconnection.firebaseapp.com/"}
	crs.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}
	r.Use(cors.New(crs))

	r.GET("/search/user/artistid/:artistID", userProfileAppImpl.SearchUsersByArtistID)
	r.GET("/search/user/artistname", userProfileAppImpl.SearchUsersByArtistName)
	r.GET("/timeline", userProfileAppImpl.TrackTimeLine)

	// Spotify ログイン処理用エンドポイント
	rSpo := r.Group("/spotify")
	{
		rSpo.GET("/callback", userProfileAppImpl.Callback)
		rSpo.GET("/login", userProfileAppImpl.Login)
		rSpo.GET("/myartist", userProfileAppImpl.MyArtist)
	}

	// 認証用エンドポイント: JWTの検証を毎回行う
	auth := r.Group("/me")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/profile", userProfileAppImpl.MyProfile)
		// mytrack 表示
		auth.GET("/track", userProfileAppImpl.MyTrack)
		// nowplayingは track のみ連携する処理
		auth.GET("/nowplaying", userProfileAppImpl.NowPlaying)
	}

	// 開発用: データ確認エンドポイント
	devUserRepo := datastore.NewDevUserRepo(gormConn)
	rDev := r.Group("/dev")
	{
		rDev.GET("/user", func(c *gin.Context) {
			users, _ := devUserRepo.CheckUser()
			c.JSON(200, users)
		})
		rDev.GET("/tag", func(c *gin.Context) {
			tags := []domain.UserArtistTag{}
			devUserRepo.DB.Find(&tags)
			c.JSON(200, tags)
		})
		rDev.GET("/artist", func(c *gin.Context) {
			artist := []domain.Artist{}
			devUserRepo.DB.Find(&artist)
			c.JSON(200, artist)
		})
		rDev.GET("/track", func(c *gin.Context) {
			track := []domain.Track{}
			devUserRepo.DB.Find(&track)
			c.JSON(200, track)
		})
		rDev.GET("/tracktag", func(c *gin.Context) {
			track := []domain.UserTrackTag{}
			devUserRepo.DB.Find(&track)
			c.JSON(200, track)
		})
		rDev.GET("/mytrack", userProfileAppImpl.MyTrack)
		rDev.GET("/timeline", userProfileAppImpl.TrackTimeLine)
		rDev.GET("/userdata", func(c *gin.Context) {
		})
		rDev.GET("/debug", userProfileAppImpl.Debug)
		rDev.GET("/deptest", func(c *gin.Context) {
			c.String(200, "deploy test")
		})
	}

	// Chat 用: authによるJWT 以下から
	chatRepo := datastore.NewChatDBRepository(gormConn)
	chatApp := usecase.ChatApplication{
		ItemRepository: itemRepo,
		ChatRepository: chatRepo,
	}
	chatAppImpl := handler.ChatApplicationImpl{
		UseCase: chatApp,
	}
	rChat := r.Group("/chat")
	rChat.Use(authMiddleware.MiddlewareFunc())
	{
		rChat.GET("/room", chatAppImpl.DisplayChatRoom)
		rChat.POST("/user/comment", chatAppImpl.UserChat)
		rChat.GET("/list/:artistID", chatAppImpl.DisplayChatListByArtist)
	}
	r.Run(":8080")
}
