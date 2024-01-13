package app

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	mock_storage "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infra/mock"
)

func TestFeedService_GetAllFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock_storage.NewMockStorage(ctrl)
	logger := logrus.New()
	feedService := NewFeedService(logger, mockStorage)

	ctx := context.Background()

	feeds := []responses.Feed{
		{
			ID:          1,
			Title:       "testAll",
			Description: "testAbout",
		},
	}
	expectedResponse := &responses.GetAllFeed{
		Feed: feeds,
	}

	// Mock the storage method call
	mockStorage.EXPECT().GetAllFeed(ctx).Return([]domain.Feed{
		{
			ID:          1,
			Title:       "testAll",
			Description: "testAbout",
		},
	}, nil)

	// Call the service method
	actualResponse, err := feedService.GetAllFeed(ctx)

	// Compare the expected and actual responses
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedResponse, actualResponse))
}

func TestFeedService_GetAllFeed_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock_storage.NewMockStorage(ctrl)
	logger := logrus.New()
	feedService := NewFeedService(logger, mockStorage)

	ctx := context.Background()
	expectedError := errors.New("storage error")

	// Mock the storage method call to return an error
	mockStorage.EXPECT().GetAllFeed(ctx).Return(nil, expectedError)

	// Call the service method
	actualResponse, err := feedService.GetAllFeed(ctx)

	// Compare the expected error with actual error
	assert.Equal(t, expectedError, err)
	assert.Nil(t, actualResponse)
}

func TestFeedService_GetFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock_storage.NewMockStorage(ctrl)
	logger := logrus.New()
	feedService := NewFeedService(logger, mockStorage)

	ctx := context.Background()
	expectedResponse := &responses.GetFeed{
		ID:             1,
		Title:          "testAll",
		Description:    "testAbout",
		RegistationURL: "testURL",
	}
	feedID := 1

	// Mock the storage method call
	mockStorage.EXPECT().GetFeed(ctx, feedID).Return(domain.Feed{
		ID:              1,
		Title:           "testAll",
		Description:     "testAbout",
		RegistrationURL: "testURL",
	}, nil)

	// Call the service method
	actualResponse, err := feedService.GetFeed(ctx, feedID)

	// Compare the expected and actual responses
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedResponse, actualResponse))
}

func TestFeedService_GetFeed_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock_storage.NewMockStorage(ctrl)
	logger := logrus.New()
	feedService := NewFeedService(logger, mockStorage)

	ctx := context.Background()
	expectedError := errors.New("storage error")
	feedID := 1

	// Mock the storage method call to return an error
	mockStorage.EXPECT().GetFeed(ctx, feedID).Return(domain.Feed{}, expectedError)

	// Call the service method
	actualResponse, err := feedService.GetFeed(ctx, feedID)

	// Compare the expected error with actual error
	assert.Equal(t, expectedError, err)
	assert.Nil(t, actualResponse)
}
