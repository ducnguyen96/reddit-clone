package enums

import "database/sql/driver"

type UserActionType int

const (
	UpVote UserActionType = iota
	DownVote
)

func (p UserActionType) String() string {
	switch p {
	case UpVote:
		return "UP_VOTES"
	default:
		return "DOWN_VOTES"
	}
}

// Values provides list valid values for Enum.
func (UserActionType) Values() []string {
	return []string{UpVote.String(), DownVote.String()}
}

// Value provides the DB a string from int.
func (p UserActionType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *UserActionType) Scan(val interface{}) error {
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
	case "UP_VOTES":
		*p = UpVote
	default:
		*p = DownVote
	}
	return nil
}