package model

import (
	"time"
)

type FeatureCollection[Feature any] struct {
	NumberMatched  *int      `json:"numberMatched,omitempty"`
	NumberReturned *int      `json:"numberReturned,omitempty"`
	Type           string    `json:"type"`
	Features       []Feature `json:"features"`
	// Links          []Link     `json:"links,omitempty"`
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
}

type Feature[Properties any, Coord any] struct {
	ID         string          `json:"id,omitempty"`
	Type       string          `json:"type"`
	Geometry   Geometry[Coord] `json:"geometry"`
	Properties Properties      `json:"properties"`
}

type Link struct {
	Href      string `json:"href" bson:"href"`
	Rel       string `json:"rel,omitempty" bson:"rel,omitempty"`
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	Hreflang  string `json:"hreflang,omitempty" bson:"hreflang,omitempty"`
	Title     string `json:"title,omitempty" bson:"title,omitempty"`
	Length    *int   `json:"length,omitempty" bson:"length,omitempty"`
	Templated *bool  `json:"templated,omitempty" bson:"templated,omitempty"`
}

// dma boundary
type DMABoundaryFeature Feature[DMABoundary, Polygon]
type DMABoundaryFeatureCollection FeatureCollection[DMABoundaryFeature]

// flow meter
type FlowMeterFeature Feature[FlowMeter, Point]
type FlowMeterFeatureCollection FeatureCollection[FlowMeterFeature]

// step test
type StepTestFeature Feature[StepTestProp, Polygon]
type StepTestFeatureCollection FeatureCollection[StepTestFeature]

type DMABoundary struct {
	ID  string `json:"_id,omitempty"`
	ID0 string `json:"id,omitempty"`

	DmaID      int        `json:"dma_id"`
	Dmaname    string     `json:"dmaname"`
	Dmano      string     `json:"dmano"`
	Globalid   string     `json:"globalid"`
	Loggerid   int        `json:"loggerid"`
	Mmno       any        `json:"mmno"`
	Pwacode    string     `json:"pwacode"`
	Recorddate *time.Time `json:"recorddate"`
	Remark     string     `json:"remark"`

	CreatedAt *time.Time `json:"_createdAt,omitempty"`
	CreatedBy *string    `json:"_createdBy,omitempty"`
	Createdat *time.Time `json:"_createdat,omitempty"`
	Createdby *string    `json:"_createdby,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`
	UpdatedBy *string    `json:"_updatedBy,omitempty"`
	Updatedat *time.Time `json:"_updatedat,omitempty"`
	Updatedby *string    `json:"_updatedby,omitempty"`
}

type FlowMeter struct {
	ID  string `json:"_id,omitempty"`
	ID0 string `json:"id,omitempty"`

	Brandcode        string    `json:"brandcode"`
	Flowmetertype    int       `json:"flowmetertype"`
	Globalid         string    `json:"globalid"`
	Inputflowchannel int       `json:"inputflowchannel"`
	Installeddate    time.Time `json:"installeddate"`
	Loggerid         int       `json:"loggerid"`
	Measuretype      int       `json:"measuretype"`
	MeterID          int       `json:"meter_id"`
	Metersize        string    `json:"metersize"`
	Model            string    `json:"model"`
	Pipesize         int       `json:"pipesize"`
	Pipetype         string    `json:"pipetype"`
	Pwacode          string    `json:"pwacode"`
	Recorddate       time.Time `json:"recorddate"`
	Remark           string    `json:"remark"`

	CreatedAt *time.Time `json:"_createdAt,omitempty"`
	CreatedBy *string    `json:"_createdBy,omitempty"`
	Createdat *time.Time `json:"_createdat,omitempty"`
	Createdby *string    `json:"_createdby,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`
	UpdatedBy *string    `json:"_updatedBy,omitempty"`
	Updatedat *time.Time `json:"_updatedat,omitempty"`
	Updatedby *string    `json:"_updatedby,omitempty"`
}

type StepTestProp struct {
	ID  string `json:"_id,omitempty"`
	ID0 string `json:"id,omitempty"`

	Dmano      string    `json:"dmano"`
	Globalid   string    `json:"globalid"`
	Jobstepid  any       `json:"jobstepid"`
	Jobstepno  any       `json:"jobstepno"`
	Pwacode    string    `json:"pwacode"`
	Recorddate time.Time `json:"recorddate"`
	Remark     string    `json:"remark"`
	StepID     int       `json:"step_id"`
	Stepname   string    `json:"stepname"`
	Stepno     string    `json:"stepno"`

	CreatedAt *time.Time `json:"_createdAt,omitempty"`
	CreatedBy *string    `json:"_createdBy,omitempty"`
	Createdat *time.Time `json:"_createdat,omitempty"`
	Createdby *string    `json:"_createdby,omitempty"`
	UpdatedAt *time.Time `json:"_updatedAt,omitempty"`
	UpdatedBy *string    `json:"_updatedBy,omitempty"`
	Updatedat *time.Time `json:"_updatedat,omitempty"`
	Updatedby *string    `json:"_updatedby,omitempty"`
}
