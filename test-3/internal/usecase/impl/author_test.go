package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/entity"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/impl"
	"github.com/fahmyabdul/gits-assignments/test-3/internal/usecase/mocks"
	"github.com/fahmyabdul/gits-assignments/test-3/pkg/logger"
)

const (
	_emptyRequest = "request must not empty"
)

var (
	repoAuthor    = new(mocks.RepoAuthor)
	usecaseAuthor = impl.NewUsecaseAuthorImpl(repoAuthor, logger.New(logger.Config{
		Level:                 "debug",
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    false,
		EncodeLogsAsJson:      false,
	}))
)

// Create(ctx context.Context, request *entity.Author) error
func TestAuthorCreate(t *testing.T) {
	type testTableStruct struct {
		name      string
		request   *entity.Author
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			request: &entity.Author{
				Name:   "[TCA-SCS] Fahmy Abdul",
				Detail: "This is the detail of Test Create Author",
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			request:   nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("Create",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.request,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseAuthor.Create(context.Background(), testCase.request)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// FetchById(ctx context.Context, id int) (*entity.Author, error)
func TestAuthorFetchById(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		response  *entity.Author
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			id:   1,
			response: &entity.Author{
				Name:   "[TCA-SCS] Fahmy Abdul",
				Detail: "This is the detail of Test Create Author",
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			id:        0,
			response:  nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("FetchById",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseAuthor.FetchById(context.Background(), testCase.id)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// FetchByName(ctx context.Context, name string) ([]*entity.Author, error)
func TestAuthorFetchByName(t *testing.T) {
	type testTableStruct struct {
		name       string
		authorName string
		response   []*entity.Author
		errStatus  error
	}

	testTable := []testTableStruct{
		{
			name:       "Success",
			authorName: "[TCA-SCS] Fahmy Abdul",
			response: []*entity.Author{
				{
					Name:   "[TCA-SCS] Fahmy Abdul",
					Detail: "This is the detail of Test Create Author",
				},
			},
			errStatus: nil,
		},
		{
			name:       "Fail",
			authorName: "",
			response:   nil,
			errStatus:  fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("FetchByName",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.authorName,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseAuthor.FetchByName(context.Background(), testCase.authorName)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// FetchAll(ctx context.Context) ([]*entity.Author, error)
func TestAuthorFetchAll(t *testing.T) {
	type testTableStruct struct {
		name      string
		response  []*entity.Author
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			response: []*entity.Author{
				{
					Name:   "[TCA-SCS] Fahmy Abdul",
					Detail: "This is the detail of Test Create Author",
				},
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			response:  []*entity.Author(nil),
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("FetchAll",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseAuthor.FetchAll(context.Background())

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// Update(ctx context.Context, id int, request *entity.Author) error
func TestAuthorUpdate(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		request   *entity.Author
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			id:   1,
			request: &entity.Author{
				Name:   "[TCA-SCS] Fahmy Abdul",
				Detail: "This is the detail of Test Create Author",
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			id:        0,
			request:   nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("Update",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
			testCase.request,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseAuthor.Update(context.Background(), testCase.id, testCase.request)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// Delete(ctx context.Context, id int) error
func TestAuthorDelete(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name:      "Success",
			id:        1,
			errStatus: nil,
		},
		{
			name:      "Fail",
			id:        0,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoAuthor.On("Delete",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseAuthor.Delete(context.Background(), testCase.id)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}
