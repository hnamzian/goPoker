package deck

import (
	"math/rand"
	"time"
)

type Deck [52]Card

func New() *Deck {
	var (
		nSuit  = 4
		nValue = 13
		d      Deck
	)

	x := 0
	for i := 0; i < nSuit; i++ {
		for j := 0; j < nValue; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return d.Shuffle()
}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	return d
}
