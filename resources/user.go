package resources

type User struct {
	ID     int64
	Email  string
	Status string
	Roles  []Role
}

type Role struct {
	ID          int64
	Name        string
	Description string
}
