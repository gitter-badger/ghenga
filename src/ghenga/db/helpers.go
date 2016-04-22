package db

import "github.com/jmoiron/modl"

// CreateFakePeople will populate the db with num fake person profiles.
func CreateFakePeople(db *modl.DbMap, num int) error {
	for i := 0; i < num; i++ {
		sex := randomdata.Male
		if i%2 == 0 {
			sex = randomdata.Female
		}
		p := Person{
			Name:    randomdata.FirstName(sex) + " " + randomdata.LastName(),
			Comment: "fake user",
		}

		db.MustExec(`INSERT INTO people (name, comment) VALUES (?, ?)`, p.Name, p.Comment)
	}

	return nil
}
