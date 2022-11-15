package repository

import "errors"

type choiceRepositoryMock struct {
	choices []Choice
}

// AddVote implements ChoiceRepository
func (c choiceRepositoryMock) AddVote(ID int) (*Choice, error) {
	for _, choice := range c.choices {
		if choice.ID == ID {
			choice.Votes++
			return &choice, nil
		}
	}
	return nil, errors.New("Choice not found")
}

// GetAllChoiceFromQuestion implements ChoiceRepository
func (c choiceRepositoryMock) GetAllChoiceFromQuestion(q Question) ([]Choice, error) {
	if q.ID == 1 {
		return c.choices[:3], nil
	} else if q.ID == 2 {
		return c.choices[3:6], nil
	} else if q.ID == 3 {
		return c.choices[6:], nil
	} else {
		return nil, errors.New("question not found")
	}
}

// GetById implements ChoiceRepository
func (c choiceRepositoryMock) GetById(ID int) (*Choice, error) {
	for _, choice := range c.choices {
		if choice.ID == ID {
			return &choice, nil
		}
	}
	return nil, errors.New("Choice not found")
}

func NewChoiceRepositoryMock() ChoiceRepository {
	repo := choiceRepositoryMock{
		choices: []Choice{
			// Question 1
			{ID: 1, ChoiceText: "Red", QuestionID: 1, Votes: 0},
			{ID: 2, ChoiceText: "Green", QuestionID: 1, Votes: 0},
			{ID: 3, ChoiceText: "Blue", QuestionID: 1, Votes: 0},
			// Question 2
			{ID: 4, ChoiceText: "Genshin Impact", QuestionID: 2, Votes: 0},
			{ID: 5, ChoiceText: "Honkai Impact", QuestionID: 2, Votes: 0},
			{ID: 6, ChoiceText: "Blue Archive", QuestionID: 2, Votes: 0},
			// Question 3
			{ID: 7, ChoiceText: "Amogus", QuestionID: 3, Votes: 0},
			{ID: 8, ChoiceText: "Impostor", QuestionID: 3, Votes: 0},
			{ID: 9, ChoiceText: "Impossible", QuestionID: 3, Votes: 0},
		},
	}
	return repo

}
