package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
)

func ErrorHook(c *gin.Context, e error) (int, interface{}) {
	code, msg := 500, http.StatusText(http.StatusInternalServerError)

	if _, ok := e.(tonic.BindError); ok {
		code, msg = 400, e.Error()
	} else {
		switch {
		case errors.IsBadRequest(e), errors.IsNotValid(e), errors.IsNotSupported(e), errors.IsNotAssigned(e), errors.IsNotProvisioned(e):
			code, msg = 400, e.Error()
		case errors.IsForbidden(e):
			code, msg = 403, e.Error()
		case errors.IsMethodNotAllowed(e):
			code, msg = 405, e.Error()
		case errors.IsNotFound(e), errors.IsUserNotFound(e):
			code, msg = 404, e.Error()
		case errors.IsUnauthorized(e):
			code, msg = 401, e.Error()
		case errors.IsAlreadyExists(e):
			code, msg = 409, e.Error()
		case errors.IsNotImplemented(e):
			code, msg = 501, e.Error()
		}
	}
	return code, gin.H{`error`: msg}
}
