package model

import (
	"net/url"
	"reflect"
	"testing"
)

func TestCreateQuery(t *testing.T) {
	type args struct {
		fieldsQueryExpression string
	}
	emptyParams := url.Values{}
	emptyParams.Add("q", "*:*")

	nonEmptyParams := url.Values{}
	nonEmptyParams.Add("q", "other:otherValue")

	tests := []struct {
		name string
		args args
		want *Query
	}{
		{
			name: "Test with empty field expression",
			args: args{
				fieldsQueryExpression: "",
			},
			want: &Query{Params: &emptyParams},
		},
		{
			name: "Test with non-empty field expression",
			args: args{
				fieldsQueryExpression: "other:otherValue",
			},
			want: &Query{Params: &nonEmptyParams},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateQuery(tt.args.fieldsQueryExpression); !reflect.DeepEqual(got.Params, tt.want.Params) {
				t.Errorf("CreateQuery() = %v, want %v", got.Params, tt.want.Params)
			}
		})
	}
}

func TestQuery_AddDistanceSort(t *testing.T) {
	type fields struct {
		Params *url.Values
	}
	type args struct {
		location *Location
	}

	locationSample := Location{Latitude: 1.0, Longitude: 1.0}
	paramsWithLocation := url.Values{}
	paramsWithLocation.Add("sort", "\"<distance,lon,lat,1.000000,1.000000,km>\"")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			name: "Test without params",
			args: args{
				location: &locationSample,
			},
			fields: fields{
				Params: nil,
			},
			want: &Query{},
		},
		{
			name: "Test without location",
			args: args{
				location: nil,
			},
			fields: fields{
				Params: &url.Values{},
			},
			want: &Query{Params: &url.Values{}},
		},
		{
			name: "Test with location",
			args: args{
				location: &locationSample,
			},
			fields: fields{
				Params: &url.Values{},
			},
			want: &Query{Params: &paramsWithLocation},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				Params: tt.fields.Params,
			}
			if got := q.AddDistanceSort(tt.args.location); !reflect.DeepEqual(got.Params, tt.want.Params) {
				t.Errorf("Query.AddDistanceSort() = %v, want %v", got.Params, tt.want.Params)
			}
		})
	}
}

func TestQuery_AddLimit(t *testing.T) {
	type fields struct {
		Params *url.Values
	}
	type args struct {
		rows int
	}

	defaultParams := url.Values{}
	defaultParams.Add("limit", "0")

	paramsWithRows := url.Values{}
	paramsWithRows.Add("limit", "10")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			name: "Test without params",
			args: args{
				rows: 0,
			},
			fields: fields{
				Params: nil,
			},
			want: &Query{},
		},
		{
			name: "Test with default rows",
			args: args{
				rows: 0,
			},
			fields: fields{
				Params: &url.Values{},
			},
			want: &Query{Params: &defaultParams},
		},
		{
			name: "Test with updated rows",
			args: args{
				rows: 10,
			},
			fields: fields{
				Params: &url.Values{},
			},
			want: &Query{Params: &paramsWithRows},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				Params: tt.fields.Params,
			}
			if got := q.AddLimit(tt.args.rows); !reflect.DeepEqual(got.Params, tt.want.Params) {
				t.Errorf("Query.AddLimit() = %v, want %v", got.Params, tt.want.Params)
			}
		})
	}
}

