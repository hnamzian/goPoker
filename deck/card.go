package deck

import (
	"fmt"
	"strconv"
)

type Card struct {
	suit  Suit
	value int
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("card value cannot be more than 13")
	}

	return Card{
		suit:  s,
		value: v,
	}
}

func (c Card) String() string {
	value := strconv.Itoa(c.value)
	if c.value == 1 {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, c.suit.String(), c.suit.Unicode())
}


