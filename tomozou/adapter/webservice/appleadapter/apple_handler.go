package appleadapter

import (
	"tomozou/domain"

	"github.com/jinzhu/gorm"
	applemusic "github.com/kohge4/go-apple-music-sdk"
)

type AppleHandler struct {
	Client *applemusic.Client
	DB     *gorm.DB
	Config *WebServiceConfig

	UserRepository      domain.UserRepository
	ItemRepository      domain.ItemRepository
	ItemChildRepository domain.ItemChildRepository
}

type WebServiceConfig struct {
	ClientID    string
	SecretKey   string
	RedirectURL string
	State       string
	Token       *WebAPIToken
}

type WebAPIToken struct {
	AccessToken  string
	RefreshToken string
}

func NewAppleHandlerByConfigToken(db *gorm.DB, config *WebServiceConfig, itemChildRepository domain.ItemChildRepository) *AppleHandler {
	//ctx := context.Background
	if config == nil {
		config = NewWebServiceConfig()
	}
	tp := applemusic.Transport{Token: config.Token.AccessToken}
	client := applemusic.NewClient(tp.Client())
	return &AppleHandler{
		Client:              client,
		DB:                  db,
		ItemChildRepository: itemChildRepository,
	}
}

func (h *AppleHandler) ConvertWebServiceAccountImpl() {}

func (h *AppleHandler) User() {}

func (h *AppleHandler) SaveUserItem() {}

func (h *AppleHandler) UpdateUserItem() {}

func (h *AppleHandler) UpdateUserItemOpt() {}

func (h *AppleHandler) DebugItem() {}
