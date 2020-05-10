package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"tomozou/adapter/webservice"
	"tomozou/domain"
	"tomozou/handler/chatappimpl"
	"tomozou/handler/mainappimpl"
	"tomozou/infra/datastore"
	"tomozou/infra/datastore/chatrepoimpl"
	"tomozou/infra/datastore/itemrepoimpl"
	"tomozou/infra/datastore/userrepoimpl"
	"tomozou/middleware/auth"
	"tomozou/usecase"
)

func main() {
	DRIVER := "mysql"
	DSN := "root:@(db:3306)/tomozou?charset=utf8&parseTime=True"
	//DSN := "root:@unix(/cloudsql/ongakuconnection:asia-northeast1:ongkdb)/tomozoudb?charset=utf8&parseTime=True"
	//"ユーザー名:パスワード@unix(/cloudsql/インスタンス接続名)/DB名"

	gormConn, _ := datastore.GormConn(DRIVER, DSN)
	userRepo := userrepoimpl.NewUserRepositoryImpl(gormConn)
	itemRepo := itemrepoimpl.NewItemRepositoryImpl(gormConn)

	useCase := usecase.NewUserProfileApplication(userRepo, itemRepo)

	spotifyHandler := webservice.NewSpotifyHandler(userRepo, itemRepo, gormConn)
	authMiddleware := auth.AuthUser()

	userProfileAppImpl := mainappimpl.UserProfileApplicationImpl{
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

	rAp := r.Group("/apple")
	{
		rAp.GET("/callback", userProfileAppImpl.CallbackApple)
		rAp.GET("/login", userProfileAppImpl.LoginApple)
		rAp.GET("/myartist", userProfileAppImpl.MyArtist)
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

	rTrk := r.Group("/comment")
	{
		rTrk.GET("/add/track/:trackID", func(c *gin.Context) {
			trackIDString := c.Param("trackID")
			trackID, _ := strconv.Atoi(trackIDString)
			//track := []domain.UserTrackTag{}
			//fmt.Println(trackID)
			c.JSON(200, gin.H{"id": trackID})
		})
		rTrk.POST("/comment/add/track/:trackID", func(c *gin.Context) {
			trackIDString := c.Param("trackID")
			trackID, _ := strconv.Atoi(trackIDString)
			//track := []domain.UserTrackTag{}
			//fmt.Println(trackID)
			c.JSON(200, gin.H{"id": trackID})
		})
	}

	// 開発用: データ確認エンドポイント
	devUserRepo := userrepoimpl.NewDevUserRepo(gormConn)
	rDev := r.Group("/dev")
	{
		rDev.GET("/user", func(c *gin.Context) {
			users, _ := devUserRepo.CheckUser()
			c.JSON(200, users)
		})
		rDev.GET("/userartisttag", func(c *gin.Context) {
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
		rDev.GET("/usertracktag", func(c *gin.Context) {
			tags := []domain.UserTrackTag{}
			devUserRepo.DB.Find(&tags)
			c.JSON(200, tags)
		})
		rDev.GET("/tracktagresp", func(c *gin.Context) {
			track := []domain.UserTrackTag{}
			//devUserRepo.DB.Raw("SELECT * FROM user_track_tag JOIN track ON user_track_tag.track_id = track.id").scan()
			c.JSON(200, track)
		})
		rDev.GET("/trackjoin", func(c *gin.Context) {
			//track := []domain.UserTrackTag{}
			track := domain.TrackResp{}
			devUserRepo.DB.Raw("SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id").Scan(&track)
			//devUserRepo.DB.Raw("SELECT * FROM user_track_tags JOIN users ON user_track_tags.user_id = users.id").Scan(&track)
			c.JSON(200, track)
		})
		/*
			rDev.GET("/trackcomment", func(c *gin.Context) {
				trackCommentFull := []domain.TrackCommentFull{}
				devUserRepo.DB.Table("track_comments")
					.Select("track_comment.id, track_comment.user_id,track_comment.track_id,track_comment.comment,track_comment.created_at, track.social_id, track.name, track.artist_name, track.artist_id")
					.JOINS("left join track on track.id = track_comment.track_id").Scan(&trackComment)
				c.JSON(200, trackCommentFull)
			})
			rDev.GET("/tracktag/:userID", func(c *gin.Context) {
				userID := c.Param("userID")
				id, _ := strconv.Atoi(userID)
				track := []domain.UserTrackTag{}
				devUserRepo.DB.Where("user_id = ?", id).Find(&track)
				c.JSON(200, track)
			})
		*/
		rDev.GET("/mytracktag", userProfileAppImpl.DebugTrackTag)
		rDev.GET("/timeline", userProfileAppImpl.TrackTimeLine)
		rDev.GET("/userdata", func(c *gin.Context) {
		})
		rDev.GET("/debug", userProfileAppImpl.Debug)
		rDev.GET("/deptest", func(c *gin.Context) {
			c.String(200, "deploy test")
		})
		rDev.GET("/chat", func(c *gin.Context) {
			chat := []domain.UserChat{}
			devUserRepo.DB.Find(&chat)
			c.JSON(200, chat)
		})
	}

	// Chat 用: authによるJWT 以下から
	chatRepo := chatrepoimpl.NewChatDBRepository(gormConn)
	chatApp := usecase.ChatApplication{
		ItemRepository: itemRepo,
		ChatRepository: chatRepo,
	}
	chatAppImpl := chatappimpl.ChatApplicationImpl{
		UseCase: chatApp,
	}
	rChat := r.Group("/chat")
	rChat.Use(authMiddleware.MiddlewareFunc())
	{
		rChat.GET("/room", chatAppImpl.DisplayChatRoom)
		rChat.POST("/artist/comment", chatAppImpl.ArtistComment)
		rChat.POST("/track/comment", chatAppImpl.TrackComment)
		rChat.POST("/user/chat", chatAppImpl.UserChat)
		rChat.GET("/list/:artistID", chatAppImpl.DisplayChatListByArtist)
	}
	r.Run(":8080")
}
