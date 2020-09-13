// Defines operation on a player
package player

import (
  "github.com/jtschelling/blackjack/pkg/deck"
  "github.com/jtschelling/blackjack/pkg/hand"
)

type Player struct {
  Hand hand.Hand
  Dealer bool
  Position int
}

func New(position int, dealer bool) Player {
  return Player {
    Hand: hand.Hand {
      Cards: make([]deck.Card, 0),
    },
    Dealer: dealer,
    Position: position,
  }
}

func SetDealer(pos int, ps []Player) []Player {
  var ndx int
  for i, p := range ps {
    if p.Position == pos {
      ndx = i
      break
    }
  }

  ps[ndx].Dealer = true
  return ps
}
