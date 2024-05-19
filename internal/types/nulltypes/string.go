package nulltypes

import "database/sql"

func NewNullString(s *string) sql.NullString {
	n := sql.NullString{
		String: "",
		Valid:  s != nil,
	}
	if s != nil {
		n.String = *s
	}
	return n
}
