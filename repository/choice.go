package repository

type Choice struct {
	ID         int    `db:"id"`
	ChoiceText string `db:"choice_text"`
	Votes      int    `db:"votes"`
}

type ChoiceRepository interface {
	GetAllChoiceFromQuestion(Question) ([]Choice, error)
	GetById(int) (*Choice, error)
}
