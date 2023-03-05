package common

import (
	"testing"
)

func TestGetRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetRandomString(6))
	}
}
