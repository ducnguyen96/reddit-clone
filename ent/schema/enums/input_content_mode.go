package enums

import "database/sql/driver"

type InputContentMode int

const (
	MarkDown InputContentMode = iota
	TextEditor
)

func (p InputContentMode) String() string {
	switch p {
	case MarkDown:
		return "MarkDown"
	default:
		return "TextEditor"
	}
}

// Values provides list valid values for Enum.
func (InputContentMode) Values() []string {
	return []string{MarkDown.String(), TextEditor.String()}
}

// Value provides the DB a string from int.
func (p InputContentMode) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *InputContentMode) Scan(val interface{}) error {
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
	case "MarkDown":
		*p = MarkDown
	default:
		*p = TextEditor
	}
	return nil
}