package main

import (
	"testing"
)

func TestMain(t *testing.T) {
}

func TestMainPass(t *testing.T) {
	t.Error("Something bad happened")
}
