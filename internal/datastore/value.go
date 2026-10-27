package datastore

type ValueType int

const (
	StringType ValueType = iota
	ListType
)

type Value struct {
	Type ValueType

	String string

	List *List
}
