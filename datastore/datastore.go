package datastore

import (
	"errors"

	"github.com/customerio/homework/serve"
)

type Datastore struct {
	Customers map[int]serve.Customer
	EventLog  map[string]EventData
}

type EventData struct {
	ID        string
	Name      string
	UserID    int
	Data      map[string]string
	Timestamp int
}

func (d Datastore) Get(id int) (*serve.Customer, error) {
	return nil, errors.New("unimplemented")
}

func (d Datastore) List(page, count int) ([]*serve.Customer, error) {
	return nil, errors.New("unimplemented")
}

func (m Datastore) Create(id int, attributes map[string]string) (*serve.Customer, error) {
	return nil, errors.New("unimplemented")
}

func (m Datastore) Update(id int, attributes map[string]string) (*serve.Customer, error) {
	return nil, errors.New("unimplemented")
}

func (m Datastore) Delete(id int) error {
	return errors.New("unimplemented")
}

func (m Datastore) TotalCustomers() (int, error) {
	return 0, errors.New("unimplemented")
}
