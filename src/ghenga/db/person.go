package db

import (
	"encoding/json"
	"time"
)

// Person is a person in the database.
type Person struct {
	ID           int64
	Name         string
	EmailAddress string
	PhoneWork    string
	PhoneMobile  string
	PhoneOther   string

	Comment string

	ChangedAt time.Time
	CreatedAt time.Time
}

// PersonJSON is the JSON representation of a Person as returned by the API.
type PersonJSON struct {
	ID           int64
	Name         string            `json:"name"`
	EmailAddress string            `json:"email_address"`
	PhoneNumbers []PhoneNumberJSON `json:"phone_numbers"`

	Comment string `json:"comment"`

	ChangedAt string `json:"changed_at"`
	CreatedAt string `json:"created_at"`
}

// PhoneNumberJSON is the JSON representation of a phone number.
type PhoneNumberJSON struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// NewPerson returns a new person record with the timestamps set to the current
// time.
func NewPerson(name string) Person {
	ts := time.Now()

	return Person{
		Name:      name,
		CreatedAt: ts,
		ChangedAt: ts,
	}
}

const timeLayout = "2006-01-02T15:04:05-07:00"

// MarshalJSON returns the JSON representation of p.
func (p *Person) MarshalJSON() ([]byte, error) {
	jp := PersonJSON{
		ID:           p.ID,
		Name:         p.Name,
		EmailAddress: p.EmailAddress,

		Comment:   p.Comment,
		ChangedAt: p.ChangedAt.Format(timeLayout),
		CreatedAt: p.CreatedAt.Format(timeLayout),
	}

	var numbers []PhoneNumberJSON

	for _, num := range []struct{ t, n string }{
		{"work", p.PhoneWork}, {"mobile", p.PhoneMobile}, {"other", p.PhoneOther},
	} {
		if num.n == "" {
			continue
		}

		numbers = append(numbers, PhoneNumberJSON{
			Type:   "work",
			Number: p.PhoneWork,
		})
	}

	return json.Marshal(jp)
}
