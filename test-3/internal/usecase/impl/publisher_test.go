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

var (
	repoPublisher    = new(mocks.RepoPublisher)
	usecasePublisher = impl.NewUsecasePublisherImpl(repoPublisher, logger.New(logger.Config{
		Level:                 "debug",
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    false,
		EncodeLogsAsJson:      false,
	}))
)

func TestPublisherCreate(t *testing.T) {
	type testTableStruct struct {
		name      string
		request   *entity.Publisher
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			request: &entity.Publisher{
				Name:   "[TCP-SCS] Republikin",
				Detail: "Republikin is one of the best publisher in Indonesia",
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
		repoPublisher.On("Create",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.request,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecasePublisher.Create(context.Background(), testCase.request)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestPublisherFetchById(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		response  *entity.Publisher
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			id:   1,
			response: &entity.Publisher{
				Name:   "[TCP-SCS] Republikin",
				Detail: "Republikin is one of the best publisher in Indonesia",
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
		repoPublisher.On("FetchById",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecasePublisher.FetchById(context.Background(), testCase.id)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestPublisherFetchByName(t *testing.T) {
	type testTableStruct struct {
		name          string
		publisherName string
		response      []*entity.Publisher
		errStatus     error
	}

	testTable := []testTableStruct{
		{
			name:          "Success",
			publisherName: "[TCP-SCS] Republikin",
			response: []*entity.Publisher{
				{
					Name: "[TCP-SCS] Republikin",
				},
			},
			errStatus: nil,
		},
		{
			name:          "Fail",
			publisherName: "",
			response:      nil,
			errStatus:     fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoPublisher.On("FetchByName",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.publisherName,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecasePublisher.FetchByName(context.Background(), testCase.publisherName)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestPublisherFetchAll(t *testing.T) {
	type testTableStruct struct {
		name      string
		response  []*entity.Publisher
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			response: []*entity.Publisher{
				{
					Name:   "[TCP-SCS] Republikin",
					Detail: "Republikin is one of the best publisher in Indonesia",
				},
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			response:  []*entity.Publisher(nil),
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoPublisher.On("FetchAll",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecasePublisher.FetchAll(context.Background())

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestPublisherUpdate(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		request   *entity.Publisher
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			id:   1,
			request: &entity.Publisher{
				Name:   "[TCP-SCS] Republikin",
				Detail: "Republikin is one of the best publisher in Indonesia",
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
		repoPublisher.On("Update",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
			testCase.request,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecasePublisher.Update(context.Background(), testCase.id, testCase.request)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestPublisherDelete(t *testing.T) {
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
		repoPublisher.On("Delete",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecasePublisher.Delete(context.Background(), testCase.id)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}
