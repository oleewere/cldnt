package model

import (
	"fmt"
	"net/url"
)

func CreateQuery(fieldsQueryExpression string) *Query {
	params := url.Values{}
	queryString := "*:*"
	if len(fieldsQueryExpression) > 0 {
		queryString = fieldsQueryExpression
	}
	params.Add("q", queryString)
	query := Query{Params: &params}
	return &query
}

func (q *Query) AddDistanceSort(location *Location) {
	q.Params.Add("sort", fmt.Sprintf("\"<distance,lon,lat,%f,%f,km>\"", location.Longitude, location.Latitude))
}

func (f *FieldQuery) ToQueryString() string {
	var result string
	if f.RangeValue != nil {
		rangeStartValue := f.RangeValue.RangeStart
		rangeEndValue := f.RangeValue.RangeStart
		result = fmt.Sprintf("%v:[%v TO %v]", f.FieldName, wildardIfEmpty(rangeStartValue), wildardIfEmpty(rangeEndValue))
	} else if f.SimpleValue != nil {
		result = fmt.Sprintf("%v:%v", f.FieldName, wildardIfEmpty(f.SimpleValue.Value))
	}
	return result
}

func (f *FieldQuery) And(field *FieldQuery) string {
	return fmt.Sprintf("[%v AND %v]", f.ToQueryString(), field.ToQueryString())
}

func (f *FieldQuery) Or(field *FieldQuery) string {
	return fmt.Sprintf("[%v OR %v]", f.ToQueryString(), field.ToQueryString())
}

func (f *FieldQuery) Negate() string {
	return fmt.Sprintf("NOT %v", f.ToQueryString())
}

func wildardIfEmpty(fieldValue string) string {
	if len(fieldValue) == 0 {
		fieldValue = "*"
	}
	return fieldValue
}
