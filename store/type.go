package store

// Type ...
type Type struct {
	Scalars      []string
	Relationhips []string
}

// HasField ...
func (t Type) HasField(n string) bool {
	for _, s := range t.Scalars {
		if n == s {
			return true
		}
	}

	for _, r := range t.Relationhips {
		if n == r {
			return true
		}
	}

	return false
}

// HasScalar ...
func (t Type) HasScalar(n string) bool {
	for _, s := range t.Scalars {
		if n == s {
			return true
		}
	}

	return false
}

// HasRelationship ...
func (t Type) HasRelationship(n string) bool {
	for _, r := range t.Relationhips {
		if n == r {
			return true
		}
	}

	return false
}
