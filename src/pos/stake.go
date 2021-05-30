package pos

type Stake struct{}

func NewStake() *Stake {
	return new(Stake)
}
