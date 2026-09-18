package httpx

type Page[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

func NewPage[T any](items []T, total int64) Page[T] {
	if items == nil {
		items = []T{}
	}
	return Page[T]{Items: items, Total: total}
}
