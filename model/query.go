package model

import (
	"fmt"
	"net/url"
)

// CreateQuery create query from fields expression
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

// AddDistanceSort extend query with distance sorting
func (q *Query) AddDistanceSort(location *Location) *Query {
	if location != nil && q.Params != nil {
		q.Params.Add("sort", fmt.Sprintf("\"<distance,lon,lat,%f,%f,km>\"", location.Longitude, location.Latitude))
	}
	return q
}

// AddLimit extend query with row limits
func (q *Query) AddLimit(rows int) *Query {
	if q.Params != nil {
		q.Params.Add("limit", fmt.Sprintf("%d", rows))
	}
	return q
}

// ToQueryString generate a string from a field query object
func (f *FieldQuery) ToQueryString() string {
	var result string
	if f.RangeValue != nil {
		rangeStartValue := f.RangeValue.RangeStart
		rangeEndValue := f.RangeValue.RangeEnd
		result = fmt.Sprintf("%v:[%v TO %v]", f.FieldName, wildardIfEmpty(rangeStartValue), wildardIfEmpty(rangeEndValue))
	} else if f.SimpleValue != nil {
		result = fmt.Sprintf("%v:%v", f.FieldName, wildardIfEmpty(f.SimpleValue.Value))
	}
	return result
}

// And provide 'and' logical expression for field query
func (f *FieldQuery) And(field *FieldQuery) string {
	if field == nil {
		return f.ToQueryString()
	}
	fQuery := f.ToQueryString()
	inputFQuery := field.ToQueryString()
	if len(inputFQuery) == 0 {
		return f.ToQueryString()
	}
	if len(fQuery) == 0 {
		return field.ToQueryString()
	}
	return fmt.Sprintf("[%v AND %v]", f.ToQueryString(), field.ToQueryString())
}

// Or provide 'or' logical expression for field query
func (f *FieldQuery) Or(field *FieldQuery) string {
	if field == nil {
		return f.ToQueryString()
	}
	fQuery := f.ToQueryString()
	inputFQuery := field.ToQueryString()
	if len(inputFQuery) == 0 {
		return f.ToQueryString()
	}
	if len(fQuery) == 0 {
		return field.ToQueryString()
	}
	return fmt.Sprintf("[%v OR %v]", f.ToQueryString(), field.ToQueryString())
}

// Negate provide 'negate' logical expression for field query
func (f *FieldQuery) Negate() string {
	if len(f.FieldName) == 0 {
		return f.ToQueryString()
	}
	if f.RangeValue == nil && f.SimpleValue == nil {
		f.SimpleValue = &SimpleValue{}
	}
	return fmt.Sprintf("NOT %v", f.ToQueryString())
}

func wildardIfEmpty(fieldValue string) string {
	if len(fieldValue) == 0 {
		fieldValue = "*"
	}
	return fieldValue
}
