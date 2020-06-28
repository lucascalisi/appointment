package resources

type QueueStorage interface {
	getQueue() (Queue, error)
	addItem(item QueueItem) error
}

type Queue struct {
	Items []QueueItem
}

type QueueItem struct {
	ID          int64
	IDPatient   int64
	IDSpecialty int64
	Status      string
}
