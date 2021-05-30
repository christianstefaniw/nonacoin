package pos

type stake struct{}

func NewStake() *stake {
	return new(stake)
}
