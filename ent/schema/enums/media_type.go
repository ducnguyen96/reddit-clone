package enums

import "database/sql/driver"

type MediaType int

const (
	Image MediaType = iota
	Video
)

func (p MediaType) String() string {
	switch p {
	case Image:
		return "Image"
	default:
		return "Video"
	}
}

// Values provides list valid values for Enum.
func (MediaType) Values() []string {
	return []string{Image.String(), Video.String()}
}

// Value provides the DB a string from int.
func (p MediaType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *MediaType) Scan(val interface{}) error {
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
	case "Image":
		*p = Image
	default:
		*p = Video
	}
	return nil
}