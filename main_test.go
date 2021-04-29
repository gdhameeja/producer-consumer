package main

import (
	"testing"
)

func TestMain(t *testing.T) {
}

func TestMainFail(t *testing.T) {
	t.Error("This is supposed to fail")
}
