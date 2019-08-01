package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	main "github.com/waltzofpearls/rollie.dev"
	"github.com/waltzofpearls/rollie.dev/libs"
)

func TestCreateApp(t *testing.T) {
	config := libs.NewConfig()
	config.Template.Path = "./"
	require.NotNil(t, config)

	app := main.NewApp(config)
	assert.NotNil(t, app)
}
