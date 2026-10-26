package object_test

import (
	"os"
	"testing"

	testconfig "github.com/ganigeorgiev/go-git/v6/internal/test/config"
)

func TestMain(m *testing.M) {
	testconfig.RegisterDefault()
	os.Exit(m.Run())
}
