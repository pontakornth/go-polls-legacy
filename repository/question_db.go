package repository

import "github.com/jmoiron/sqlx"

type questionRepositoryDb struct {
	db *sqlx.DB
}

// GetAll implements QuestionRepository
func (questionRepositoryDb) GetAll() ([]Question, error) {
	panic("unimplemented")
}

// GetById implements QuestionRepository
func (questionRepositoryDb) GetById(int) (*Question, error) {
	panic("unimplemented")
}

func NewQuestionRepository(db *sqlx.DB) QuestionRepository {
	return questionRepositoryDb{db: db}
}
