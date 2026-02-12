package vallaris

import (
	"context"
	"encoding/json"
	"net/url"
	"workshop-pwa-api/model"
)

func (v *VallarisAPI) Intersection(collectionID string, coordinate model.Coordinates, excludeIds ...string) (bool, error) {
	var (
		ctx = context.TODO()
		wkt = coordinate.ToWKT()
	)

	query := url.Values{}
	query.Add("intersects", wkt)
	query.Add("limit", "1") // มีวิธีที่ดีกว่านี้
	reader, err := v.doGetRequest(ctx, []string{"collections", collectionID, "items"}, query)
	if err != nil {
		return false, err
	}
	defer reader.Close()

	var data model.FeatureCollection[map[string]any]
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return false, err
	}

	return len(data.Features) > 0, nil
}
