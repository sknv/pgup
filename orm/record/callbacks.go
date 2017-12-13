package record

type IBeforeDeleter interface {
	BeforeDelete()
}

type IBeforeInserter interface {
	BeforeInsert()
}

type IBeforeSaver interface {
	BeforeSave()
}

type IBeforeUpdater interface {
	BeforeUpdate()
}

type IAfterDeleter interface {
	AfterDelete()
}

type IAfterInserter interface {
	AfterInsert()
}

type IAfterSaver interface {
	AfterSave()
}

type IAfterUpdater interface {
	AfterUpdate()
}
