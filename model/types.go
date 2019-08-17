package model

import "net/url"

type Query struct {
	Params *url.Values
}

type FieldQuery struct {
	FieldName   string
	RangeValue  *RangeValue
	SimpleValue *SimpleValue
}

type RangeValue struct {
	RangeStart string
	RangeEnd   string
}

type SimpleValue struct {
	Value string
}

type Location struct {
	Longitude float64
	Latitude  float64
}

type Airport struct {
	Name     string
	Location *Location
}
