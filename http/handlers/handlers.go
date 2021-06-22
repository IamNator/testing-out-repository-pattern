package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gradelyng/testRepo/repo/internal/models"
	"github.com/gradelyng/testRepo/repo/internal/repo"
)

//Handler ...
type Handler struct {
	repo         repo.Repo
	user1, user2 []models.User
}	

//New ...
func New(repo repo.Repo, user1, user2 []models.User) Handler {
	return Handler{repo: repo, user1: user1, user2: user2}
}

//Handle uses orm for database operations
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllHandler1(w, r)
	case http.MethodPost:
		h.postHandler1(w, r)
	case http.MethodPut:
		h.updateHandler1(w, r)
	case http.MethodDelete:
		h.deleteHandler1(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("method not allowed"))
	}
}

var userids []uint

func (h *Handler) getAllHandler1(w http.ResponseWriter, r *http.Request) {

	users, er := h.repo.GetAllUsers()
	if er != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	if er := json.NewEncoder(w).Encode(users); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}
}

func (h *Handler) postHandler1(w http.ResponseWriter, r *http.Request) {

	users := h.user1

	usersCreated, er := h.repo.CreateUsers(users, &userids)
	if er != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	if er := json.NewEncoder(w).Encode(usersCreated); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}
}

func (h *Handler) updateHandler1(w http.ResponseWriter, r *http.Request) {

	users := h.user2
	_, er := h.repo.UpdateUsers(users, &userids)
	if er != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	if er := json.NewEncoder(w).Encode(userids); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

}

func (h *Handler) deleteHandler1(w http.ResponseWriter, r *http.Request) {

	er := h.repo.DeleteUsers(&userids)
	if er != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

	w.Header().Set("content-type", "application/json")
	if er := json.NewEncoder(w).Encode(map[string]interface{}{"message": "successfully deleted"}); er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(er.Error()))
		log.Println(er.Error())
		return
	}

}
