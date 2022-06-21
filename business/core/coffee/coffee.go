package coffee

import (
	"context"
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee/db"
	"time"
)

type Core struct {
	store db.Store
}

func (c Core) Create(ctx context.Context, np NewCoffeeBean, now time.Time) (Bean, error) {
	dbPrd := db.Bean{
		ID:          "",
		Name:        np.Name,
		Roaster:     np.Roaster,
		Origin:      np.Origin,
		RoastDate:   np.RoastDate,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.store.Create(ctx, dbPrd); err != nil {
		return Bean{}, fmt.Errorf("create: %w", err)
	}

	return toBean(dbPrd), nil
}
