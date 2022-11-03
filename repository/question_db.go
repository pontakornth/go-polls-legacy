package repository

import "github.com/jmoiron/sqlx"

type questionRepositoryDb struct {
	db *sqlx.DB
}

// GetAll implements QuestionRepository
func (repo questionRepositoryDb) GetAll() ([]Question, error) {
	questions := []Question{}
	err := repo.db.Select(&questions, "SELECT * FROM questions")
	if err != nil {
		return nil, err
	}
	return questions, err
}

// GetById implements QuestionRepository
func (questionRepositoryDb) GetById(int) (*Question, error) {
	panic("unimplemented")
}

func NewQuestionRepository(db *sqlx.DB) QuestionRepository {
	return questionRepositoryDb{db: db}
}
