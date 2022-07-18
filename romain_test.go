package numlang_test

import (
	"testing"

	"github.com/jakofys/numlang"
)

func TestParser(t *testing.T) {
	if numlang.Parse(1840) != "MDCCCXL" {
		t.Error()
	}
	if numlang.Parse(1) != "I" {
		t.Error()
	}
	if numlang.Parse(267) != "CCLXVII" {
		t.Error()
	}
}
func TestInt(t *testing.T) {
	if n := numlang.Int("VVVVV"); n != 20 {
		t.Error(n)
	}
	if numlang.Int("CCLXVII") != 267 {
		t.Error()
	}
}
