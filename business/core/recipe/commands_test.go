package recipe

import (
	"github.com/baransonmez/coff.app/business/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRecipe_Validate(t *testing.T) {
	type fields struct {
		UserID      string
		CoffeeID    string
		Description string
		Steps       []Step
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{name: "fill all fields success", fields: fields{
			UserID:      "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			CoffeeID:    "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			Description: "description description description",
			Steps: []Step{{
				Description:       "step 1",
				DurationInSeconds: 12,
			}, {
				Description:       "step 2",
				DurationInSeconds: 14,
			}},
		}, err: nil},
		{name: "user_id empty fail", fields: fields{
			UserID:      "",
			CoffeeID:    "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			Description: "description description description",
			Steps: []Step{{
				Description:       "step 1",
				DurationInSeconds: 12,
			}, {
				Description:       "step 2",
				DurationInSeconds: 14,
			}},
		}, err: &common.CannotBeEmptyError{Field: "user_id"}},
		{name: "coffee_id empty fail", fields: fields{
			UserID:      "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			CoffeeID:    "",
			Description: "description description description",
			Steps: []Step{{
				Description:       "step 1",
				DurationInSeconds: 12,
			}, {
				Description:       "step 2",
				DurationInSeconds: 14,
			}},
		}, err: &common.CannotBeEmptyError{Field: "coffee_id"}},
		{name: "description empty fail", fields: fields{
			UserID:      "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			CoffeeID:    "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			Description: "",
			Steps: []Step{{
				Description:       "step 1",
				DurationInSeconds: 12,
			}, {
				Description:       "step 2",
				DurationInSeconds: 14,
			}},
		}, err: &common.CannotBeEmptyError{Field: "description"}},
		{name: "steps size smaller than 1 fail", fields: fields{
			UserID:      "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			CoffeeID:    "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			Description: "desc desc",
			Steps:       []Step{},
		}, err: &common.CannotBeSmallerError{Field: "steps length", Limit: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NewRecipe{
				UserID:      tt.fields.UserID,
				CoffeeID:    tt.fields.CoffeeID,
				Description: tt.fields.Description,
				Steps:       tt.fields.Steps,
			}
			err := r.Validate()
			assert.Equal(t, err, tt.err)
		})
	}
}
