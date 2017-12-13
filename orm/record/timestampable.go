package record

import "time"

type Timestampable struct {
	Base `db:",inline"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (t *Timestampable) BeforeInsert() {
	t.CreatedAt = time.Now()
}

func (t *Timestampable) BeforeSave() {
	t.UpdatedAt = time.Now()
}
