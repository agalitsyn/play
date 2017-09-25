package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Port         string
	ResponseCode int
}

func main() {
	var (
		port         = flag.String("port", "5050", "HTTP port")
		responseCode = flag.Int("response-code", http.StatusOK, "Response HTTP code")
	)
	flag.Parse()

	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}
	if os.Getenv("RESPONSE_CODE") != "" {
		envCode, err := strconv.Atoi(os.Getenv("RESPONSE_CODE"))
		if err != nil {
			log.Fatal(err)
		}
		*responseCode = envCode
	}

	conf := Config{Port: *port, ResponseCode: *responseCode}
	log.Printf("conf = %+v\n", conf)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(
			w,
			fmt.Sprintf("Requested URL: %s \nResponse code: %v \nMessage: %s", r.URL.Path, conf.ResponseCode, http.StatusText(conf.ResponseCode)),
			conf.ResponseCode,
		)

		log.Printf("%s %s %v | remote_ip: %s | http_referer: %s | http_user_agent: %s | http_proto: %s",
			r.Method, r.URL.Path, conf.ResponseCode, r.RemoteAddr, r.Referer(), r.UserAgent(), r.Proto)
	})

	httpAddr := net.JoinHostPort("", conf.Port)
	log.Printf("Listening at %s", httpAddr)
	http.ListenAndServe(httpAddr, nil)
}
