package types_test

import (
	"bytes"
	"testing"

	"github.com/cockroachdb/apd"
	"github.com/mfcochauxlaberge/gqlgen-extra/types"
)

func TestMarshalDecimal(t *testing.T) {
	// 1
	d, _, _ := apd.NewFromString("1")

	marshaler := types.MarshalDecimal(*d)

	buf := &bytes.Buffer{}
	marshaler.MarshalGQL(buf)

	if buf.String() != `"1"` {
		t.Errorf("expected %q, got %q", `"1"`, buf.String())
	}

	// 1.2
	d, _, _ = apd.NewFromString("1.2")

	marshaler = types.MarshalDecimal(*d)

	buf = &bytes.Buffer{}
	marshaler.MarshalGQL(buf)

	if buf.String() != `"1.2"` {
		t.Errorf("expected %q, got %q", `"1.2"`, buf.String())
	}
}

func TestUnmarshalDecimal(t *testing.T) {
	cmp := apd.New(0, 0)

	// 1
	dec, _ := types.UnmarshalDecimal("1")
	exp, _, _ := apd.NewFromString("1")

	_, err := apd.BaseContext.Cmp(cmp, &dec, exp)
	if err != nil {
		t.Errorf("Cmp call failed: %s", err)
	}

	if !cmp.IsZero() {
		t.Errorf("%v != %v", &dec, exp)
	}

	// 1.2
	dec, _ = types.UnmarshalDecimal("1.2")
	exp, _, _ = apd.NewFromString("1.2")

	_, err = apd.BaseContext.Cmp(cmp, &dec, exp)
	if err != nil {
		t.Errorf("Cmp call failed: %s", err)
	}

	if !cmp.IsZero() {
		t.Errorf("%v != %v", &dec, exp)
	}

	// Error (wrong format)
	dec, err = types.UnmarshalDecimal("invalid")

	if err == nil {
		t.Errorf("UnmarshalDecimal(\"invalid\") should fail")
	}

	// Error (wrong type)
	dec, err = types.UnmarshalDecimal(123)

	if err == nil {
		t.Errorf("UnmarshalDecimal(123) should fail")
	}
}
