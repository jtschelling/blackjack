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
  Round int
}

var blackjackStartingHandSize = 2

func (g Game) Play() {
  g = newRound(g)
  findWinner(g.Players)
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
    Round: 1,
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

func findWinner(ps []player.Player) player.Player {
  var highScore int
  var pos int
  for ndx, p := range ps {
    fmt.Println(p)
    hand.Show(p.Hand)
    if hand.Points(p.Hand) > highScore {
      highScore = hand.Points(p.Hand)
      pos = ndx
    }
  }

  return ps[pos]
}

func newRound(g Game) Game {
  var d deck.Deck
  var h hand.Hand
  for _, p := range g.Players {
    d, h = DiscardHand(d, p.Hand)
  }


  return Game {
    Deck: g.Deck,
    Players: g.Players,
    PlayerPTR: 0,
    Round: (g.Round + 1),
  }
}
