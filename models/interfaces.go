package models

type ProtoMessage[T any] interface {
	Proto() *T
}

type ProtoEnum[T any] interface {
	Proto() T
}
