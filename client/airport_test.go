package client

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/oleewere/cldnt/model"
)

type AirportClientMock struct {
}

type AirportClientMockWithError struct {
}

func (a *AirportClientMock) ProcessRequest(request *http.Request) ([]byte, error) {
	outputJsonStr := `{"rows": [{
        "fields": {
            "name": "Budapest Keleti",
            "lat": 47.0,
            "lon": 19.0
        }
    },
    {
        "fields": {
            "name": "Ferihegy",
            "lat": 47.0,
            "lon": 19.2
        }
	}]}`
	return []byte(outputJsonStr), nil
}

func (a *AirportClientMockWithError) ProcessRequest(request *http.Request) ([]byte, error) {
	return nil, errors.New("wrong request")
}

func TestListAirportsByDistance(t *testing.T) {
	mockClient := AirportClientMock{}
	mockErrorClient := AirportClientMockWithError{}

	expectedAirports := make([]model.Airport, 0)
	airport1 := model.Airport{
		AirportFields: model.AirportFields{
			Name:     "Budapest Keleti",
			Location: model.Location{Latitude: 47.0, Longitude: 19.0},
		},
	}
	airport2 := model.Airport{
		AirportFields: model.AirportFields{
			Name:     "Ferihegy",
			Location: model.Location{Latitude: 47.0, Longitude: 19.2},
		},
	}
	expectedAirports = append(expectedAirports, airport1)
	expectedAirports = append(expectedAirports, airport2)

	type args struct {
		client           Client
		location         model.Location
		rows             int
		airportSearchUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Airport
		wantErr bool
	}{
		{
			name: "Test airport response conversion",
			args: args{
				client: &mockClient,
			},
			want:    expectedAirports,
			wantErr: false,
		},
		{
			name: "Test with http client error response",
			args: args{
				client: &mockErrorClient,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListAirportsByDistance(tt.args.client, tt.args.location, tt.args.rows, tt.args.airportSearchUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAirportsByDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAirportsByDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
