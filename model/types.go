package model

import "net/url"

// Query holds query object that holds uri query params
type Query struct {
	Params *url.Values
}

// FieldQuery holds field expressions with name and values
type FieldQuery struct {
	FieldName   string
	RangeValue  *RangeValue
	SimpleValue *SimpleValue
}

// RangeValue holds 'from' and 'to' range values for field queries
type RangeValue struct {
	RangeStart string
	RangeEnd   string
}

// SimpleValue holds simple values for field queries (equals)
type SimpleValue struct {
	Value string
}

// Location holds position (geo location) data
type Location struct {
	Longitude float64
	Latitude  float64
}

// Location holds airport data (name + location)
type Airport struct {
	Name     string
	Location *Location
}
