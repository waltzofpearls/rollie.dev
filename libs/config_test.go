package libs_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/rollie.dev/libs"
)

func TestEmptyConfig(t *testing.T) {
	libs.NewConfig()
}

func TestValidConfigFile(t *testing.T) {
	fixtures := []string{
		`{}`,
		`{"env":""}`,
		`{"env":"testing"}`,
		`{"listen":{}}`,
		`{"github":{}}`,
		`{"template":{}}`,
		`{"listen":{"http":"localhost:1234"}}`,
		`{"github":{"token":"xxxxxxxxxx"}}`,
		`{"github":{"username":"testuser"}}`,
		`{"template":{"path":"../"}}`,
	}

	for testn, fixture := range fixtures {
		func() {
			// Setup
			config, err := ioutil.TempFile("", "tetris-config")
			require.Nil(t, err)
			if os.Getenv("TEST_PRESERVE") == "" {
				defer os.Remove(config.Name())
			}

			_, err = config.WriteString(fixture)
			require.Nil(t, err)
			err = config.Close()
			require.Nil(t, err)

			// Verification
			var conf *libs.Config

			assert.NotPanics(t, func() {
				conf = libs.NewConfigFile(config.Name())
			}, "[%d:%s] parse errors", testn, fixture)
			assert.NotNil(t, conf, "[%d:%s] invalid config", testn, fixture)
		}()
	}
}

func TestEnvConfig(t *testing.T) {
	fixture := `{"env":"development"}`

	config, err := ioutil.TempFile("", "tetris-config")
	require.Nil(t, err)
	defer os.Remove(config.Name())

	_, err = config.WriteString(fixture)
	require.Nil(t, err)
	err = config.Close()
	require.Nil(t, err)

	os.Setenv("ENV_NAME", "testing")
	os.Setenv("GITHUB_TOKEN", "testToken")

	conf := libs.NewConfigFile(config.Name())

	assert.Equal(t, "testing", conf.Env)
	assert.Equal(t, "testToken", conf.Github.Token)
}
