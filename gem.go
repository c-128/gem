package gem

import "mime"

type Handler func(ctx *Ctx) error

type Map map[string]any

const (
	maxURILength = 1024
	geminiScheme = "gemini"

	StatusInput          = 10
	StatusSensitiveInput = 11

	StatusSuccess = 20

	StautsTemporaryRedirection = 30
	StatusPermanentRedirection = 31

	StatusTemporaryFailure  = 40
	StatusServerUnavailable = 41
	StautsCGIError          = 42
	StautsProxyError        = 43
	StautsSlowDown          = 44

	StatusPermanentFailure    = 50
	StatusNotFound            = 51
	StatusGone                = 52
	StatusProxyRequestRefused = 53
	StatusBadRequest          = 59

	StatusClientCert              = 60
	StatusClientCertNotAuthorized = 61
	StatusClientCertNotValid      = 62
)

func init() {
	mime.AddExtensionType(".gmi", "text/gemini")
}
