package service

import (
	"context"
	"gold-panel/internal/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_ApplicationFormSave(t *testing.T) {
	service := newService(t)

	form := entity.ApplicationForm{
		ListenFrom: "forum",
	}

	err := service.ApplicationFormSave(
		context.Background(),
		&ApplicationFormSaveDTO{
			KeyID: "tg-5678973",
			Form:  &form,
		},
	)

	require.Empty(t, err)
}

func TestService_ApplicationFormGet(t *testing.T) {
	service := newService(t)

	keyID := "tg-5678973"

	form := entity.ApplicationForm{
		ListenFrom: "forum",
	}
	err := service.ApplicationFormSave(
		context.Background(),
		&ApplicationFormSaveDTO{
			KeyID: keyID,
			Form:  &form,
		},
	)
	require.Empty(t, err)

	form2, err := service.ApplicationFormGet(
		context.Background(),
		&ApplicationFormGetDTO{
			KeyID: keyID,
		},
	)
	require.Empty(t, err)
	require.Equal(t, form.ListenFrom, form2.ListenFrom)
}

func TestService_ApplicationFormDel(t *testing.T) {
	service := newService(t)

	keyID := "tg-5678973"

	//add form
	form := entity.ApplicationForm{
		ListenFrom: "forum",
	}
	err := service.ApplicationFormSave(
		context.Background(),
		&ApplicationFormSaveDTO{
			KeyID: keyID,
			Form:  &form,
		},
	)
	require.Empty(t, err)

	//del form
	err = service.ApplicationFormDel(
		context.Background(),
		&ApplicationFormDelDTO{
			KeyID: keyID,
		},
	)
	require.Empty(t, err)

	//get form
	_, err = service.ApplicationFormGet(
		context.Background(),
		&ApplicationFormGetDTO{
			KeyID: keyID,
		},
	)
	require.NotEmpty(t, err)
}
