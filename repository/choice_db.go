package repository

import "github.com/jmoiron/sqlx"

type choiceRepository struct {
	db *sqlx.DB
}

// GetAllChoiceFromQuestion implements ChoiceRepository
func (choiceRepository) GetAllChoiceFromQuestion(Question) ([]Choice, error) {
	panic("unimplemented")
}

// GetById implements ChoiceRepository
func (choiceRepository) GetById(int) (*Choice, error) {
	panic("unimplemented")
}

func NewChoiceRepository(db *sqlx.DB) ChoiceRepository {
	return choiceRepository{db: db}
}
