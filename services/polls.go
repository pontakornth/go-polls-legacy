package services

type QuestionListResponse struct {
	ID    int    `json:"question_id"`
	Title string `json:"title"`
}

type ChoiceResponse struct {
	ID         int    `json:"choice_id"`
	ChoiceText string `json:"choice_text"`
	Votes      int    `json:"votes"`
}

type QuestionDetailResponse struct {
	ID      int              `json:"question_id"`
	Title   string           `json:"title"`
	Choices []ChoiceResponse `json:"choices"`
}

type PollsService interface {
	GetAllQuestions() ([]QuestionListResponse, error)
	GetQuestionById(int) (*QuestionDetailResponse, error)
	VoteChoiceById(int) (*ChoiceResponse, error)
}
