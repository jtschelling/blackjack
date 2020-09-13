package blackjack

import (
  "testing"
)

func TestNew(t *testing.T) {
  numPlayers := 2
  g := New(numPlayers)

  if len(g.Deck.Cards) != 52 {
    t.Errorf("Deck is not the correct number of cards.")
  }

  if g.Players[numPlayers - 1].Dealer != true {
    t.Errorf("Dealer is not the last player in the list.")
  }
}

func TestDeal(t *testing.T) {
  numPlayers := 3
  g := New(numPlayers)

  g = Deal(g)

  for i := 0; i < numPlayers; i++ {
    if len(g.Players[i].Hand.Cards) != 2 {
      t.Errorf("Players were not all dealt the correct number of cards.")
    }
  }
}
