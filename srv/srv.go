package srv

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/nikonor/blot/domain"
)

type Blot interface {
	GetToken(login string) error
	Confirm(login, code string) (string, error)
	AddLink(link string) error
	AddNotify(link string, duration string) error
}

type Response struct {
	Status string
	Token  string `json:"token,omitempty"`
}

func NewSrv(b Blot, port string) {
	e := echo.New()

	e.POST("/get_token", func(c echo.Context) error {
		var req domain.GetTokenRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		if err := b.GetToken(req.Login); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})

	e.POST("/confirm", func(c echo.Context) error {
		var req domain.ConfirmRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		token, err := b.Confirm(req.Login, req.Code)
		if err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok", Token: token})
	})
	e.POST("/add_link", func(c echo.Context) error {
		var req domain.AddLinkRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		if err := b.AddLink(req.Link); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		return c.JSON(200, Response{Status: "Ok"})
	})
	e.POST("/add_notify", func(c echo.Context) error {
		var req domain.AddNotifyRequest
		if err := getRequest(c.Request().Body, &req); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
		}
		if err := b.AddNotify(req.Link, req.Duration); err != nil {
			return c.JSON(500, domain.ErrorResponse{ErrorCode: 500, ErrorMessage: err.Error()})
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
