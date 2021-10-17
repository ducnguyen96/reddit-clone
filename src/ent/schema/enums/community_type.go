package enums

import "database/sql/driver"

type CommunityType int

const (
	Public CommunityType = iota
	Restricted
	Private
)

func (p CommunityType) String() string {
	switch p {
	case Public:
		return "Public"
	case Restricted:
		return "Restricted"
	default:
		return "Private"
	}
}

// Values provides list valid values for Enum.
func (CommunityType) Values() []string {
	return []string{Public.String(), Restricted.String(), Private.String()}
}

// Value provides the DB a string from int.
func (p CommunityType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *CommunityType) Scan(val interface{}) error {
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
	case "Public":
		*p = Public
	case "Restricted":
		*p = Restricted
	default:
		*p = Private
	}
	return nil
}
