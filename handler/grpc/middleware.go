package handler

import (
	"log"
	"net/http"
)

func (g *GrpcHandler) PanicHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panicHandler error: %v", err)

				http.Error(
					w,
					`{"code": 13, "message": "Internal Server Error", "details": []}`,
					http.StatusInternalServerError,
				)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
