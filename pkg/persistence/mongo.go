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
	session  *MongoSession
	entity   *MongoEntity
	database string
}

// NewMongoRepository factory method to get a new mongo repository
func NewMongoRepository(session *MongoSession, entity *MongoEntity, database string) *MongoRepository {
	return &MongoRepository{session, entity, database}
}

// Persist persists an entity and returns the id
func (m *MongoRepository) Persist(i interface{}) error {
	s := m.session.copy()
	defer s.Close()

	return m.collection(s).Insert(i)
}

// Update updates an existing entity
func (m *MongoRepository) Update(ID string, i interface{}) error {
	s := m.session.copy()
	defer s.Close()

	return m.collection(s).UpdateId(ID, i)
}

// List returns a collection of entities
func (m *MongoRepository) List() (interface{}, error) {
	s := m.session.copy()
	defer s.Close()

	l := m.entity.list()
	err := m.collection(s).Find(nil).Sort("_id").All(l)

	return l, err
}

// Get retrieves a single entity by the ID
func (m *MongoRepository) Get(ID string) (interface{}, error) {
	s := m.session.copy()
	defer s.Close()

	e := m.entity.one()
	err := m.collection(s).FindId(ID).One(e)

	if err != nil && err == mgo.ErrNotFound {
		return nil, nil
	}

	return e, err
}

// Delete deletes an entity by the ID
func (m *MongoRepository) Delete(ID string) error {
	s := m.session.copy()
	defer s.Close()

	err := m.collection(s).RemoveId(ID)

	if err != nil && err == mgo.ErrNotFound {
		return nil
	}

	return err
}

// DeleteAll deletes all entities in the collection
func (m *MongoRepository) DeleteAll() error {
	s := m.session.copy()
	defer s.Close()

	_, err := m.collection(s).RemoveAll(nil)
	return err
}

func (m *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB(m.database).C(m.entity.collection)
}

// MongoSession A wrapper to handle the connection to mongo
type MongoSession struct {
	config  *MongoConfig
	session *mgo.Session
	sync.RWMutex
}

// NewMongoSession factory method to gest a connection to mongo
func NewMongoSession(config *MongoConfig) *MongoSession {
	return &MongoSession{
		config: config,
	}
}

// Connect connects to the mongo server
func (s *MongoSession) Connect() error {
	s.Lock()
	defer s.Unlock()

	info := &mgo.DialInfo{
		Addrs:    []string{s.config.Address},
		Database: s.config.AuthDB,
		Username: s.config.User,
		Password: s.config.Pass,
		Timeout:  timeout * time.Second,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return logging.Errors(ErrConnectionMongo, err)
	}

	s.session = session
	return nil
}

// Close closes the connection to mongo
func (s *MongoSession) Close() {
	s.Lock()
	defer s.Unlock()

	s.session.Close()
}

func (s *MongoSession) copy() *mgo.Session {
	s.RLock()
	defer s.RUnlock()

	return s.session.Copy()
}
