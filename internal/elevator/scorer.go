package elevator

type Scorer interface {
	Score(floor int, direction Direction) int
}
