package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func echoStream(c echo.Context, respBody io.ReadCloser) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	_, err := io.Copy(c.Response().Writer, respBody)
	return err
}
