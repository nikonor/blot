package srv

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/nikonor/blot/blot"
)

type Blot interface {
	Register(login, password, rePassword string) error
	Confirm(login, code string) (string, error)
	ResurrectToken(login string) error
	ConfirmResurrect(login, code string) (string, error)
	AddLink(link string) error
	AddNotify(link string, duration string) error
}

type Response struct {
	Status string
	Token  string `json:"token,omitempty"`
}

func NewSrv(b *blot.Blot, port string) {
	e := echo.New()

	e.POST("/register", func(c echo.Context) error {
		var req blot.RegisterRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		if err := b.Register(req.Login, req.Password, req.RePassword); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})

	e.POST("/confirm", func(c echo.Context) error {
		var req blot.ConfirmRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		token, err := b.Confirm(req.Login, req.Code)
		if err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok", Token: token})
	})
	e.POST("/resurrect_token", func(c echo.Context) error {
		var req blot.ResurrectTokenRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		if err := b.ResurrectToken(req.Login); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})
	e.POST("/confirm_resurrect", func(c echo.Context) error {
		var req blot.ConfirmResurrectRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		token, err := b.ConfirmResurrect(req.Login, req.Code)
		if err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok", Token: token})
	})
	e.POST("/add_link", func(c echo.Context) error {
		var req blot.AddLinkRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		if err := b.AddLink(req.Link); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})
	e.POST("/add_notify", func(c echo.Context) error {
		var req blot.AddNotifyRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		if err := b.AddNotify(req.Link, req.Duration); err != nil {
			return c.JSON(500, Response{Status: "Error::" + err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})

	if err := e.Start(port); err != nil {
		panic(err.Error())
	}
}

func getRequest(r io.ReadCloser, ret any) error {
	body, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &ret); err != nil {
		return err
	}

	return nil
}
