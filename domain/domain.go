package domain

const (
	AuthTokenName = "X-Blot-Auth-Token"
)

type ErrorResponse struct {
	ErrorCode    int
	ErrorMessage string
}

type GetTokenRequest struct {
	Login string
}

type GetTokenResponse struct {
	Code string
}

type ConfirmRequest struct {
	Login, Code string
}

type AddLinkRequest struct {
	Link string
}

type AddNotifyRequest struct {
	Link, Duration string
}
