package web

import (
	"log"
	"net/http"
)

type APIHandler func(w http.ResponseWriter, r *http.Request) error

func Handle(h APIHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Printf("response failed: %v", err)
			_, err := w.Write([]byte(err.Error()))

			if err != nil {
				log.Printf("response failed: %v", err)
			}
		}
	}
}
