package libs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/rolli3.net/libs"
)

func TestCreateApp(t *testing.T) {
	config := libs.NewConfig()
	config.Template.Path = "../"
	require.NotNil(t, config)

	app := libs.NewApp(config)
	assert.NotNil(t, app)
}
