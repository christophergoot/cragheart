package datastore

import (
	"errors"
	"time"

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
	Timestamp int64
}

func (d Datastore) Get(id int) (*serve.Customer, error) {
	c, exists := d.Customers[id]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return &c, nil
}

func (d Datastore) List(page, count int) ([]*serve.Customer, error) {
	var list []*serve.Customer
	counter := 0
	first := (page - 1) * count
	last := first + count

	for id := range d.Customers {
		counter++
		if counter > last {
			break
		}
		if counter >= first {
			c := d.Customers[id]
			list = append(list, &c)
		}
	}

	return list, nil
}

func (m Datastore) Create(id int, attributes map[string]string) (*serve.Customer, error) {
	return m.createUser(id, attributes, time.Now().Unix())
}

func (m Datastore) Update(id int, attributes map[string]string) (*serve.Customer, error) {
	return m.updateUser(id, attributes, time.Now().Unix())
}

func (m Datastore) Delete(id int) error {
	_, exists := m.Customers[id]
	if !exists {
		return errors.New("customer not found")
	}
	delete(m.Customers, id)

	return nil
}

func (m Datastore) TotalCustomers() (int, error) {
	return len(m.Customers), nil
}
