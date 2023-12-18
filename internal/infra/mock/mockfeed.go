package cache

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

type MockFeedRepository struct{}

func NewMockFeedRepository() *MockFeedRepository {
	return &MockFeedRepository{}
}

func (m *MockFeedRepository) GetAllFeed() ([]responses.Feed, error) {
	return []responses.Feed{
		{
			ID:          1,
			Title:       "Title",
			Description: "Description",
		},
	}, nil
}

func (m *MockFeedRepository) GetFeed() (responses.Feed, error) {
	return responses.Feed{
		ID:          1,
		Title:       "Title",
		Description: "Description",
	}, nil
}
