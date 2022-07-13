package user

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser_Validate(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{name: "fill all fields success", fields: fields{Name: "Baran Sonmez"}, err: nil},
		{name: "name empty fail", fields: fields{Name: ""}, err: &common.CannotBeEmptyError{Field: "name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &NewUser{
				Name: tt.fields.Name,
			}
			err := u.Validate()
			assert.Equal(t, err, tt.err)
		})
	}
}
