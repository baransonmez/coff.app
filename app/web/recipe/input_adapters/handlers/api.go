package recipe

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"github.com/baransonmez/coff.app/foundation/web"
	"net/http"
)

type Handlers struct {
	RecipeService *recipe.Service
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {

	var nr recipe.NewRecipe
	if err := web.Decode(r, &nr); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.RecipeService.CreateNewRecipe(ctx, nr)
	if err != nil {
		return fmt.Errorf("creating new coffee bean, nr[%+v]: %w", nr, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) error {

	ctx := r.Context()

	recipeUUID, err := web.ReadIDParam(r)
	if err != nil {
		return err
	}

	prod, err := h.RecipeService.GetRecipe(ctx, recipeUUID)
	if err != nil {
		return fmt.Errorf("getting recipe, recipeUUID[%+v]: %w", recipeUUID, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}
