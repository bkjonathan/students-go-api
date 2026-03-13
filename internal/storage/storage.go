package storage

type Storage interface {
	SaveStudent(name string, email string, age int) (int64, error)
}
