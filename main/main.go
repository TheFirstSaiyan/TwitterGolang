package main

import (
	"example/layered-architecture/handlers"
	"example/layered-architecture/repositories"
	"example/layered-architecture/services"
	"log"
	"net/http"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func setUpRoutes(handler *handlers.Handler) {
	r := mux.NewRouter().StrictSlash(true)

	//routes for the apis
	r.HandleFunc("/api/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/api/user", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/user", handler.AddUser).Methods("POST")
	r.HandleFunc("/api/tweet", handler.AddTweet).Methods("POST")
	r.HandleFunc("/api/user/tweets/{username}", handler.GetTweetsOfUser).Methods("GET")
	r.HandleFunc("/api/user/followees/{username}", handler.GetFolloweesOfUser).Methods("GET")
	r.HandleFunc("/api/follow", handler.AddFollowee).Methods("POST")
	r.HandleFunc("/api/tweet/{tweetid}", handler.DeleteTweet).Methods("DELETE")
	r.HandleFunc("/api/user/followees/{username}/{followeename}", handler.DeleteFollowee).Methods("DELETE")
	r.HandleFunc("/api/user/followees/{username}/{followeename}", handler.CheckFollowing).Methods("GET")

	//allowing CORS for the client
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})

	h := c.Handler(r)
	//start server
	err := http.ListenAndServe(":8000", h)
	if err != nil {
		log.Fatal("cant start server")
	}

}

func main() {
	const dsn = "root:FBenGax11@@tcp(127.0.0.1:3306)/demodb?parseTime=true"

	repository := repositories.NewMySqlRepository(dsn)
	service := services.NewUserService(repository)
	handler := handlers.NewHandler(service)

	setUpRoutes(handler)
}
