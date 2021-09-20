package types

type Money int64

type Category string

type Currency string

type Status string

const (
	StatusOK         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type Payment struct {
	ID       int
	Amount   Money
	Status   Status
	Category Category
}
