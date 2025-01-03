package main

import (
	// "fmt"
	"log"
	"net/http"
)

/* type apiConfig struct {
	fileserverHits int
}
*/
func main() {
	const filepathRoot = "."
	const port = "8080"

	/* apiCfg := apiConfig{
		fileserverHits: 0,
	}
	*/
	

	mux := http.NewServeMux()
	mux.Handle("/app/", /*apiCfg.middlewareMetricsInc( */ http.StripPrefix("/app",http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("GET /healthz", handlerReadiness)
	/*
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.handlerReset)
	*/

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
/*
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits)))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}
*/