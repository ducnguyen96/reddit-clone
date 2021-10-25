package enums

import "database/sql/driver"

type UserActionTargetType int

const (
	POST UserActionTargetType = iota
	COMMENT
)

func (p UserActionTargetType) String() string {
	switch p {
	case POST:
		return "POST"
	default:
		return "COMMENT"
	}
}

func (UserActionTargetType) Values() []string {
	return []string{POST.String(), COMMENT.String()}
}

// Value provides the DB a string from int.
func (p UserActionTargetType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *UserActionTargetType) Scan(val interface{}) error {
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
	case "POST":
		*p = POST
	default:
		*p = COMMENT
	}
	return nil
}