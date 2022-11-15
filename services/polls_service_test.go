package services_test

import (
	"os"
	"testing"

	"github.com/pontakornth/go-polls/repository"
	"github.com/pontakornth/go-polls/services"
)

var (
	mockChoice   repository.ChoiceRepository
	mockQuestion repository.QuestionRepository
)

func TestGetOneQuestion(t *testing.T) {
	polls := services.NewPollsService(mockQuestion, mockChoice)
	question, err := polls.GetQuestionById(1)
	if err != nil {
		t.Error(err)
		return
	}
	if question.ID != 1 {
		t.Error("Incorrect question returned")
	}
}

func TestGetAllQuestions(t *testing.T) {

	polls := services.NewPollsService(mockQuestion, mockChoice)
	questions, err := polls.GetAllQuestions()
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(questions); l != 3 {
		t.Errorf("Wrong number of questions, expect 3 got %v", l)
	}
}

func TestVoteChoiceById(t *testing.T) {
	polls := services.NewPollsService(mockQuestion, mockChoice)
	choice, err := polls.VoteChoiceById(1)
	if err != nil {
		t.Errorf("Error %v", err)
	}
	if choice.Votes != 1 {
		t.Errorf("Wrong number of votes, expect 1 got %v", choice.Votes)
	}
}

func TestMain(m *testing.M) {

	mockChoice = repository.NewChoiceRepositoryMock()
	mockQuestion = repository.NewQuestionRepositoryMock()
	os.Exit(m.Run())
}
