package database

import (
	"fmt"
	"os"
	"testing"
)

var (
	db Database
)

func TestMain(m *testing.M) {
	db = New()
	fmt.Println("Test")
	os.Exit(m.Run())
}