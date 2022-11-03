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
func (polls pollsService) GetQuestionById(ID int) (*QuestionDetailResponse, error) {
	question, err := polls.questionRepo.GetById(ID)
	if err != nil {
		return nil, err
	}
	choices, err := polls.choiceRepo.GetAllChoiceFromQuestion(*question)
	if err != nil {
		return nil, err
	}
	detailedQuestion := &QuestionDetailResponse{
		ID:      ID,
		Title:   question.Title,
		Choices: []ChoiceResponse{},
	}
	for _, c := range choices {
		detailedQuestion.Choices = append(detailedQuestion.Choices, ChoiceResponse{
			ID:         c.ID,
			ChoiceText: c.ChoiceText,
			Votes:      c.Votes,
		})
	}
	return detailedQuestion, nil
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
