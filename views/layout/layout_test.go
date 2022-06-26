package layout

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting testing tassets")
	exitcode := m.Run()
	os.Exit(exitcode)
}
func TestLayout(t *testing.T) {
	lo := "Lay <inner_content_1/><inner_content_2/><inner_content_3/> out"
	fmt.Println(lo)
	ic1 := "inner_content_1"
	fmt.Println(ic1)
	ic2 := "inner_content_2"
	fmt.Println(ic2)
	ic3 := "inner_content_3"
	fmt.Println(ic3)

	var expected string = "Lay inner_content_1inner_content_2inner_content_3 out"
	got := Layout(
		lo,
		ic1,
		ic2,
		ic3,
	)
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
	t.Logf("layout = %s", lo)
	t.Logf("result = %s", got)
}
