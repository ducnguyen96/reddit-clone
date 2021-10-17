package enums

import "database/sql/driver"

type Gender int

const (
	Male Gender = iota
	Female
)

func (p Gender) String() string {
	switch p {
	case Female:
		return "FEMALE"
	default:
		return "MALE"
	}
}

// Values provides list valid values for Enum.
func (Gender) Values() []string {
	return []string{Male.String(), Female.String()}
}

// Value provides the DB a string from int.
func (p Gender) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *Gender) Scan(val interface{}) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}
	switch s {
	case "FEMALE":
		*p = Female
	default:
		*p = Male
	}
	return nil
}
