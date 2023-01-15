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
	repoBook    = new(mocks.RepoBook)
	usecaseBook = impl.NewUsecaseBookImpl(repoBook, logger.New(logger.Config{
		Level:                 "debug",
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    false,
		EncodeLogsAsJson:      false,
	}))
)

func TestBookCreate(t *testing.T) {
	type testTableStruct struct {
		name      string
		request   *entity.Book
		authorId  int
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			request: &entity.Book{
				Name:        "[TCB-SCS] Pulang & Pergi",
				Pages:       100,
				PublisherId: 1,
			},
			authorId:  1,
			errStatus: nil,
		},
		{
			name:      "Fail",
			request:   nil,
			authorId:  0,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoBook.On("Create",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.request,
			testCase.authorId,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseBook.Create(context.Background(), testCase.request, testCase.authorId)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestBookFetchById(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		response  *entity.Book
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			id:   1,
			response: &entity.Book{
				Name:        "[TCB-SCS] Pulang & Pergi",
				Pages:       100,
				PublisherId: 1,
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
		repoBook.On("FetchById",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseBook.FetchById(context.Background(), testCase.id)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestBookFetchByName(t *testing.T) {
	type testTableStruct struct {
		name      string
		bookName  string
		response  []*entity.Book
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name:     "Success",
			bookName: "[TCB-SCS] Pulang & Pergi",
			response: []*entity.Book{
				{
					Name:        "[TCB-SCS] Pulang & Pergi",
					Pages:       100,
					PublisherId: 1,
				},
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			bookName:  "",
			response:  nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoBook.On("FetchByName",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.bookName,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseBook.FetchByName(context.Background(), testCase.bookName)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestBookFetchAll(t *testing.T) {
	type testTableStruct struct {
		name      string
		response  []*entity.Book
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name: "Success",
			response: []*entity.Book{
				{
					Name:        "[TCB-SCS] Pulang & Pergi",
					Pages:       100,
					PublisherId: 1,
				},
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			response:  []*entity.Book(nil),
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoBook.On("FetchAll",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseBook.FetchAll(context.Background())

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestBookUpdate(t *testing.T) {
	type testTableStruct struct {
		name      string
		id        int
		authorId  int
		request   *entity.Book
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name:     "Success",
			id:       1,
			authorId: 1,
			request: &entity.Book{
				Name:        "[TCB-SCS] Pulang & Pergi",
				Pages:       100,
				PublisherId: 1,
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			id:        0,
			authorId:  0,
			request:   nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoBook.On("Update",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
			testCase.authorId,
			testCase.request,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseBook.Update(context.Background(), testCase.id, testCase.authorId, testCase.request)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

func TestBookDelete(t *testing.T) {
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
		repoBook.On("Delete",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.id,
		).Return(testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			err := usecaseBook.Delete(context.Background(), testCase.id)

			assert.Equal(t, testCase.errStatus, err)
		})
	}
}

// func (p *UsecaseBookImpl) FetchByAuthorId(ctx context.Context, author_id int) (*entity.BookByAuthor, error) {
func TestBookFetchByAuthorId(t *testing.T) {
	type testTableStruct struct {
		name      string
		author_id int
		response  *entity.BookByAuthor
		errStatus error
	}

	testTable := []testTableStruct{
		{
			name:      "Success",
			author_id: 1,
			response: &entity.BookByAuthor{
				AuthorId:   1,
				AuthorName: "[TCA-SCS] Fahmy Abdul",
				Books: []entity.Book{
					{
						Id:          1,
						Name:        "[TCB-SCS] Pulang & Pergi",
						Pages:       100,
						PublisherId: 1,
					},
				},
			},
			errStatus: nil,
		},
		{
			name:      "Fail",
			author_id: 0,
			response:  nil,
			errStatus: fmt.Errorf(_emptyRequest),
		},
	}

	for _, testCase := range testTable {
		repoBook.On("FetchByAuthorId",
			mock.MatchedBy(func(_ context.Context) bool { return true }),
			testCase.author_id,
		).Return(testCase.response, testCase.errStatus).Once()

		t.Run(testCase.name, func(t *testing.T) {
			response, err := usecaseBook.FetchByAuthorId(context.Background(), testCase.author_id)

			assert.Equal(t, testCase.response, response)
			assert.Equal(t, testCase.errStatus, err)
		})
	}
}
