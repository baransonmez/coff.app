package coffee

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/stretchr/testify/assert"
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
		name   string
		fields fields
		err    error
	}{
		{name: "fill all fields success", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Now(),
		}, err: nil},
		{name: "roast date empty success", fields: fields{
			Name:    "montag coffee",
			Roaster: "montag",
			Origin:  "ethiopia",
			Price:   23,
		}, err: nil},
		{name: "name empty fail", fields: fields{
			Name:      "",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Time{},
		}, err: &common.CannotBeEmptyError{Field: "Name"}},
		{name: "roaster empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "",
			Origin:    "ethiopia",
			Price:     23,
			RoastDate: time.Time{},
		}, err: &common.CannotBeEmptyError{Field: "Roaster"}},
		{name: "origin empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "",
			Price:     23,
			RoastDate: time.Time{},
		}, err: &common.CannotBeEmptyError{Field: "Origin"}},
		{name: "price empty fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			RoastDate: time.Time{},
		}, err: &common.CannotBeSmallerError{Field: "Price", Limit: 1}},
		{name: "price smaller than 1 fail", fields: fields{
			Name:      "montag coffee",
			Roaster:   "montag",
			Origin:    "ethiopia",
			Price:     -23,
			RoastDate: time.Time{},
		}, err: &common.CannotBeSmallerError{Field: "Price", Limit: 1}},
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
			err := c.Validate()
			assert.Equal(t, err, tt.err)
		})
	}
}
