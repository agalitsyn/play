package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:5000", "")
	secret := flag.String("secret", "somesecret", "")
	flag.Parse()

	jwt := NewJWT(*secret, false, time.Hour, time.Hour)
	authProviders := makeAuthProviders(jwt)
	auth := Authenticator{
		JWTService: jwt,
		Providers:  authProviders,
	}
	r := chi.NewRouter()

	// auth routes for all providers
	r.Route("/auth", func(r chi.Router) {
		for _, provider := range auth.Providers {
			r.Mount("/"+provider.Name, provider.Routes()) // mount auth providers as /auth/{name}
		}
		if len(auth.Providers) > 0 {
			// shortcut, can be any of providers, all logouts do the same - removes cookie
			r.Get("/logout", auth.Providers[0].LogoutHandler)
		}
	})

	r.Route("/api", func(rapi chi.Router) {
		// rapi.Use(auth.Auth(true))
		rapi.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// ctx := r.Context.Get()
			username := "not real"
			w.Write([]byte(fmt.Sprintf("hello %s", username)))
		})
	})

	devAuth := &DevAuthServer{Provider: authProviders[len(authProviders)-1]}
	go devAuth.Run() // dev oauth2 server on :8084

	log.Printf("[INFO] api listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}

func makeAuthProviders(jwtService *JWT) []Provider {
	makeParams := func(cid, secret string) Params {
		return Params{
			JwtService: jwtService,
			RemarkURL:  "http://127.0.0.1:5000",
			Cid:        cid,
			Csecret:    secret,
		}
	}

	var providers []Provider
	providers = append(providers, NewDev(makeParams("", "")))

	if len(providers) == 0 {
		log.Printf("[WARN] no auth providers defined")
	}

	return providers
}
