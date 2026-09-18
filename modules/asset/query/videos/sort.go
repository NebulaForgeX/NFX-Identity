package videos

type SortField int

const (
	_ SortField = iota
	SortByCreatedAt
	SortByUpdatedAt
	SortFieldMax
)

var allowedSortFields = makeRangeSet(SortFieldMax)

func makeRangeSet(max SortField) map[SortField]struct{} {
	set := make(map[SortField]struct{}, int(max))
	for i := SortField(1); i < max; i++ {
		set[i] = struct{}{}
	}
	return set
}
