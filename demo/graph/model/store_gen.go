package model

import (
	"time"
)

func (Article) ValsFromCols(cols []string) []interface{} {
	return []interface{}{
		new(string),
		new(time.Time),
		new(string),
		new(string),
	}
}

func (Article) ScanInto(v []interface{}) Article {
	obj := Article{
		ID:        *v[0].(*string),
		CreatedAt: *v[1].(*time.Time),
		Title:     *v[2].(*string),
		Content:   *v[3].(*string),
	}

	return obj
}
