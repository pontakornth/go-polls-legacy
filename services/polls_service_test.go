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

func TestPollsService(t *testing.T) {
	t.Run("Get one question", func(t *testing.T) {
		polls := services.NewPollsService(mockQuestion, mockChoice)
		question, err := polls.GetQuestionById(1)
		if err != nil {
			t.Error(err)
			return
		}
		if question.ID != 1 {
			t.Error("Incorrect question returned")
		}
	})
}

func TestMain(m *testing.M) {

	mockChoice = repository.NewChoiceRepositoryMock()
	mockQuestion = repository.NewQuestionRepositoryMock()
	os.Exit(m.Run())
}