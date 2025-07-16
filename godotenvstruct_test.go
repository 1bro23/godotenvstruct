package godotenvstruct

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ConfigCustomName struct {
	HostCustom string `env:"Config__Host"`
	PortCustom string `env:"Config__Port"`
}

type Config struct {
	Host string
	Port string
}

func TestBind_WithCustomName(t *testing.T) {
	os.Setenv("TEST_PREFIX_Config__Host", "localhost")
	os.Setenv("TEST_PREFIX_Config__Port", "8080")

	var config ConfigCustomName
	err := Bind("TEST_PREFIX_", &config)

	assert.Nil(t, err)
	assert.Equal(t, "localhost", config.HostCustom)
	assert.Equal(t, "8080", config.PortCustom)
}

func TestBind_WithoutCustomName(t *testing.T) {
	os.Setenv("TEST_PREFIX_Config__Host", "localhost")
	os.Setenv("TEST_PREFIX_Config__Port", "8080")

	var config Config
	err := Bind("TEST_PREFIX_", &config)

	assert.Nil(t, err)
	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, "8080", config.Port)
}
