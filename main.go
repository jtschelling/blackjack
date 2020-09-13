package main

import (
  "github.com/jtschelling/blackjack/pkg/blackjack"
)

func main() {
  // Start game loop
  blackjack.New(2).Play()
}
