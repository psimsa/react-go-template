package web

import (
	"fmt"
	"main/models"
	"main/utils"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve(files http.FileSystem, addr string) {
	fileServer := http.FileServer(files)

	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/data", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", models.NewApiData("Hello world!")))
	}).Methods("GET", "OPTIONS")
	api.HandleFunc("/metrics", getMetrics).Methods("GET", "OPTIONS")
	api.HandleFunc("/data", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", "Successfully saved"))
	}).Methods("POST", "OPTIONS")

	router.PathPrefix("/").Handler(fileServer)
	cors := handlers.AllowedOrigins([]string{"*"})

	srv := handlers.CombinedLoggingHandler(os.Stdout, router)
	server := &http.Server{
		Addr:    addr,
		Handler: handlers.CORS(cors)(srv),
	}
	fmt.Println("Listening on:", addr)
	err := server.ListenAndServe()
	fmt.Println(err)
}

func getMetrics(rw http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(rw, models.NewApiResponse("OK", "some metrics would be here"))
}
