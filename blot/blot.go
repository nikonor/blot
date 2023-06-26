package blot

type Blot struct {
}

func NewBlot() *Blot {
	return &Blot{}
}

type RegisterRequest struct {
	Login, Password, RePassword string
}

func (b Blot) Register(login, password, rePassword string) error {
	// TODO implement me
	panic("implement me")
}

type ConfirmRequest struct {
	Login, Code string
}

func (b Blot) Confirm(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

type ResurrectTokenRequest struct {
	Login string
}

func (b Blot) ResurrectToken(login string) error {
	// TODO implement me
	panic("implement me")
}

type ConfirmResurrectRequest struct {
	Login, Code string
}

func (b Blot) ConfirmResurrect(login, code string) (string, error) {
	// TODO implement me
	panic("implement me")
}

type AddLinkRequest struct {
	Link string
}

func (b Blot) AddLink(link string) error {
	// TODO implement me
	panic("implement me")
}

type AddNotifyRequest struct {
	Link, Duration string
}

func (b Blot) AddNotify(link string, duration string) error {
	// TODO implement me
	panic("implement me")
}
