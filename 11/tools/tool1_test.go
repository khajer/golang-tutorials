package tools

import (
	"testing"
)

func TestText(t *testing.T) {
	if CheckData() != "hello" {
		t.Errorf("expect error")
	}
}
