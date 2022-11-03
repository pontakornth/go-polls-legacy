package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pontakornth/go-polls/services"
)

type pollsHandler struct {
	pollService services.PollsService
}

func NewPollsHandler(poll services.PollsService) pollsHandler {
	return pollsHandler{pollService: poll}
}

func (p pollsHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := p.pollService.GetAllQuestions()
	// TODO: Error handler
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func (p pollsHandler) GetQuestionById(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
