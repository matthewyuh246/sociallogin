package controller

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/matthewyuh246/socallogin/internal/usecase"
	"github.com/x/oauth2"
	"golang.org/x/oauth2/google"
)

type IUserController interface {
	Authentication(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

var (
	clientID     string
	clientSecret string
	redirectURL  = "http://localhost:8080/signup"
	provider     *oidc.Provider
	config       *oauth2.Config
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	clientID = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	var err error
	provider, err = oidc.NewProvider(context.Background(), "http://accounts.google.com")
	if err != nil {
		log.Fatalf("failed to get provider: %v", err)
	}

	config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     google.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
}

func (uc *userController) Authentication(c echo.Context) error {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}
