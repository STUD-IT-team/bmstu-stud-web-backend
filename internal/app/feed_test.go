package app

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/requests"
	"github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain/responses"
	mock "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/infrastructure/mock"
)

type FeedServiceTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	t           *testing.T
	mockStorage *mock.MockfeedServiceStorage
	feedService *FeedService
}

func NewFeedServiceTestSuite(t *testing.T) *FeedServiceTestSuite {
	return &FeedServiceTestSuite{t: t}
}

func (suite *FeedServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.t)
	suite.mockStorage = mock.NewMockfeedServiceStorage(suite.ctrl)
	// logger := logrus.New()
	// suite.feedService = NewFeedService(logger, suite.mockStorage)
}

func (suite *FeedServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *FeedServiceTestSuite) TestGetAllFeed() {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		filter           *requests.GetFeedByFilter
		request          []domain.Feed
		expectedResponse *responses.GetAllFeed
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			filter: &requests.GetFeedByFilter{
				Offset: mo.None[int](),
				Limit:  mo.None[int](),
			},
			expectedResponse: &responses.GetAllFeed{
				Feed: []responses.Feed{
					{
						ID:          1,
						Title:       "testAll",
						Description: "testAbout",
					},
				},
			},
			request: []domain.Feed{
				{
					ID:          1,
					Title:       "testAll",
					Description: "testAbout",
				}},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			filter:           &requests.GetFeedByFilter{},
			expectedResponse: nil,
			request:          []domain.Feed{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetAllFeed(ctx).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.feedService.GetFeedByFilter(ctx, *test.filter)

		// Compare the expected and actual responses
		assert.Equal(suite.T(), test.expectedError, actualError)
		assert.True(suite.T(), reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *FeedServiceTestSuite) TestGetFeed() {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest         string
		request          domain.Feed
		expectedResponse *responses.GetFeed
		expectedError    error
	}{
		1: {
			nameTest: "Test Ok",
			expectedResponse: &responses.GetFeed{
				ID:          1,
				Title:       "testAll",
				Description: "testAbout",
			},
			request: domain.Feed{
				ID:    1,
				Title: "testAll",

				Description: "testAbout",
			},
			expectedError: nil,
		},
		2: {
			nameTest:         "Test Error",
			expectedResponse: nil,
			request:          domain.Feed{},
			expectedError:    errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().GetFeed(ctx, 1).Return(test.request, test.expectedError)

		// Call the service method
		actualResponse, actualError := suite.feedService.GetFeed(ctx, 1)

		// Compare the expected and actual responses
		assert.Equal(suite.T(), test.expectedError, actualError)
		assert.True(suite.T(), reflect.DeepEqual(test.expectedResponse, actualResponse))
	}
}

func (suite *FeedServiceTestSuite) TestDeleteFeed() {
	ctx := context.Background()
	testCase := map[int]struct {
		nameTest      string
		request       int
		expectedError error
	}{
		1: {
			nameTest:      "Test Ok",
			request:       1,
			expectedError: nil,
		},
		2: {
			nameTest:      "Test Error",
			request:       0,
			expectedError: errors.New("storage error")},
	}

	for _, test := range testCase {
		// Mock the storage method call
		suite.mockStorage.EXPECT().DeleteFeed(ctx, test.request).Return(test.expectedError)

		// Call the service method
		actualError := suite.feedService.DeleteFeed(ctx, test.request)

		// Compare the expected and actual responses
		assert.Equal(suite.T(), test.expectedError, actualError)
	}
}

func TestFeedService_GetAllFeed(t *testing.T) {
	suite.Run(t, NewFeedServiceTestSuite(t))
}
