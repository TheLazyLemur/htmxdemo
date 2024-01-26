package types

type ContextKey string

const (
	MyDataKey  ContextKey = "isAuthenticated"
	MyRouteKey ContextKey = "route"
)

type Transaction struct {
	ID      int64 `json:"id"`
	From    int64 `json:"from"`
	To      int64 `json:"to"`
	Amount  int64 `json:"amount"`
	Status  int   `json:"status"`
	Flagged bool  `json:"flagged"`
}
