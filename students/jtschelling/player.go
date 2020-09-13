// Defines operation on a player
package player

import (
  "github.com/jtschelling/blackjack/students/jtschelling/deck"
  "github.com/jtschelling/blackjack/students/jtschelling/hand"
)

type Player struct {
  Hand hand.Hand
  Dealer bool
  Position int
}

func New() Player {
  return Player {
    Hand: hand.Hand {
      Cards: []deck.Card {
        deck.Card {
          Value: 0,
          Suit: "Heart",
        },
      },
    },
    Dealer: false,
  }
}
