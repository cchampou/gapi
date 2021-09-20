package logger

import "testing"

func TestLog(t *testing.T) {
	num, _ := Log("")
	if num != 1 {
		t.Errorf("should return null")
	}
}
