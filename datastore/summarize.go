package datastore

import (
	"strconv"

	"github.com/customerio/homework/serve"
	"github.com/customerio/homework/stream"
)

func NewDatastoreFromChannel(ch <-chan *stream.Record) (*Datastore, error) {
	ds := Datastore{}
	ds.Customers = make(map[int]serve.Customer)
	ds.EventLog = make(map[string]EventData)
	for rec := range ch {
		ds.ParseRecord(rec)
	}

	return &ds, nil
}

func (d Datastore) ParseRecord(rec *stream.Record) error {
	userID, err := strconv.Atoi(rec.UserID)
	if err != nil {
		return err
	}

	if rec.Type == "attributes" {
		if _, exists := d.Customers[userID]; !exists {
			_, err = d.createUser(userID, rec.Data, rec.Timestamp)
		} else {
			_, err = d.updateUser(userID, rec.Data, rec.Timestamp)
		}
	}

	if rec.Type == "event" {
		if _, exists := d.EventLog[rec.ID]; exists { //event previously recorded, no need to proceed
			return nil
		}

		d.EventLog[rec.ID] = EventData{ID: rec.ID, Name: rec.Name, UserID: userID, Data: rec.Data, Timestamp: rec.Timestamp}

		c, exists := d.Customers[userID]
		if !exists { // create customer if not yet encountered
			newCustomer, err := d.createUser(userID, make(map[string]string), rec.Timestamp)
			if err != nil {
				return err
			}
			c = *newCustomer
		}

		c.Events[rec.Name] = c.Events[rec.Name] + 1

		d.Customers[userID] = c
	}
	return err
}

func (m Datastore) createUser(id int, attributes map[string]string, timestamp int64) (*serve.Customer, error) {
	c := serve.Customer{ID: id, Attributes: attributes, LastUpdated: timestamp}
	c.Events = make(map[string]int)
	m.Customers[id] = c

	return &c, nil
}

func (m Datastore) updateUser(id int, attributes map[string]string, timestamp int64) (*serve.Customer, error) {
	c := m.Customers[id]
	if c.LastUpdated < timestamp {
		c.LastUpdated = timestamp
		for k, v := range c.Attributes {
			c.Attributes[k] = v
		}
	}
	m.Customers[id] = c

	return &c, nil
}
