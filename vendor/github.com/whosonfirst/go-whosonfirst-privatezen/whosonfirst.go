package privatezen

import (
	"errors"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-placetypes"
	"strconv"
)

func WOFPlaceFromFeature(f geojson.Feature) (*WOFPlace, error) {

	id := whosonfirst.Id(f)

	if id == -1 {
		return nil, errors.New("Invalid WOF ID")
	}

	str_pt := whosonfirst.Placetype(f)

	if !placetypes.IsValidPlacetype(str_pt) {
		return nil, errors.New("Invalid WOF placetype")
	}

	pt, err := placetypes.GetPlacetypeByName(str_pt)

	if err != nil {
		return nil, err
	}

	pl := WOFPlace{
		Id:        id,
		Placetype: pt,
	}

	return &pl, nil
}

type WOFPlace struct {
	Id        int64
	Placetype *placetypes.WOFPlacetype
}

func (w *WOFPlace) String() string {
	return strconv.FormatInt(w.Id, 10)
}
