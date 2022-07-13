package coffee

import (
	"testing"
	"time"
)

func TestNewCoffeeBean_Validate(t *testing.T) {
	type fields struct {
		Name      string
		Roaster   string
		Origin    string
		Price     int
		RoastDate time.Time
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "fill all fields success", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Now(),
		}, wantErr: false},
		{name: "roast date empty success", fields: fields{
			Name:    "montag coffee",
			Roaster: "montag",
			Origin:  "ethiopia",
			Price:   23,
		}, wantErr: false},
		{name: "name empty fail", fields: fields{
			Name:      "",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Time{},
		}, wantErr: true},
		{name: "roaster empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Time{},
		}, wantErr: true},
		{name: "origin empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "",
			Price:     23,
			RoastDate: time.Time{},
		}, wantErr: true},
		{name: "price empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			RoastDate: time.Time{},
		}, wantErr: true},
		{name: "price smaller than 1 fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     -23,
			RoastDate: time.Time{},
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NewCoffeeBean{
				Name:      tt.fields.Name,
				Roaster:   tt.fields.Roaster,
				Origin:    tt.fields.Origin,
				Price:     tt.fields.Price,
				RoastDate: tt.fields.RoastDate,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
