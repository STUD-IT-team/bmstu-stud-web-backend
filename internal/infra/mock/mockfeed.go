package cache

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

type MockFeedRepository struct{}

func NewMockFeedRepository() *MockFeedRepository {
	return &MockFeedRepository{}
}

func (m *MockFeedRepository) GetAllFeed() ([]responses.GetAllFeed, error) {
	return []responses.GetAllFeed{
		{
			ID:          1,
			Title:       "Title",
			Description: "Description",
		},
	}, nil
}

func (m *MockFeedRepository) GetFeed() (responses.GetFeed, error) {
	return responses.GetFeed{

		ID:          1,
		Title:       "Title",
		Description: "Description",
	}, nil
}
