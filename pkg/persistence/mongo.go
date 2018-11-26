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
	session *mgo.Session
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
	config MongoConfig
	entity *MongoEntity
}

// NewMongoRepository factory method to get a new mongo repository
func NewMongoRepository(config MongoConfig, entity *MongoEntity) (*MongoRepository, error) {
	err := connect(config)
	if err != nil {
		return nil, err
	}

	return &MongoRepository{config, entity}, nil
}

// Persist persists an entity and returns the id
func (m *MongoRepository) Persist(i interface{}) error {
	return m.collection().Insert(i)
}

// Update updates an existing entity
func (m *MongoRepository) Update(ID string, i interface{}) error {
	return m.collection().UpdateId(ID, i)
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

// CloseMongoSession closes the session to mongo
func CloseMongoSession() {
	if session != nil {
		session.Close()
	}
}

func (m *MongoRepository) collection() *mgo.Collection {
	return session.DB(m.config.Database).C(m.entity.collection)
}

func connect(config MongoConfig) error {
	mtx := &sync.RWMutex{}
	mtx.RLock()
	defer mtx.RUnlock()

	if session == nil {
		info := &mgo.DialInfo{
			Addrs:    []string{config.Address},
			Database: config.AuthDB,
			Username: config.User,
			Password: config.Pass,
			Timeout:  timeout * time.Second,
		}

		var err error
		session, err = mgo.DialWithInfo(info)
		if err != nil {
			return logging.Errors(ErrConnectionMongo, err)
		}

		session.SetMode(mgo.Monotonic, true)
	}

	return nil
}
