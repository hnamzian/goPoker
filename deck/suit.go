package deck

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamond
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Hearts:
		return "HEARTS"
	case Diamond:
		return "DIAMOND"
	case Clubs:
		return "CLUBS"
	default:
		panic("invalid suit value")
	}
}

func (s Suit) Unicode() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamond:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid suit value")
	} 
}