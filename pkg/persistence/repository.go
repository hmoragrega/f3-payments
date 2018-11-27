package persistence

// Repository is an general infrastructure service to persist entities
type Repository interface {
	Persist(i interface{}) error
	Update(ID string, i interface{}) error
	List() (interface{}, error)
	Get(ID string) (interface{}, error)
	Delete(ID string) error
	DeleteAll() error
}
