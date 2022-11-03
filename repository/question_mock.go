package repository

import "errors"

type questionRepositoryMock struct {
	questions []Question
}

// GetAll implements QuestionRepository
func (q questionRepositoryMock) GetAll() ([]Question, error) {
	return q.questions, nil
}

// GetById implements QuestionRepository
func (q questionRepositoryMock) GetById(ID int) (*Question, error) {
	for _, question := range q.questions {
		if question.ID == ID {
			return &question, nil
		}
	}
	return nil, errors.New("question not found")
}

func NewQuestionRepositoryMock() QuestionRepository {
	repo := questionRepositoryMock{
		questions: []Question{
			{ID: 1, Title: "What is your favorite color?"},
			{ID: 2, Title: "What is your favorite game"},
			{ID: 3, Title: "Who is the best programmer?"},
		},
	}
	return repo

}
