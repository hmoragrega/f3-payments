package persistence

import (
	"errors"
	"sync"
	"time"

	"github.com/globalsign/mgo"
	"github.com/hmoragrega/f3-payments/pkg/logging"
)

const timeout = 60

var (
	// ErrConnectionMongo is triggered when the connection to mongo fails
	ErrConnectionMongo = errors.New("The connection to mongo server has failed")

	// single connection instance shared accross all mongo entity repositories
	mongoDB *mgo.Database
)

// MongoEntity is a entity that can be manged and persisted in mongo
// name: Name of the entity
// one: A factory that creates an empty entity to be filled by the repository
// list: A factory that creates an empty collection to be filled by the repository
type MongoEntity struct {
	collection string
	one        func() interface{}
	list       func() interface{}
}

// NewMongoEntity factory method to creates a new mongo entity
func NewMongoEntity(collection string, one func() interface{}, list func() interface{}) *MongoEntity {
	return &MongoEntity{collection, one, list}
}

// MongoConfig configuration parameters for connecting to mongo
type MongoConfig struct {
	Database string
	Address  string
	AuthDB   string
	User     string
	Pass     string
}

// MongoRepository mongo DB implementation of repository
type MongoRepository struct {
	entity *MongoEntity
}

// NewMongoRepository factory method to get a new mongo repository
func NewMongoRepository(config MongoConfig, entity *MongoEntity) (Repository, error) {
	err := connect(config)
	if err != nil {
		return nil, err
	}

	return &MongoRepository{entity}, nil
}

// Persist persists an entity and returns the id
func (m *MongoRepository) Persist(i interface{}) error {
	return m.collection().Insert(i)
}

// List returns a collection of entities
func (m *MongoRepository) List() (interface{}, error) {
	l := m.entity.list()
	err := m.collection().Find(nil).Sort("_id").All(l)

	return l, err
}

// Get retrieves a single entity by the ID
func (m *MongoRepository) Get(ID string) (interface{}, error) {
	e := m.entity.one()
	err := m.collection().FindId(ID).One(e)

	if err != nil && err == mgo.ErrNotFound {
		return nil, nil
	}

	return e, err
}

// Delete deletes an entity by the ID
func (m *MongoRepository) Delete(ID string) error {
	err := m.collection().RemoveId(ID)

	if err != nil && err == mgo.ErrNotFound {
		return nil
	}

	return err
}

func (m *MongoRepository) collection() *mgo.Collection {
	return mongoDB.C(m.entity.collection)
}

func connect(config MongoConfig) error {
	mtx := &sync.RWMutex{}
	mtx.RLock()
	defer mtx.RUnlock()

	if mongoDB == nil {
		info := &mgo.DialInfo{
			Addrs:    []string{config.Address},
			Database: config.AuthDB,
			Username: config.User,
			Password: config.Pass,
			Timeout:  timeout * time.Second,
		}

		session, err := mgo.DialWithInfo(info)
		if err != nil {
			return logging.Errors(ErrConnectionMongo, err)
		}

		mongoDB = session.DB(config.Database)
	}

	return nil
}
