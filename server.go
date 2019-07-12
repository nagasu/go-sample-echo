package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		ss := allUsers()
		return c.JSON(http.StatusOK, ss)
	})

	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/show", show)

	e.Logger.Fatal(e.Start(":8000"))
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
