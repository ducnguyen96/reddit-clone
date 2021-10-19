package enums

import "database/sql/driver"

type PostType int

const (
	Post PostType = iota
	Image_Video
	Link
)

func (p PostType) String() string {
	switch p {
	case Post:
		return "Post"
	case Image_Video:
		return "Image_Video"
	default:
		return "Link"
	}
}

// Values provides list valid values for Enum.
func (PostType) Values() []string {
	return []string{Post.String(), Image_Video.String(), Link.String()}
}

// Value provides the DB a string from int.
func (p PostType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *PostType) Scan(val interface{}) error {
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
	case "Post":
		*p = Post
	case "Image_Video":
		*p = Image_Video
	default:
		*p = Link
	}
	return nil
}