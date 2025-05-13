package tests

import (
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"os"
	"testing"
)

func Test_Connect_Success(t *testing.T) {
	assert.NotPanics(t, func() {
		database.ConnectDb()
	}, "Code panics")
}

func Test_Connect_Panic(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "test")

	assert.Panics(t, func() {
		database.ConnectDb()
	}, "The code did not panic")
}
