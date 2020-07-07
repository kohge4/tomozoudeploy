package main

import (
	"context"
	"net/url"
	"reflect"
	"strconv"

	"tomozou/adapter/webservice/appleadapter"
	"tomozou/adapter/webservice/spotifyadapter"
	"tomozou/domain"
	"tomozou/handler/backgroundexec/connectorappimpl"
	"tomozou/handler/chatappimpl"
	"tomozou/handler/mainappimpl"
	"tomozou/infra/datastore"
	"tomozou/infra/datastore/chatrepoimpl"
	"tomozou/infra/datastore/itemchildrepoimpl"
	"tomozou/infra/datastore/itemrepoimpl"
	"tomozou/infra/datastore/userrepoimpl"
	"tomozou/middleware/auth"
	"tomozou/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	applemusic "github.com/kohge4/go-apple-music-sdk"
	"github.com/kohge4/spotify"
)

func main() {
	DRIVER := "mysql"
	//DSN := "root:@(db:3306)/tomozou?charset=utf8&parseTime=True"
	DSN := "root:@unix(/cloudsql/ongakuconnection:asia-northeast1:ongkdb)/tomozoudb?charset=utf8&parseTime=True"
	//"ユーザー名:パスワード@unix(/cloudsql/インスタンス接続名)/DB名"

	gormConn, _ := datastore.GormConn(DRIVER, DSN)
	userRepo := userrepoimpl.NewUserRepositoryImpl(gormConn)
	itemRepo := itemrepoimpl.NewItemRepositoryImpl(gormConn)

	itemChildRepo := itemchildrepoimpl.NewItemChildRepositoryImpl(gormConn)

	useCase := usecase.NewUserProfileApplication(userRepo, itemRepo, itemChildRepo)

	spotifyHandler := spotifyadapter.NewSpotifyHandler(userRepo, itemRepo, gormConn)
	authMiddleware := auth.AuthUser()
	appleHandler := appleadapter.NewAppleHandlerByConfigToken(gormConn, nil, itemChildRepo)

	userProfileAppImpl := mainappimpl.UserProfileApplicationImpl{
		UseCase: useCase,

		Handler:        spotifyHandler,
		AuthMiddleware: authMiddleware,
		Connector:      appleHandler,
	}

	// =====諸説あり ========= connectorAppImpl 多分いらない
	connectorAppImpl := connectorappimpl.ConnectorApplicationImpl{
		AppleHandler:        appleHandler,
		ItemRepository:      itemRepo,
		ItemChildRepository: itemChildRepo,
	}
	// ======================================

	r := gin.Default()

	crs := cors.DefaultConfig()
	crs.AllowOrigins = []string{"http://localhost:8080", "https://ongakuconnection.com", "https://ongakuconnectionsns.firebaseapp.com"}
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
		rTrk.POST("/trackcomment/add", userProfileAppImpl.AddTrackComment)
		rTrk.GET("/get/trackcomment/:trackID", userProfileAppImpl.GetTrackCommentWithUserByTrackID)
	}

	rConn := r.Group("/connector")
	{
		rConn.GET("/apple/:trackID", connectorAppImpl.CreateAppleTrackWebServiceTagByTrackID)
	}

	// 開発用: データ確認エンドポイント
	devUserRepo := userrepoimpl.NewDevUserRepo(gormConn)
	devItemChildRepo := itemchildrepoimpl.NewItemChildRepositoryImpl(gormConn)
	rDev := r.Group("/dev")
	{
		// テーブル簡易確認
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
		rDev.GET("/trackcomment", func(c *gin.Context) {
			trackComment := []domain.TrackComment{}
			devUserRepo.DB.Find(&trackComment)
			c.JSON(200, trackComment)
		})
		rDev.GET("/uttag", func(c *gin.Context) {
			//uttag, _ := itemRepo.ReadUserTrackTagByUserIDANDTagName(2, "nowplaying")
			userTrackTags := []domain.UserTrackTagFull{}
			sql := "SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id WHERE user_track_tags.user_id = ? AND user_track_tags.tag_name = ?"
			//sql := "SELECT * FROM user_track_tags JOIN tracks ON user_track_tags.track_id = tracks.id JOIN users ON user_track_tags.user_id = users.id WHERE user_track_tags.user_id = ? AND user_track_tags.tag_name = ?"
			devUserRepo.DB.Raw(sql, 2, "nowplaying").Scan(&userTrackTags)
			c.JSON(200, userTrackTags)
		})
		rDev.POST("/addtrackcomment", userProfileAppImpl.AddTrackComment)
		rDev.GET("/gettrackcomment/:trackID", userProfileAppImpl.GetTrackCommentWithUserByTrackID)
		rDev.GET("/track/artists/:trackID", func(c *gin.Context) {
			trackIDString := c.Param("trackID")
			trackID, _ := strconv.Atoi(trackIDString)
			t, _ := itemRepo.ReadTrackWithArtistListByTrackID(trackID)
			c.JSON(200, t)
		})
		rDev.GET("/atag", func(c *gin.Context) {
			t := []domain.ArtistWebServiceTag{}
			devUserRepo.DB.Find(&t)
			c.JSON(200, t)
		})
		//======SPOTIFY API 確認
		rDev.GET("/spotify/toptrack", func(c *gin.Context) {
			timerange := "short"
			limit := 5
			opt := &spotify.Options{
				Timerange: &timerange,
				Limit:     &limit,
			}
			results, err := userProfileAppImpl.Handler.Client.GetUserTopTracks2Opt(opt)
			if err != nil {
				c.String(401, err.Error())
				return
			}
			c.JSON(200, results)
		})

		// ===================== apple connector 用
		rDev.GET("/apple/tracktag", func(c *gin.Context) {
			tag := []domain.TrackWebServiceTag{}
			devUserRepo.DB.Find(&tag)
			c.JSON(200, tag)
		})
		rDev.GET("/apple/connector", connectorAppImpl.AppleConnectorByTrack)
		rDev.GET("/ap/search/:word", func(c *gin.Context) {
			//word := c.Param("word")
			// https://api.music.apple.com/v1/catalog/jp/search?term=cero+orphans&types=songs
			// term=cero%2Borphans&types=songs
			searchOpt := &applemusic.SearchOptions{
				Term: "james+bro",
				//Types: "songs",
			}
			ctx := context.Background()
			resp, _, err := appleHandler.Client.Catalog.Search(ctx, "jp", searchOpt)
			if err != nil {
				c.String(401, err.Error())
			}
			c.JSON(200, resp)
		})
		rDev.GET("/apple/tagjoin", func(c *gin.Context) {
			tags, _ := devItemChildRepo.ReadTrackWithTrackWebServiceTagByTrackID(1)
			c.JSON(200, tags)
		})
		rDev.GET("/apple/conn/createtag/:trackID", connectorAppImpl.CreateAppleTrackWebServiceTagByTrackID)
		rDev.GET("/show/apple/:trackID", userProfileAppImpl.ShowAppleMusic)
		rDev.GET("/apple/trackwebservicetag", func(c *gin.Context) {
			tags, _ := devItemChildRepo.ReadTrackWithTrackWebServiceTagByTrackID(3)
			c.JSON(200, tags)
		})
		// ========================

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

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}
	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}
	u.RawQuery = qs.Encode()
	return u.String(), nil
}
