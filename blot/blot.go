package blot

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
	AddLink(login, link string) error
	AddNotify(login, duration, link string) error
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
	// TODO implement me
	panic("implement me")
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
	// TODO implement me
	panic("implement me")
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
func (b Blot) AddLink(link string) error {
	// TODO implement me
	panic("implement me")
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
func (b Blot) AddNotify(link string, duration string) error {
	// TODO implement me
	panic("implement me")
}
