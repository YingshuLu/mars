package room

const (
	TextMsg = iota
)

func roomID(id int64) int32 {
	return int32(id >> 32)
}
func messageID(id int64) int32 {
	mask := (int64(1)<<32 - 1)
	return int32(id & mask)
}

type Message interface {
	ID() int64
	Type() int
	Sender() string
	Content() []byte
}

type message struct {
	id      int64
	typ     int
	sender  string
	content string
}
