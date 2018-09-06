package main

import (
	"net/http"
)

func SecureWithScope(scope string, f http.HandlerFunc) http.Handler {
	realHandler := BuildSecureHandler(f)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !hasScope(r, scope) {
			message := "Insufficient scope."
			responseJSON(message, w, http.StatusForbidden)
			return
		}

		realHandler.ServeHTTP(w, r)
	})
}

func BuildSecureHandler(f http.HandlerFunc) http.Handler {
	jwtMiddleware := BuildMiddleware()

	return jwtMiddleware.Handler(http.HandlerFunc(f))
}
