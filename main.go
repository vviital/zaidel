package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vviital/zaidel/handler"
)

func main() {
	r := mux.NewRouter()

	// Peaks endpoints
	r.HandleFunc("/peaks/settings", handler.V1GetDefaultSettingsForPeaks).Methods(http.MethodGet)
	r.HandleFunc("/peaks/{id}", handler.V1GetPeaksByID).Methods(http.MethodGet)
	r.HandleFunc("/peaks/{id}", handler.V1UpdatePeaksByID).Methods(http.MethodPut, http.MethodPatch)
	r.HandleFunc("/peaks", handler.V1CalculatePeaks).Methods(http.MethodPost)

	// Spectrum lines endpoints
	r.HandleFunc("/spectrumlines/settings", handler.V1GetDefaultSettingsForSpectrumLines).Methods(http.MethodGet)
	r.HandleFunc("/spectrumlines/{id}", handler.V1UpdateSpectrumLinesByID).Methods(http.MethodGet)
	r.HandleFunc("/spectrumlines/{id}", handler.V1GetSpectrumLinesByID).Methods(http.MethodPut, http.MethodPatch)
	r.HandleFunc("/spectrumlines", handler.V1CalculateSpectrumLines).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server is up and running on %s", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
