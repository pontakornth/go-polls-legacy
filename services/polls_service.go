package services

import "github.com/pontakornth/go-polls/repository"

type pollsService struct {
	questionRepo repository.QuestionRepository
	choiceRepo   repository.ChoiceRepository
}

// GetAllQuestions implements PollsService
func (pollsService) GetAllQuestions() ([]QuestionListResponse, error) {
	panic("unimplemented")
}

// GetQuestionById implements PollsService
func (pollsService) GetQuestionById(int) (*QuestionDetailResponse, error) {
	panic("unimplemented")
}

// VoteChoice implements PollsService
func (pollsService) VoteChoice(int) (*ChoiceResponse, error) {
	panic("unimplemented")
}

func NewPollsService(questionRepo repository.QuestionRepository, choiceRepo repository.ChoiceRepository) PollsService {
	return pollsService{
		questionRepo: questionRepo,
		choiceRepo:   choiceRepo,
	}

}
