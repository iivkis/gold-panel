package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	LoadFrom(".")
	fmt.Println(config)
	require.NotEmpty(t, config.Panel.BotToken)
}
