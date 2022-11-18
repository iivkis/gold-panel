package repo

import (
	"context"
	"gold-panel/internal/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo_ApplicationInsert(t *testing.T) {
	repo := newRepo()

	table := []struct {
		name    string
		dto     *ApplicationInsertDTO
		handler func(insertID int64, err error)
	}{
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-69300483",
				Tag:   "SquirrelHat",
			},
			handler: func(insertID int64, err error) {
				require.NotEmpty(t, insertID, insertID)
				require.Empty(t, err)
			},
		},

		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "vk-954456789",
				Tag:   "AlphaRomeo",
			},
			handler: func(insertID int64, err error) {
				require.NotEmpty(t, insertID, insertID)
				require.Empty(t, err)
			},
		},

		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "vk-954456789",
				Tag:   "Degga",
			},
			handler: func(insertID int64, err error) {
				require.Empty(t, insertID)
				require.NotEmpty(t, err) // duplicate KeyID with previous query
			},
		},
	}

	for _, item := range table {
		t.Run(item.name, func(t *testing.T) {
			insertID, err := repo.ApplicationInsert(
				context.Background(),
				item.dto,
			)
			item.handler(insertID, err)
		})
	}
}

func TestRepo_ApplicationAccept(t *testing.T) {
	repo := newRepo()

	table := []struct {
		name    string
		dto     *ApplicationInsertDTO
		handler func(error)
	}{
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-4567893420",
				Tag:   "SquirrelHat",
			},
		},
	}

	for _, item := range table {
		_, err := repo.ApplicationInsert(
			context.Background(),
			item.dto,
		)
		require.Empty(t, err)

		{
			affected, err := repo.ApplicationAccept(
				context.Background(),
				&ApplicationAcceptDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.NotEqual(t, 0, affected)
		}

		{
			application, err := repo.ApplicationGet(
				context.Background(),
				&ApplicationAcceptDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.Equal(t, true, application.Invited)
		}
	}
}

func TestRepo_ApplicationDiscard(t *testing.T) {
	repo := newRepo()

	table := []struct {
		name    string
		dto     *ApplicationInsertDTO
		handler func(error)
	}{
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-3545352",
				Tag:   "SquirrelHat",
			},
		},
	}

	for _, item := range table {
		_, err := repo.ApplicationInsert(
			context.Background(),
			item.dto,
		)
		require.Empty(t, err)

		{
			affected, err := repo.ApplicationAccept(
				context.Background(),
				&ApplicationAcceptDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.NotEqual(t, 0, affected)
		}

		{
			affected, err := repo.ApplicationDiscard(
				context.Background(),
				&ApplicationDiscardDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.NotEqual(t, 0, affected)
		}

		{
			application, err := repo.ApplicationGet(
				context.Background(),
				&ApplicationAcceptDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.Equal(t, false, application.Invited)
		}
	}
}

func TestRepo_ApplicationSelect(t *testing.T) {
	repo := newRepo()

	//create new applications
	table := []struct {
		name    string
		dto     *ApplicationInsertDTO
		handler func(error)
	}{
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-456789",
				Tag:   "SquirrelHat",
			},
		},
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-384757",
				Tag:   "SquirrelHat",
			},
		},
		{
			name: "OK",
			dto: &ApplicationInsertDTO{
				KeyID: "tg-9282748",
				Tag:   "SquirrelHat",
			},
		},
	}

	for _, item := range table {
		_, err := repo.ApplicationInsert(
			context.Background(),
			item.dto,
		)
		require.Empty(t, err)
	}

	//select applications
	table2 := []struct {
		name    string
		dto     *ApplicationSelectDTO
		handler func(applications []entity.Application, err error)
	}{
		{
			name: "OK",
			dto: &ApplicationSelectDTO{
				Limit:  0,
				Offset: 0,
			},
			handler: func(applications []entity.Application, err error) {
				require.Empty(t, err)
				require.NotEqual(t, 0, len(applications))
			},
		},

		{
			name: "OK",
			dto: &ApplicationSelectDTO{
				Limit:  0,
				Offset: 500,
			},
			handler: func(applications []entity.Application, err error) {
				require.Empty(t, err)
				require.Equal(t, 0, len(applications))
			},
		},

		{
			name: "OK",
			dto: &ApplicationSelectDTO{
				Limit:  1,
				Offset: 0,
			},
			handler: func(applications []entity.Application, err error) {
				require.Empty(t, err)
				require.Equal(t, 1, len(applications))
			},
		},
	}

	for _, item := range table2 {
		t.Run(item.name, func(t *testing.T) {
			applications, err := repo.ApplicationSelect(
				context.Background(),
				item.dto,
			)
			item.handler(applications, err)
		})
	}
}

func TestRepo_ApplicationGet(t *testing.T) {
	repo := newRepo()

	//create new applications
	table := []struct {
		dto *ApplicationInsertDTO
	}{
		{
			dto: &ApplicationInsertDTO{
				KeyID: "tg-728688797819478",
				Tag:   "SquirrelHat",
			},
		},

		{
			dto: &ApplicationInsertDTO{
				KeyID: "tg-769654567373382",
				Tag:   "SquirrelHat",
			},
		},
	}

	for _, item := range table {
		{
			insertID, err := repo.ApplicationInsert(
				context.Background(),
				item.dto,
			)
			require.Empty(t, err)
			require.NotEmpty(t, insertID)
		}

		{
			application, err := repo.ApplicationGet(
				context.Background(),
				&ApplicationAcceptDTO{
					KeyID: item.dto.KeyID,
				},
			)
			require.Empty(t, err)
			require.Equal(t, item.dto.KeyID, application.KeyID)
			require.Equal(t, item.dto.Tag, application.Tag)
			require.Equal(t, false, application.Invited)
		}
	}
}
