package types

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cockroachdb/apd"
)

func MarshalDecimal(d apd.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte{'"'})
		_, _ = w.Write([]byte(d.String()))
		_, _ = w.Write([]byte{'"'})
	})
}

func UnmarshalDecimal(v interface{}) (apd.Decimal, error) {
	switch v := v.(type) {
	case string:
		d, _, err := apd.NewFromString(v)
		if err != nil {
			return apd.Decimal{}, fmt.Errorf("decimal format is invalid")
		}

		return *d, nil
	default:
		return apd.Decimal{}, fmt.Errorf("cannot parse %T", v)
	}
}
