package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	op := calculate(3, 5, "+")
	var res int64
	res = 8
	assert.Equal(t, op, res)
}

func TestRemove(t *testing.T) {
	op := calculate(3, 5, "-")
	var res int64
	res = -2
	assert.Equal(t, op, res)
}

func TestError(t *testing.T) {
	assert.Panics(t, func() {
		_ = calculate(3, 5, "x")
	})
}

func TestMainFunc(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "3", "+", "5"}
	main()
}

func TestMainFuncFails(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "3", "/", "5"}
	assert.Panics(t, func() {
		main()
	})
}

func TestMainWrongParams(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"app", "3", "5"}
	assert.Panics(t, func() {
		main()
	})
}
