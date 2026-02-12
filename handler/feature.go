package handler

import (
	"errors"
	"fmt"
	"net/http"
	"workshop-pwa-api/model"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetFeatures(c echo.Context) error {
	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
		queries      = c.QueryParams()
	)

	respBody, err := h.vaAPI.GetFeatures(ctx, collectionId, queries)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	defer respBody.Close()

	return echoStream(c, respBody)
}

func (h *handler) GetFeature(c echo.Context) error {
	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
		featureId    = c.Param("featureId")
	)

	respBody, err := h.vaAPI.GetFeature(ctx, collectionId, featureId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	defer respBody.Close()

	return echoStream(c, respBody)
}

func (h *handler) CreateFeatures(c echo.Context) error {
	var (
		collectionId = c.Param("collectionId")
		err          error
	)

	switch collectionId {
	case flowMeter:
		err = h.createFlowMeters(c)
	case stepTest:
		err = h.createStepTests(c)
	case dmaBoundary:
		err = h.createDMAs(c)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Unsuported feature collections"))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (h *handler) createDMAs(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.DMABoundaryFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}

	// topo check
	data := make([]any, len(fc.Features))
	for i, feature := range fc.Features {
		intersect, err := h.vaAPI.Intersection(collectionId, feature.Geometry.Coordinates)
		if err != nil {
			return errors.New("process data failed, error to check topology")
		}
		if intersect {
			return fmt.Errorf("feature[%d] not pass topology", i)
		}
		data[i] = feature
	}

	reader, err := h.vaAPI.CreateFeatures(ctx, collectionId, data)
	if err != nil {
		return errors.New("process data failed, error to create data")
	}
	defer reader.Close()

	return echoStream(c, reader)
}

func (h *handler) createFlowMeters(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.FlowMeterFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}

	// topo check
	data := make([]any, len(fc.Features))
	for i, feature := range fc.Features {
		intersect, err := h.vaAPI.Intersection(collectionId, feature.Geometry.Coordinates)
		if err != nil {
			return errors.New("process data failed, error to check topology")
		}
		if intersect {
			return fmt.Errorf("feature[%d] not pass topology", i)
		}
		data[i] = feature
	}

	reader, err := h.vaAPI.CreateFeatures(ctx, collectionId, data)
	if err != nil {
		return errors.New("process data failed, error to create data")
	}
	defer reader.Close()

	return echoStream(c, reader)
}

func (h *handler) createStepTests(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.StepTestFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}

	// topo check
	data := make([]any, len(fc.Features))
	for i, feature := range fc.Features {
		intersect, err := h.vaAPI.Intersection(collectionId, feature.Geometry.Coordinates)
		if err != nil {
			return errors.New("process data failed, error to check topology")
		}
		if intersect {
			return fmt.Errorf("feature[%d] not pass topology", i)
		}
		data[i] = feature
	}

	reader, err := h.vaAPI.CreateFeatures(ctx, collectionId, data)
	if err != nil {
		return errors.New("process data failed, error to create data")
	}
	defer reader.Close()

	return echoStream(c, reader)
}

func (h *handler) UpdateFeatures(c echo.Context) error {
	var (
		collectionId = c.Param("collectionId")
		err          error
	)

	switch collectionId {
	case flowMeter:
		err = h.updateFlowmeters(c)
	case stepTest:
		err = h.updateStepTests(c)
	case dmaBoundary:
		err = h.updateDMAs(c)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Unsuported feature collections"))
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (h *handler) updateDMAs(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.DMABoundaryFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}
	for i, feature := range fc.Features {
		// intersect, err := h.vaAPI.Intersection(collectionId, feature.Geometry.Coordinates)
		// if err != nil {
		// 	return errors.New("process data failed, error to check topology")
		// }
		// if intersect {
		// 	return fmt.Errorf("feature[%d] not pass topology", i)
		// }
		if feature.ID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, id is required", i))
		}
	}
	features := make([]model.DMABoundaryFeature, 0, len(fc.Features))
	for i, feature := range fc.Features {
		reader, err := h.vaAPI.UpdateFeature(ctx, collectionId, feature.ID, feature)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, %w", i, err))
		}
		reader.Close()

		features = append(features, feature)
	}

	fc.Features = features

	return c.JSON(http.StatusOK, fc)
}
func (h *handler) updateStepTests(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.StepTestFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}
	for i, feature := range fc.Features {
		if feature.ID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, id is required", i))
		}
	}
	features := make([]model.StepTestFeature, 0, len(fc.Features))
	for i, feature := range fc.Features {
		reader, err := h.vaAPI.UpdateFeature(ctx, collectionId, feature.ID, feature)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, %w", i, err))
		}
		reader.Close()

		features = append(features, feature)
	}

	fc.Features = features

	return c.JSON(http.StatusOK, fc)
}

func (h *handler) updateFlowmeters(c echo.Context) error {

	var (
		ctx          = c.Request().Context()
		collectionId = c.Param("collectionId")
	)

	var fc model.FlowMeterFeatureCollection
	if err := c.Bind(&fc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if len(fc.Features) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Required at lest 1 item"))
	}
	for i, feature := range fc.Features {
		if feature.ID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, id is required", i))
		}
	}
	features := make([]model.FlowMeterFeature, 0, len(fc.Features))
	for i, feature := range fc.Features {
		reader, err := h.vaAPI.UpdateFeature(ctx, collectionId, feature.ID, feature)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("update feature[%d] failed, %w", i, err))
		}
		reader.Close()

		features = append(features, feature)
	}

	fc.Features = features

	return c.JSON(http.StatusOK, fc)
}

func (h *handler) DeleteFeatures(c echo.Context) error {
	var (
		collectionId = c.Param("collectionId")
		ctx          = c.Request().Context()
	)

	switch collectionId {
	case flowMeter:
		var fc model.FlowMeterFeatureCollection
		if err := c.Bind(&fc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for i, feature := range fc.Features {
			if feature.ID == "" {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, id is required", i))
			}
			if err := h.vaAPI.DeleteFeature(ctx, collectionId, feature.ID); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, %w", i, err))
			}
		}
	case stepTest:
		var fc model.StepTestFeatureCollection
		if err := c.Bind(&fc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for i, feature := range fc.Features {
			if feature.ID == "" {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, id is required", i))
			}
			if err := h.vaAPI.DeleteFeature(ctx, collectionId, feature.ID); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, %w", i, err))
			}
		}
	case dmaBoundary:
		var fc model.DMABoundaryFeatureCollection
		if err := c.Bind(&fc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for i, feature := range fc.Features {
			if feature.ID == "" {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, id is required", i))
			}
			if err := h.vaAPI.DeleteFeature(ctx, collectionId, feature.ID); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("delete feature[%d] failed, %w", i, err))
			}
		}
	default:
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Unsuported feature collections"))
	}
	return c.String(http.StatusNoContent, "")
}
