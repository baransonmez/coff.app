package web

import (
	"log"
	"net/http"
)

type DomainFunc func(w http.ResponseWriter, r *http.Request) error

func Handle(d DomainFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := d(w, r)
		if err != nil {
			log.Printf("response failed: %v", err)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				log.Printf("response failed: %v", err)
			}
		}
	}
}
