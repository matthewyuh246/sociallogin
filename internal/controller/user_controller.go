package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-oidc"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/matthewyuh246/socallogin/internal/domain"
	"github.com/matthewyuh246/socallogin/internal/usecase"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type IUserController interface {
	Authentication(c echo.Context) error
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
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

func (uc *userController) SignUp(c echo.Context) error {
	ctx := context.Background()
	oauth2Token, err := config.Exchange(ctx, c.QueryParam("code"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "No id_token field in oauth2 token.")
	}

	idToken, err := provider.Verifier(&oidc.Config{ClientID: clientID}).Verify(ctx, rawIDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to verify ID Token: "+err.Error())
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.JSON(http.StatusInternalServerError, "Faile to get user profile: "+err.Error())
	}

	if domain, ok := profile["hd"].(string); !ok || domain != "g.kogakuin.jp" {
		return c.JSON(http.StatusUnauthorized, "Unauthorized domain")
	}

	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Email = profile["email"].(string)

	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)

}

func (uc *userController) LogIn(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
