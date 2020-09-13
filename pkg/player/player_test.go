package player

import (
  "testing"
)

// Test New() function and expects a
func TestNew(t *testing.T) {
  p := New(0, false)

  if len(p.Hand.Cards) != 0 {
    t.Errorf("Player's hand does not contain a card.")
  }
}
