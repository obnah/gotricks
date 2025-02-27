package must

import (
	"fmt"
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestSucceed(t *testing.T) {
	answerIs := func(n int) (int, error) {
		if n == 42 {
			return 42, nil
		}
		return -1, fmt.Errorf("%s", "the answer is 42!")
	}

	if 42 != Succeed(answerIs(42)) {
		t.Errorf("test failed")
	}

	assertPanic(t, func() { Succeed(answerIs(9)) })
}

func TestBeTrue(t *testing.T) {
	snowIs := func(s string) (string, bool) {
		if s == "white" {
			return "white", true
		}
		return "", false
	}

	if BeTrue(snowIs("white")) != "white" {
		t.Errorf("test failed")
	}

	assertPanic(t, func() { BeTrue(snowIs("black")) })
}
