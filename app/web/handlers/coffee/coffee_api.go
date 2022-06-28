package coffee

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/foundation/web"
	"net/http"
)

type Handlers struct {
	CoffeeService *coffee.Service
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {

	var ncb coffee.NewCoffeeBean
	if err := web.Decode(r, &ncb); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.CoffeeService.CreateCoffeeBean(ctx, ncb)
	if err != nil {
		return fmt.Errorf("creating new coffee bean, ncb[%+v]: %w", ncb, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}
