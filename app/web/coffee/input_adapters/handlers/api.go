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

func (h Handlers) GetCoffee(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	id, err := web.ReadIDParam(r)
	if err != nil {
		return err
	}
	prod, err := h.CoffeeService.GetCoffeeBean(ctx, id)
	if err != nil {
		return fmt.Errorf("getting coffee bean, id[%+v]: %w", id, err)
	}

	return web.Respond(w, prod, http.StatusOK)
}
