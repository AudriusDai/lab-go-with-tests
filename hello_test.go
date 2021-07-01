package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	actual := Hello("Chris")
	expect := "Hello, Chris!"

	assert.Equalf(t, expect, actual, "expected %q, got %q", expect, actual)
}
