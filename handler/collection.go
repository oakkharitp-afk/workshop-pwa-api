package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCollections(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		queries = c.QueryParams()
	)

	respBody, err := h.vaAPI.GetCollections(ctx, queries)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	defer respBody.Close()

	return echoStream(c, respBody)
}

func (h *handler) GetCollection(c echo.Context) error {
	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	respBody, err := h.vaAPI.GetCollection(ctx, collectionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	defer respBody.Close()

	return echoStream(c, respBody)
}
