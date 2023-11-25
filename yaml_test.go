package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	_, _ = os.Create("test.yaml")
	defer os.Remove("test.yaml")

	_ = os.WriteFile("test.yaml", []byte("port: 8080\nhost: localhost\ntimeout: 60"), 0644)

	c, err := GetConfig("test.yaml")
	assert.Nil(t, err)

	assert.Equal(t, 8080, c.Port)
	assert.Equal(t, "localhost", c.Host)
	assert.Equal(t, 60, c.Timeout)
}

func TestGetConfigNoFile(t *testing.T) {
	c, err := GetConfig("test.yaml")
	assert.NotNil(t, err)

	assert.Nil(t, c)
}

func TestGetConfigBadYaml(t *testing.T) {
	_, _ = os.Create("test.yaml")
	defer os.Remove("test.yaml")
	_ = os.WriteFile("test.yaml", []byte("port: str"), 0644)

	c, err := GetConfig("test.yaml")
	assert.NotNil(t, err)
	assert.Nil(t, c)

}
