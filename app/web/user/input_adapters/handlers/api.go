package user

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/common"
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/baransonmez/coff.app/foundation/web"
	"net/http"
	"strings"
)

type Handlers struct {
	UserService *user.Service
}

func (h Handlers) Create(w http.ResponseWriter, r *http.Request) error {

	var nu user.NewUser
	if err := web.Decode(r, &nu); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	ctx := r.Context()

	prod, err := h.UserService.CreateNewUser(ctx, nu)
	if err != nil {
		return fmt.Errorf("creating new user, nu[%+v]: %w", nu, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) error {

	id := strings.TrimPrefix(r.URL.Path, "/user/")

	ctx := r.Context()

	userUUID, err := common.StringToID(id)
	if err != nil {
		return fmt.Errorf("string to uuid, id[%+v]: %w", id, err)
	}

	prod, err := h.UserService.GetUser(ctx, userUUID)
	if err != nil {
		return fmt.Errorf("getting recipe, userUUID[%+v]: %w", userUUID, err)
	}

	return web.Respond(w, prod, http.StatusCreated)
}
