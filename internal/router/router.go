package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matthewyuh246/socallogin/internal/controller"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://your-frontend-domain.com"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowCredentials: true,
	}))

	e.GET("/auth", uc.Authentication)
	e.GET("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	return e
}
