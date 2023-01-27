package handlers

import (
	"encoding/json"
	"example/layered-architecture/models"
	"example/layered-architecture/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service services.ServiceInterface
}

func NewHandler(service services.ServiceInterface) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//create container for the incoming user
	var user models.User

	//take the data from the request body and put it in the empty container
	json.NewDecoder(r.Body).Decode(&user)

	//username and password length validation
	if len(user.Name) < 3 || len(user.Password) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//call the UserService
	err := h.service.AddUser(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//take the data and show it in browser
	json.NewEncoder(w).Encode(&user)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	err := h.service.SignIn(&user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.service.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func (h *Handler) AddTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tweet models.Tweet
	//get the tweet from request
	json.NewDecoder(r.Body).Decode(&tweet)

	err := h.service.AddTweet(&tweet)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&tweet)

}

func (h *Handler) GetTweetsOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tweets, err := h.service.GetTweetsOfUser(params["username"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(tweets)

}

func (h *Handler) GetFolloweesOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	followees, err := h.service.GetFolloweesOfUser(params["username"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(followees)

}

func (h *Handler) AddFollowee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var follow models.Follows
	json.NewDecoder(r.Body).Decode(&follow)

	err := h.service.AddFollowee(&follow)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&follow)
}

func (h *Handler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	val, err := strconv.Atoi(params["tweetid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTweet(val)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("deleted tweet")

}

func (h *Handler) DeleteFollowee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	err := h.service.DeleteFollowee(params["username"], params["followeename"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("deleted followee")

}

func (h *Handler) CheckFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	err := h.service.CheckFollowing(params["username"], params["followeename"])

	if err != nil {
		w.WriteHeader(http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)

}
