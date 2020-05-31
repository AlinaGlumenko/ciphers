package abstractions

type Encoder interface {
	Encode(string) (string, bool)
}

type Decoder interface {
	Decode(string) (string, bool)
}
