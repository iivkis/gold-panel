package actstore

import (
	"context"
	"sync"
	"time"
)

type (
	Action = string
)

type Options struct {
	ActionDefault  Action
	ActionLifetime time.Duration
}

type ActionStore struct {
	mx      sync.Mutex
	store   map[int64]*Item
	options *Options
}

func NewStore(ctx context.Context, opts Options) *ActionStore {
	if opts.ActionLifetime.String() == "0s" {
		opts.ActionLifetime = time.Hour
	}

	storage := &ActionStore{
		store:   make(map[int64]*Item),
		options: &opts,
	}

	go storage.collector(ctx)
	return storage
}

func (s *ActionStore) SetAction(ctx context.Context, id int64, action Action) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.store[id] = &Item{
		Action:    action,
		ExpiresIn: time.Now().Add(s.options.ActionLifetime),
	}
}

func (s *ActionStore) GetAction(ctx context.Context, id int64) Action {
	s.mx.Lock()
	defer s.mx.Unlock()

	if item, ok := s.store[id]; ok {
		return item.Action
	}

	return s.options.ActionDefault
}

func (s *ActionStore) collector(ctx context.Context) {
cleaning:
	for {
		select {

		case <-ctx.Done():
			break cleaning

		case <-time.After(s.options.ActionLifetime):
			now := time.Now()

			for id := range s.store {
				if s.store[id].ExpiresIn.Before(now) {
					s.mx.Lock()
					delete(s.store, id)
					s.mx.Unlock()
				}
			}
		}
	}
}
