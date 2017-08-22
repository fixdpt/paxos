package paxos

type monotonic_value struct {
	uuid string
	n    uint64
}

func New() *identifier {
	return &monotonic_value{"", 1}
}
