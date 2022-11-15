package repository

import "github.com/jmoiron/sqlx"

type choiceRepository struct {
	db *sqlx.DB
}

// AddVote implements ChoiceRepository
func (choiceRepository) AddVote(int) (*Choice, error) {
	panic("unimplemented")
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
func (c choiceRepository) GetById(ID int) (*Choice, error) {
	var choice Choice
	err := c.db.Get(&choice, "SELECT * FROM choices c WHERE c.id == $1", ID)
	if err != nil {
		return nil, err
	}
	return &choice, err
}

func NewChoiceRepository(db *sqlx.DB) ChoiceRepository {
	return choiceRepository{db: db}
}
