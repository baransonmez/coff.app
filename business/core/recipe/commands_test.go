package recipe

import "testing"

func TestNewRecipe_Validate(t *testing.T) {
	type fields struct {
		UserID      string
		CoffeeID    string
		Description string
		Steps       []Step
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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
		}, wantErr: false},
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
		}, wantErr: true},
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
		}, wantErr: true},
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
		}, wantErr: true},
		{name: "description empty fail", fields: fields{
			UserID:      "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			CoffeeID:    "9141f16c-cdb4-47eb-93a5-681c93e297cf",
			Description: "desc desc",
			Steps:       []Step{},
		}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NewRecipe{
				UserID:      tt.fields.UserID,
				CoffeeID:    tt.fields.CoffeeID,
				Description: tt.fields.Description,
				Steps:       tt.fields.Steps,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
