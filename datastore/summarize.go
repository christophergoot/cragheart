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

	if rec.Type == "attribute" {
		if _, exists := d.Customers[userID]; !exists {
			_, err = d.Create(userID, rec.Data)
		} else {
			_, err = d.Update(userID, rec.Data)
		}
	}

	if rec.Type == "event" {
		if _, exists := d.EventLog[rec.ID]; exists { //event previously recorded, no need to proceed
			return nil
		}

		d.EventLog[rec.ID] = EventData{ID: rec.ID, Name: rec.Name, UserID: userID, Data: rec.Data, Timestamp: int(rec.Timestamp)}

		c, exists := d.Customers[userID]
		if !exists { // create customer if not yet encountered
			newCustomer, err := d.Create(userID, rec.Data)
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
