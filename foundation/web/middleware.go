package web

import (
	"log"
	"net/http"
)

type Res func(w http.ResponseWriter, r *http.Request) error

func Handle(h Res) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Printf("response failed: %v", err)
			if err != nil {
				return
			}
		}
	}
}
