package workerbot

import (
	"context"
	mock_service "gold-panel/internal/service/v1/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	token = "5688392235:AAGtbfDZlp5AbIrd4qE-XYb-vmpLN0w0BIk"
)

func TestNewWorkerBot(t *testing.T) {
	ctrl := gomock.NewController(t)

	table := []struct {
		Name     string
		Token    string
		WorkerID int64
		Handler  func(w *WorkerBot, err error)
	}{
		{
			Name:     "OK",
			Token:    token,
			WorkerID: 1,
			Handler: func(w *WorkerBot, err error) {
				require.Empty(t, err)
				require.NotEmpty(t, w)
			},
		},
		{
			Name:     "BAD",
			Token:    "5688392235:AAGtbfDZlp5AbIrd4qE-XYb-vmpLN0w0BIx",
			WorkerID: 2,
			Handler: func(w *WorkerBot, err error) {
				require.NotEmpty(t, err)
				require.Empty(t, w)
			},
		},
	}

	mockService := mock_service.NewMockIService(ctrl)

	ctx := context.Background()

	for _, item := range table {
		t.Run(item.Name, func(t *testing.T) {
			mockService.
				EXPECT().
				WorkerGetToken(ctx, item.WorkerID).
				Return(item.Token, nil)

			worker, err := NewWorkerBot(ctx, item.WorkerID, mockService)
			item.Handler(worker, err)
		})
	}
}

func TestNewWorkerBot_Listen(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockService := mock_service.NewMockIService(ctrl)
	ctx := context.Background()

	var (
		workerID int64 = 1
		err      error
	)

	mockService.
		EXPECT().
		WorkerGetToken(ctx, workerID).
		Return(token, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	worker, err := NewWorkerBot(ctx, workerID, mockService)
	require.Empty(t, err)

	err = worker.Listen()
	require.Empty(t, err)
}
