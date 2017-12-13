package record

type (
	IIdentifier interface {
		GetId() int64
	}

	Base struct {
		Id int64 `db:"id,omitempty" json:"id"`
	}
)

func (r *Base) GetId() int64 {
	return r.Id
}
