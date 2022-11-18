package tgbotmessage

import (
	"context"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMessagesHandler_New(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockActionSetterGetter(ctrl)

	handler := NewHandler(mockStore)

	require.NotNil(t, handler)
}

func TestMessagesHandler_GetAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := struct {
		ID     int64
		Return Action
	}{
		ID:     1,
		Return: "#",
	}

	mockStore := NewMockActionSetterGetter(ctrl)

	mockStore.
		EXPECT().
		GetAction(context.Background(), data.ID).
		Return(data.Return)

	handler := NewHandler(mockStore)

	action := handler.GetAction(data.ID)
	require.Equal(t, data.Return, action)
}

func TestMessagesHandler_SetAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := struct {
		ID     int64
		Action Action
	}{
		ID:     1,
		Action: "#",
	}

	mockStore := NewMockActionSetterGetter(ctrl)

	mockStore.
		EXPECT().
		SetAction(context.Background(), data.ID, data.Action).
		Return()

	handler := NewHandler(mockStore)
	handler.SetAction(data.ID, data.Action)
}

func TestMessagesHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGetAction := struct {
		ID     int64
		Return Action
	}{
		ID:     1,
		Return: "#",
	}

	testsTable := []struct {
		Name                   string
		Action                 Action
		Update                 tgbotapi.Update
		expectedFuncRunned     bool
		expectedNotFoundRunned bool
	}{
		{
			Name:   "OK_command",
			Action: "/start",
			Update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Text: "/start",
					Entities: []tgbotapi.MessageEntity{
						{
							Type:   "bot_command",
							Length: 6,
						},
					},
					From: &tgbotapi.User{
						ID: mockGetAction.ID,
					},
				},
			},
			expectedFuncRunned:     true,
			expectedNotFoundRunned: false,
		},
		{
			Name:   "OK_action",
			Action: "#",
			Update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Text: "hi!",
					From: &tgbotapi.User{
						ID: mockGetAction.ID,
					},
				},
			},
			expectedFuncRunned:     true,
			expectedNotFoundRunned: false,
		},
		{
			Name:   "OK_command_not_found",
			Action: "#",
			Update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Text: "/start",
					Entities: []tgbotapi.MessageEntity{
						{
							Type:   "bot_command",
							Length: 6,
						},
					},
					From: &tgbotapi.User{
						ID: mockGetAction.ID,
					},
				},
			},
			expectedFuncRunned:     false,
			expectedNotFoundRunned: true,
		},
		{
			Name:   "OK_action_not_found",
			Action: "#foo",
			Update: tgbotapi.Update{
				Message: &tgbotapi.Message{
					Text: "hi!",
					From: &tgbotapi.User{
						ID: mockGetAction.ID,
					},
				},
			},
			expectedFuncRunned:     false,
			expectedNotFoundRunned: true,
		},
	}

	mockStore := NewMockActionSetterGetter(ctrl)

	mockStore.
		EXPECT().
		GetAction(context.Background(), mockGetAction.ID).
		Return(mockGetAction.Return).
		AnyTimes()

	for _, item := range testsTable {
		handler := NewHandler(mockStore)

		t.Run(item.Name, func(t *testing.T) {
			var (
				funcRunned     bool
				notFoundRunned bool
			)

			handler.Add(item.Action, func(upd tgbotapi.Update) (next bool) {
				funcRunned = true
				return
			})

			handler.NotFound = func(upd tgbotapi.Update) (next bool) {
				notFoundRunned = true
				return
			}

			handler.Handle(item.Update)

			require.Equal(t, item.expectedFuncRunned, funcRunned)
			require.Equal(t, item.expectedNotFoundRunned, notFoundRunned)
		})
	}
}
