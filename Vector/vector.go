package vector

type	Rank	int

type Vector interface {
	Size()
	Get(r Rank) (interface{}, bool)
	Put(r Rank, e interface{}) bool
	Insert(r Rank, e interface{}) bool
	Remove(r Rank) (interface{}, bool)
}