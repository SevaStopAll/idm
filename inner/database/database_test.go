package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func Test_Connect_Success(t *testing.T) {
//	assert.NotPanics(t, func() {
//		ConnectDb()
//	}, "Code panics")
//}

func Test_Connect_Panic(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "test")

	assert.Panics(t, func() {
		ConnectDb()
	}, "The code did not panic")
}
