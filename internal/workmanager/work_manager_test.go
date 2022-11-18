package workmanager

import (
	"context"
	"gold-panel/internal/service/v1"
	mock_service "gold-panel/internal/service/v1/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const token = "5688392235:AAGtbfDZlp5AbIrd4qE-XYb-vmpLN0w0BIk"

func TestWorkManager_Add(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockService := mock_service.NewMockIService(ctrl)
	ctx := context.Background()

	workman := NewWorkManager(mockService)

	me, err := workman.GetBotMe(token)
	require.Empty(t, err)

	workerAddDTO := service.WorkerAddDTO{
		ID:       me.ID,
		Token:    token,
		Username: me.UserName,
	}

	mockService.
		EXPECT().
		WorkerAdd(ctx, &workerAddDTO).
		Return(nil)

	workman.Add(&AddDTO{
		ID:       workerAddDTO.ID,
		Token:    workerAddDTO.Token,
		Username: workerAddDTO.Username,
	})
}

func TestWorkManager_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := mock_service.NewMockIService(ctrl)
	ctx := context.Background()

	table := struct {
		Name     string
		Token    string
		WorkerID int64
	}{
		Name:     "OK",
		Token:    token,
		WorkerID: 1,
	}

	mockService.
		EXPECT().
		WorkerGetToken(ctx, table.WorkerID).
		Return(table.Token, nil)

	workman := NewWorkManager(mockService)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	err := workman.Run(ctx, table.WorkerID)
	require.Empty(t, err)
}
