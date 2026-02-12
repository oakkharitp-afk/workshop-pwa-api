package vallaris

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"workshop-pwa-api/model"

	"github.com/labstack/echo/v4"
)

const (
	versionRead  = "1.1"
	versionWrite = "1.1-beta"
)

type VallarisAPI struct {
	baseURLRead  string
	baseURLWrite string
	apiKey       string
	client       *http.Client
}

func debugIo(io io.ReadCloser) {
	var data map[string]any
	json.NewDecoder(io).Decode(&data)
	b, _ := json.Marshal(data)
	fmt.Println("debug:", string(b))
}

func NewVallarisAPI(baseURL, apiKey string, client *http.Client) *VallarisAPI {
	if client == nil {
		client = &http.Client{}
	}

	return &VallarisAPI{
		apiKey:       apiKey,
		client:       client,
		baseURLRead:  baseURL + "/" + versionRead,
		baseURLWrite: baseURL + "/" + versionWrite,
	}
}

func (v *VallarisAPI) doGetRequest(
	ctx context.Context,
	pathSegments []string,
	query url.Values,
) (io.ReadCloser, error) {

	base, err := url.Parse(v.baseURLRead)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	base = base.JoinPath(pathSegments...)

	q := base.Query()
	q.Set("api_key", v.apiKey)

	for key, vals := range query {
		if key == "api_key" {
			continue
		}
		for _, val := range vals {
			q.Add(key, val)
		}
	}

	base.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call vallaris api: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (v *VallarisAPI) GetCollections(ctx context.Context, query url.Values) (io.ReadCloser, error) {
	return v.doGetRequest(ctx, []string{"collections"}, query)
}

func (v *VallarisAPI) GetCollection(ctx context.Context, collectionId string) (io.ReadCloser, error) {
	return v.doGetRequest(ctx, []string{"collections", collectionId}, nil)
}

func (v *VallarisAPI) GetFeatures(ctx context.Context, collectionId string, query url.Values) (io.ReadCloser, error) {
	return v.doGetRequest(ctx, []string{"collections", collectionId, "items"}, query)
}

func (v *VallarisAPI) GetFeature(ctx context.Context, collectionId, featureID string) (io.ReadCloser, error) {
	return v.doGetRequest(ctx, []string{"collections", collectionId, "items", featureID}, nil)
}

func (v *VallarisAPI) CreateFeatures(ctx context.Context, collectionId string, features []any) (io.ReadCloser, error) {
	base, err := url.Parse(v.baseURLWrite)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	base = base.JoinPath("collections", collectionId, "items")

	q := base.Query()
	q.Set("api_key", v.apiKey)
	base.RawQuery = q.Encode()

	fc := model.FeatureCollection[any]{}
	fc.Type = "FeatureCollection"
	fc.Features = features
	raw, err := json.Marshal(fc)
	if err != nil {
		return nil, fmt.Errorf("call vallaris api: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base.String(), bytes.NewBuffer(raw))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call vallaris api: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		defer resp.Body.Close()
		// debugIo(resp.Body)
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (v *VallarisAPI) UpdateFeature(ctx context.Context, collectionId string, featureId string, feature any) (io.ReadCloser, error) {
	base, err := url.Parse(v.baseURLWrite)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	base = base.JoinPath("collections", collectionId, "items", featureId)

	q := base.Query()
	q.Set("api_key", v.apiKey)
	base.RawQuery = q.Encode()

	raw, err := json.Marshal(feature)
	if err != nil {
		return nil, fmt.Errorf("call vallaris api: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, base.String(), bytes.NewBuffer(raw))
	if err != nil {
		return nil, fmt.Errorf("update request: %w", err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call vallaris api: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		defer resp.Body.Close()
		// debugIo(resp.Body)
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (v *VallarisAPI) DeleteFeature(ctx context.Context, collectionId string, featureId string) error {
	base, err := url.Parse(v.baseURLWrite)
	if err != nil {
		return fmt.Errorf("invalid base url: %w", err)
	}

	base = base.JoinPath("collections", collectionId, "items", featureId)

	q := base.Query()
	q.Set("api_key", v.apiKey)
	base.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, base.String(), nil)
	if err != nil {
		return fmt.Errorf("delete request: %w", err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := v.client.Do(req)
	if err != nil {
		return fmt.Errorf("call vallaris api: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		defer resp.Body.Close()
		// debugIo(resp.Body)
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}
