package blot

import (
	"errors"
	"strings"

	"github.com/nikonor/blot/domain"
)

type Blot struct {
	repo   Repo
	sender Sender
}

type Sender interface {
	Send(login string, typeOfMsg string, values map[string]string) error
}

type Repo interface {
	GetToken(login string) (string, error)          // code
	ConfirmUser(login, code string) (string, error) // token
	AddLink(token, link string) error
	AddNotify(token, duration, link string) error
}

func NewBlot(repo Repo, sender Sender) *Blot {
	return &Blot{repo: repo, sender: sender}
}

// GetToken
// @Tags public
// @Summary создание или обновление токена
// @Accept json
// @Produce json
// @Param body body domain.GetTokenRequest true "Тело запроса"
// @Success 200 {object} domain.GetTokenResponse
// @Failure 500 {object} domain.ErrorResponse "Внутрення ошибка"
// @Router /get_token [POST]
func (b Blot) GetToken(login string) error {
	switch {
	case len(login) == 0:
		return errors.New("login is required")
	case !strings.Contains(login, "@"):
		return errors.New("login must be login")
	}

	code, err := b.repo.GetToken(login)
	if err != nil {
		return err
	}

	if err = b.sender.Send(login, "sendCode::"+code, nil); err != nil {
		return err
	}

	return err
}

// Confirm
// @Tags public
// @Summary подтверждение владения адресом
// @Accept json
// @Produce json
// @Param body body domain.ConfirmRequest true "Тело запроса"
// @Success 200
// @Failure 500 {object} domain.ErrorResponse "Внутрення ошибка"
// @Router /confirm [POST]
func (b Blot) Confirm(login, code string) (string, error) {
	switch {
	case len(login) == 0:
		return "", errors.New("login is required")
	case len(code) == 0:
		return "", errors.New("code is required")
	case !strings.Contains(login, "@"):
		return "", errors.New("login must be login")
	}

	token, err := b.repo.ConfirmUser(login, code)
	if err != nil {
		return "", err
	}

	return token, nil
}

// AddLink
// @Tags public
// @Summary добавление ссылки
// @Accept json
// @Produce json
// @Param body body domain.AddLinkRequest true "Тело запроса"
// @Success 200
// @Failure 500 {object} domain.ErrorResponse "Внутрення ошибка"
// @Router /add_link [POST]
func (b Blot) AddLink(token string, link string) error {
	switch {
	case len(token) == 0:
		return errors.New("token " + domain.AuthTokenName + " is required")
	case len(link) == 0:
		return errors.New("link is required")
	}

	if err := b.repo.AddLink(token, link); err != nil {
		return err
	}

	return nil
}

// AddNotify
// @Tags public
// @Summary подтверждение
// @Accept json
// @Produce json
// @Param body body domain.AddNotifyRequest true "Тело запроса"
// @Success 200
// @Failure 500 {object} domain.ErrorResponse "Внутрення ошибка"
// @Router /add_notify [POST]
func (b Blot) AddNotify(token string, link string, duration string) error {
	switch {
	case len(token) == 0:
		return errors.New("token " + domain.AuthTokenName + " is required")
	case len(duration) == 0:
		return errors.New("duration is required")
	case len(link) == 0:
		return errors.New("link is required")
	}

	if err := b.repo.AddNotify(token, duration, link); err != nil {
		return err
	}

	return nil
}
