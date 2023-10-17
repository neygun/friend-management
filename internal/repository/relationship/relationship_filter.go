package relationship

// Filter defines filtering options for relationship repo methods
type Filter struct {
	RequestorID int64
	TargetID    int64
	Type        string
}
