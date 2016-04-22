package db

import (
	"github.com/jmoiron/modl"
	"github.com/manveru/faker"
)

// CreateFakePeople will populate the db with num fake person profiles.
func CreateFakePeople(dbm *modl.DbMap, num int) error {
	f, err := faker.New("en")
	if err != nil {
		return err
	}

	for i := 0; i < num; i++ {
		p := NewPerson(f.Name())

		p.EmailAddress = f.Email()
		p.PhoneMobile = f.CellPhoneNumber()
		p.PhoneWork = f.PhoneNumber()

		p.Comment = "fake profile"

		err := dbm.Insert(p)
		if err != nil {
			return err
		}
	}

	return nil
}
