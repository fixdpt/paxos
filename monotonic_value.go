package paxos

type monotonic_value struct {
	uuid string
	n    uint64
}

func New() *monotonic_value {
	return &monotonic_value{"", 1}
}
