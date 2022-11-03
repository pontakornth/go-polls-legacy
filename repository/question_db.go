package repository

import "github.com/jmoiron/sqlx"

type questionRepositoryDb struct {
	db *sqlx.DB
}

// GetAll implements QuestionRepository
func (q questionRepositoryDb) GetAll() ([]Question, error) {
	questions := []Question{}
	err := q.db.Select(&questions, "SELECT * FROM questions")
	if err != nil {
		return nil, err
	}
	return questions, err
}

// GetById implements QuestionRepository
func (q questionRepositoryDb) GetById(ID int) (*Question, error) {
	var question Question
	err := q.db.Get(&question, "SELECT * FROM questions q WHERE q.id = $1", ID)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func NewQuestionRepository(db *sqlx.DB) QuestionRepository {
	return questionRepositoryDb{db: db}
}
