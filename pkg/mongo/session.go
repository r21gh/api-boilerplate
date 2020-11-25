package mongo

import (
	"gopkg.in/mgo.v2"
)

// Session struct for mongo connection
type Session struct {
	session *mgo.Session
}

// NewSession creates new connection
func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	return &Session{session}, err
}

// Copy returns a copy of session
func(s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

// GetCollection returns collection of
// speific db and collection name
func(s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

// Close is responsible for closing a specific session
func (s *Session) Close() {
	if(s.session != nil) {
		s.session.Close()
	}
}


// DropDatabase is responsible for
// dropping the specified database
func (s *Session) DropDatabase(db string) error {
	if s.session != nil {
		return s.session.DB(db).DropDatabase()
	}
	return nil
}