package message

type Message interface {
	Decode([]byte)

	Encode() []byte

	GetPayload() []byte
}
