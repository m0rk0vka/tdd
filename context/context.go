package context

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Store interface {
	Fetch(context.Context) (string, error)
}

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.Fetch(r.Context())
		if err != nil {
			log.Printf("While fetching: %v", err)
			return
		}

		fmt.Fprint(w, data)
	}
}
