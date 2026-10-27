package datastore

type ValueType int

const (
	StringType ValueType = iota
	ListType
	HashType
	SetType
)

type Value struct {
	Type   ValueType
	String string
	List   *List
	Hash   *Hash
	Set    *Set
}
