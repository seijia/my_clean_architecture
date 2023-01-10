package middleware

import (
	"api_client/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ErrorResponse struct {
	Message string `json:"errors"`
}

func JSONErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)
	// printout error string with field
	c.Echo().Logger.Error(err.Error())

	if he, ok := err.(*utils.ErrorResp); ok {
		code = he.StatusCode
		msg = he.ToResponse()
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = map[string]interface{}{"errors": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
