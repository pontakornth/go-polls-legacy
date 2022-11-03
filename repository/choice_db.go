package repository

import "github.com/jmoiron/sqlx"

type choiceRepository struct {
	db *sqlx.DB
}

// GetAllChoiceFromQuestion implements ChoiceRepository
func (c choiceRepository) GetAllChoiceFromQuestion(q Question) ([]Choice, error) {
	choices := []Choice{}
	err := c.db.Select(&choices, "SELECT * FROM choices c WHERE c.question_id = $1", q.ID)
	if err != nil {
		return nil, err
	}
	return choices, err
}

// GetById implements ChoiceRepository
func (choiceRepository) GetById(int) (*Choice, error) {
	panic("unimplemented")
}

func NewChoiceRepository(db *sqlx.DB) ChoiceRepository {
	return choiceRepository{db: db}
}
