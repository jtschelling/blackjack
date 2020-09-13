package player

import (
  "testing"
)

// Test New() function and expects a
func TestNew(t *testing.T) {
  p := New()

  if p.Hand.Cards[0].Value != 0 {
    t.Errorf("Player's hand does not contain a hand.")
  }
}
