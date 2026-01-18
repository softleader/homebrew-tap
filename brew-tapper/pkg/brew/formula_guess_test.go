package brew

import (
	"fmt"
	"os"
	"testing"
)

func TestFormula_Guess(t *testing.T) {
	os, _, name, version, err := guess("slctl-v1.2.3-darwin.tgz")
	if err != nil {
		t.Error(err)
		t.SkipNow()
	}
	if v := os; v != "darwin" {
		t.Errorf("os must be darwin, but got %s", v)
	}
	if v := name; v != "slctl" {
		t.Errorf("name must be slctl, but got %s", v)
	}
	if v := version; v != "v1.2.3" {
		t.Errorf("name must be v1.2.3, but got %s", v)
	}
}

func TestFormula_Guess2(t *testing.T) {
	path := "/Users/matt/tmp/nu/temp"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Skipf("path %q not exist", path)
	}
	f := &Formula{}
	f.Guess(path)
	fmt.Printf("%#v\n", f)
}
