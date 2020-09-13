package blackjack

import (
  "fmt"
  "github.com/jtschelling/blackjack/pkg/deck"
  "github.com/jtschelling/blackjack/pkg/hand"
  "github.com/jtschelling/blackjack/pkg/player"
)

type Game struct {
  Deck deck.Deck
  Players []player.Player
  PlayerPTR int
}

var blackjackStartingHandSize = 2

func (g Game) Play() {
  fmt.Println(g)
}

func New(numPlayers int) Game {
  players := make([]player.Player, numPlayers)
  for i := 0; i < numPlayers; i++ {
    players[i] = player.New(i, false)
  }

  players = chooseDealer(players)

  return Game {
    Deck: deck.Shuffle(deck.New([]string{}, 0), "random"),
    Players: players,
    PlayerPTR: 0,
  }
}

func Deal(g Game) Game {
  for i := 0; i < len(g.Players) * blackjackStartingHandSize; i++ {
    g.Deck, g.Players[i % len(g.Players)].Hand = hand.Draw(g.Deck, g.Players[i % len(g.Players)].Hand, 1)
  }

  return g
}

/////////////
// PRIVATE //
/////////////

func chooseDealer(p []player.Player) []player.Player {
  return player.SetDealer(len(p) - 1, p)
}
