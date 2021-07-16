package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("say hello with name", func(t *testing.T) {
		assert.Equal(t, "Hello, Chris!", Hello("Chris", ""))
	})

	t.Run("say hello to world, when name is not present", func(t *testing.T) {
		assert.Equal(t, "Hello, World!", Hello("", ""))
	})

	t.Run("say hello to world in spanish", func(t *testing.T) {
		assert.Equal(t, "Hola, World!", Hello("", spanish))
	})

	t.Run("say hello to world in lithuanian", func(t *testing.T) {
		assert.Equal(t, "Labas, World!", Hello("", lithuanian))
	})
}
