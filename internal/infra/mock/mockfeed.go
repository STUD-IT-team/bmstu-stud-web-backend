package mock_storage

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"

type MockFeedRepository struct{}

func NewMockFeedRepository() *MockFeedRepository {
	return &MockFeedRepository{}
}

func (m *MockFeedRepository) GetAllFeed() (responses.GetAllFeed, error) {
	feed := []responses.Feed{
		{
			ID:          1,
			Title:       "Title",
			Description: "Description",
		},
	}
	return responses.GetAllFeed{
		Feed: feed,
	}, nil
}

func (m *MockFeedRepository) GetFeed() (responses.GetFeed, error) {
	return responses.GetFeed{
		ID:          1,
		Title:       "Title",
		Description: "Description",
	}, nil
}
