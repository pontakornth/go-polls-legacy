package repository

type Question struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
}

type QuestionRepository interface {
	GetAll() ([]Question, error)
	GetById(int) (*Question, error)
}
