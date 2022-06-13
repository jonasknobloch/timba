package tim

import "time"

type Record struct {
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Note  string `json:"node,omitempty"`
}

func NewRecord(start, end time.Time, note string) *Record {
	return &Record{
		Start: start.UnixMilli(),
		End:   end.UnixMilli(),
		Note:  note,
	}
}
