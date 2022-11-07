package room

func Open(id string) Room {
	var r Room

	return r
}

type Member interface {
	ID() string
	Notify() chan<- Message
}

type Room interface {
	ID() string
	Name() string
	Members() []Member
	Pull(id, scope uint64) ([]Message, error)
	Push(string) (Message, error)
	Notify(Member, Message) error
	Join(Member) error
	Leave(Member) error
}
