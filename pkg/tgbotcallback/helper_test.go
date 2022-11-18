package tgbotcallback

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitCallback(t *testing.T) {
	data := "start|123|null"

	expCallback := "start"
	expPayload := "123|null"

	callback, payload := SplitCallback(data)

	require.Equal(t, expCallback, callback)
	require.Equal(t, expPayload, payload)
}
