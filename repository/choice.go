package repository

type Choice struct {
	ID         int    `db:"id"`
	ChoiceText string `db:"choice_text"`
	QuestionID int    `db:"question_id"`
	Votes      int    `db:"votes"`
}

type ChoiceRepository interface {
	GetAllChoiceFromQuestion(Question) ([]Choice, error)
	GetById(int) (*Choice, error)
	AddVote(int) (*Choice, error)
}
