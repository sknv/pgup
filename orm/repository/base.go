package repository

import (
	"upper.io/db.v3"

	"github.com/sknv/pgup/orm/record"
)

var MaxFetchLimit = 50

type (
	Base struct {
		Session        db.Database
		CollectionName string
	}

	PagingParams struct {
		Limit  int
		Offset int
		Order  []interface{}
	}
)

func (r *Base) GetCollection() db.Collection {
	return r.Session.Collection(r.CollectionName)
}

func (r *Base) Find(query ...interface{}) db.Result {
	col := r.GetCollection()
	return col.Find(query...)
}

func (r *Base) FindPage(params PagingParams, query ...interface{}) db.Result {
	res := r.Find(query...)

	// Set limit and offset params.
	limit := MaxFetchLimit
	if params.Limit > 0 && params.Limit < limit {
		limit = params.Limit // Restrict fetching limit.
	}
	res = res.Limit(limit)

	if params.Offset > 0 {
		res = res.Offset(params.Offset)
	}

	res = res.OrderBy(params.Order...)

	return res
}

func (r *Base) Insert(record interface{}) (interface{}, error) {
	col := r.GetCollection()

	// Before callbacks section.
	doBeforeInsertIfNeeded(record)
	doBeforeSaveIfNeeded(record)

	id, err := col.Insert(record)

	// After callbacks section.
	doAfterInsertIfNeeded(record)
	doAfterSaveIfNeeded(record)

	return id, err
}

func (r *Base) Update(record interface{}, query ...interface{}) error {
	res := r.Find(query...)

	// Before callbacks section.
	doBeforeUpdateIfNeeded(record)
	doBeforeSaveIfNeeded(record)

	err := res.Update(record)

	// After callbacks section.
	doAfterUpdateIfNeeded(record)
	doAfterSaveIfNeeded(record)

	return err
}

func (r *Base) UpdateRecord(record record.IIdentifier) error {
	return r.Update(record, "id", record.GetId())
}

func (r *Base) Delete(query ...interface{}) error {
	res := r.Find(query...)
	return res.Delete()
}

func (r *Base) DeleteRecord(record record.IIdentifier) error {
	res := r.Find("id", record.GetId())

	// Before callbacks section.
	doBeforeDeleteIfNeeded(record)

	err := res.Delete()

	// After callbacks section.
	doAfterDeleteIfNeeded(record)

	return err
}
