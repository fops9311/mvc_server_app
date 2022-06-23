package tassets

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting testing tassets")
	InitDir("./testtemplates")
	exitcode := m.Run()

	os.Exit(exitcode)
}
func TestGetAsset(t *testing.T) {
	fmt.Println(templates)

	var expected string = "dd"
	got := GetAsset("testtemplates/testfile1.txt")
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
	t.Logf("%s", templates)
}
