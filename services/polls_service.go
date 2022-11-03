package services

import "github.com/pontakornth/go-polls/repository"

type pollsService struct {
	questionRepo repository.QuestionRepository
	choiceRepo   repository.ChoiceRepository
}

// GetAllQuestions implements PollsService
func (polls pollsService) GetAllQuestions() ([]QuestionListResponse, error) {
	questions, err := polls.questionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	resQuestions := []QuestionListResponse{}
	for _, v := range questions {
		resQuestions = append(resQuestions, QuestionListResponse{
			ID:    v.ID,
			Title: v.Title,
		})
	}
	return resQuestions, nil
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
