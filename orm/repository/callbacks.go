package repository

import "github.com/sknv/pgup/orm/record"

func doBeforeInsertIfNeeded(doc interface{}) {
	if beforeInserter, ok := doc.(record.IBeforeInserter); ok {
		beforeInserter.BeforeInsert()
	}
}

func doBeforeDeleteIfNeeded(doc interface{}) {
	if beforeDeleter, ok := doc.(record.IBeforeDeleter); ok {
		beforeDeleter.BeforeDelete()
	}
}

func doBeforeSaveIfNeeded(doc interface{}) {
	if beforeSaver, ok := doc.(record.IBeforeSaver); ok {
		beforeSaver.BeforeSave()
	}
}

func doBeforeUpdateIfNeeded(doc interface{}) {
	if beforeUpdater, ok := doc.(record.IBeforeUpdater); ok {
		beforeUpdater.BeforeUpdate()
	}
}

func doAfterInsertIfNeeded(doc interface{}) {
	if afterInserter, ok := doc.(record.IAfterInserter); ok {
		afterInserter.AfterInsert()
	}
}

func doAfterDeleteIfNeeded(doc interface{}) {
	if afterDeleter, ok := doc.(record.IAfterDeleter); ok {
		afterDeleter.AfterDelete()
	}
}

func doAfterSaveIfNeeded(doc interface{}) {
	if afterSaver, ok := doc.(record.IAfterSaver); ok {
		afterSaver.AfterSave()
	}
}

func doAfterUpdateIfNeeded(doc interface{}) {
	if afterUpdater, ok := doc.(record.IAfterUpdater); ok {
		afterUpdater.AfterUpdate()
	}
}
