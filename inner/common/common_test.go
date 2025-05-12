package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_getConfig_withoutNeededFile(t *testing.T) {
	config := GetConfig(".enx")
	assert.Equal(t, "", config.Dsn)
	assert.Equal(t, "", config.DbDriverName)
}

func Test_getConfig_withoutVars(t *testing.T) {

	config := GetConfig(("test.env"))
	assert.Equal(t, "", config.Dsn)
	assert.Equal(t, "", config.DbDriverName)
}

func Test_getConfig_withCorrectEnvFile(t *testing.T) {
	config := GetConfig(".env")

	assert.Equal(t, "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", config.Dsn)
	assert.Equal(t, "postgres", config.DbDriverName)
}

func Test_getConfig_withEnvironmentVars(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "test")
	os.Setenv("DB_DSN", "test")
	config := GetConfig("test.env")

	assert.Equal(t, "test", config.Dsn)
	assert.Equal(t, "test", config.DbDriverName)
}

func Test_getConfig_withEnvironmentVarsAndEnvFileVars(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "test")
	os.Setenv("DB_DSN", "test")
	config := GetConfig(".env")

	assert.Equal(t, "test", config.Dsn)
	assert.Equal(t, "test", config.DbDriverName)
}
