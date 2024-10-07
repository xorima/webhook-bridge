package app

import (
	"github.com/xorima/hmacvalidator"
	"github.com/xorima/slogger"
	"io"
	"log/slog"
	"net/http"
)

type HmacConfig interface {
	HmacSecret() string
	HmacEnabled() bool
}

type AuthHmac struct {
	log  *slog.Logger
	cfg  HmacConfig
	hash hmacvalidator.Hash
}

func NewAuthHmacMiddleware(log *slog.Logger, cfg HmacConfig) *AuthHmac {
	return &AuthHmac{
		log:  slogger.SubLogger(log, "auth-hmac"),
		cfg:  cfg,
		hash: hmacvalidator.HashSha256,
	}
}

func (a *AuthHmac) AuthHmacMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.cfg.HmacEnabled() {
			next.ServeHTTP(w, r)
		}
		validator := hmacvalidator.NewHMACValidator(hmacvalidator.HashSha256, a.cfg.HmacSecret())

		hmac := r.Header.Get("X-Hub-Signature-256")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			a.log.WarnContext(r.Context(), "unable to parse body", slogger.ErrorAttr(err))
			NewResponse(400, "Bad Request").WriteResponse(w)
		}
		if validator.IsInvalid(b, hmac) {
			NewResponse(401, "Unauthorized").WriteResponse(w)
			return
		}
		// If authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