func TestFieldQuery_ToQueryString(t *testing.T) {
	type fields struct {
		FieldName   string
		RangeValue  *RangeValue
		SimpleValue *SimpleValue
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test with empty fields",
			fields: fields{},
			want:   "",
		},
		{
			name: "Test with only field name",
			fields: fields{
				FieldName: "myField",
			},
			want: "",
		},
		{
			name: "Test with empty field value",
			fields: fields{
				FieldName:   "myField",
				SimpleValue: &SimpleValue{},
			},
			want: "myField:*",
		},
		{
			name: "Test with field value",
			fields: fields{
				FieldName:   "myField",
				SimpleValue: &SimpleValue{Value: "myValue"},
			},
			want: "myField:myValue",
		},
		{
			name: "Test with empty range values",
			fields: fields{
				FieldName:  "myField",
				RangeValue: &RangeValue{},
			},
			want: "myField:[* TO *]",
		},
		{
			name: "Test with only range start",
			fields: fields{
				FieldName:  "myField",
				RangeValue: &RangeValue{RangeStart: "1"},
			},
			want: "myField:[1 TO *]",
		},
		{
			name: "Test with only range end",
			fields: fields{
				FieldName:  "myField",
				RangeValue: &RangeValue{RangeEnd: "10"},
			},
			want: "myField:[* TO 10]",
		},
		{
			name: "Test with range values",
			fields: fields{
				FieldName:  "myField",
				RangeValue: &RangeValue{RangeStart: "1", RangeEnd: "10"},
			},
			want: "myField:[1 TO 10]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FieldQuery{
				FieldName:   tt.fields.FieldName,
				RangeValue:  tt.fields.RangeValue,
				SimpleValue: tt.fields.SimpleValue,
			}
			if got := f.ToQueryString(); got != tt.want {
				t.Errorf("FieldQuery.ToQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldQuery_And(t *testing.T) {
	type fields struct {
		FieldName   string
		RangeValue  *RangeValue
		SimpleValue *SimpleValue
	}
	type args struct {
		field *FieldQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test add on empty field",
			args: args{
				field: &FieldQuery{
					FieldName: "myField",
					SimpleValue: &SimpleValue{
						Value: "myValue",
					},
				},
			},
			fields: fields{},
			want:   "myField:myValue",
		},
		{
			name: "Test on nil field input",
			args: args{
				field: &FieldQuery{},
			},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "myField1:myValue1",
		},
		{
			name: "Test on empty field input",
			args: args{},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "myField1:myValue1",
		},
		{
			name: "Test adding fields",
			args: args{
				field: &FieldQuery{
					FieldName: "myField",
					SimpleValue: &SimpleValue{
						Value: "myValue",
					},
				},
			},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "[myField1:myValue1 AND myField:myValue]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FieldQuery{
				FieldName:   tt.fields.FieldName,
				RangeValue:  tt.fields.RangeValue,
				SimpleValue: tt.fields.SimpleValue,
			}
			if got := f.And(tt.args.field); got != tt.want {
				t.Errorf("FieldQuery.And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldQuery_Or(t *testing.T) {
	type fields struct {
		FieldName   string
		RangeValue  *RangeValue
		SimpleValue *SimpleValue
	}
	type args struct {
		field *FieldQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test or on empty field",
			args: args{
				field: &FieldQuery{
					FieldName: "myField",
					SimpleValue: &SimpleValue{
						Value: "myValue",
					},
				},
			},
			fields: fields{},
			want:   "myField:myValue",
		},
		{
			name: "Test on nil field input",
			args: args{
				field: &FieldQuery{},
			},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "myField1:myValue1",
		},
		{
			name: "Test on empty field input",
			args: args{},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "myField1:myValue1",
		},
		{
			name: "Test or fields",
			args: args{
				field: &FieldQuery{
					FieldName: "myField",
					SimpleValue: &SimpleValue{
						Value: "myValue",
					},
				},
			},
			fields: fields{
				FieldName: "myField1",
				SimpleValue: &SimpleValue{
					Value: "myValue1",
				},
			},
			want: "[myField1:myValue1 OR myField:myValue]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FieldQuery{
				FieldName:   tt.fields.FieldName,
				RangeValue:  tt.fields.RangeValue,
				SimpleValue: tt.fields.SimpleValue,
			}
			if got := f.Or(tt.args.field); got != tt.want {
				t.Errorf("FieldQuery.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldQuery_Negate(t *testing.T) {
	type fields struct {
		FieldName   string
		RangeValue  *RangeValue
		SimpleValue *SimpleValue
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Test without field name",
			fields: fields{},
			want:   "",
		},
		{
			name: "Test negate without any value",
			fields: fields{
				FieldName: "myField",
			},
			want: "NOT myField:*",
		},
		{
			name: "Test negate with empty value",
			fields: fields{
				FieldName:   "myField",
				SimpleValue: &SimpleValue{},
			},
			want: "NOT myField:*",
		},
		{
			name: "Test negate",
			fields: fields{
				FieldName: "myField",
				SimpleValue: &SimpleValue{
					Value: "myVal",
				},
			},
			want: "NOT myField:myVal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FieldQuery{
				FieldName:   tt.fields.FieldName,
				RangeValue:  tt.fields.RangeValue,
				SimpleValue: tt.fields.SimpleValue,
			}
			if got := f.Negate(); got != tt.want {
				t.Errorf("FieldQuery.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}
