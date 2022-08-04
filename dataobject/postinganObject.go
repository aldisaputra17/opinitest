package dataobject

type PostCreateObj struct {
	ID          uint64
	Title       string
	Description string
	User_ID     uint64
}

type PostUpdatedObj struct {
	Title       string
	Description string
	User_ID     uint64
}

type PostDeletedObj struct {
	ID          uint64
	Title       string
	Description string
	User_ID     uint64
}
