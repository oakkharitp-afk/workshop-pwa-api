package vallaris

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
	"workshop-pwa-api/model"
)

type Topofeature model.Feature[map[string]any, any]

func (v *VallarisAPI) Intersection(
	collectionID string,
	coordinate model.Coordinates,
	excludeIds ...string,
) (bool, error) {

	ctx := context.TODO()
	wkt := coordinate.ToWKT()

	// We need more than 1 result if we are excluding some IDs.
	limit := 1 + len(excludeIds)

	query := url.Values{}
	query.Add("intersects", wkt)
	query.Add("limit", fmt.Sprint(limit))

	reader, err := v.doGetRequest(ctx, []string{"collections", collectionID, "items"}, query)
	if err != nil {
		return false, err
	}
	defer reader.Close()

	var fc model.FeatureCollection[Topofeature]
	if err := json.NewDecoder(reader).Decode(&fc); err != nil {
		return false, err
	}

	for _, f := range fc.Features {
		// Skip excluded IDs (for update case, avoid checking itself)
		if slices.Contains(excludeIds, f.ID) {
			continue
		}
		return true, nil
	}

	return false, nil
}
