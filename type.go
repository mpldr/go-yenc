package yenc

type Encoder struct {
	LineLength int
	BlockSize  int
}

func NewEncoder() Encoder {
	return Encoder{
		LineLength: 128,
	}
}
