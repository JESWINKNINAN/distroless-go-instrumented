package main

import (
	"distroless-go-instrumented/pkg/web"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/web", web.SampleHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:9091", nil))
}
