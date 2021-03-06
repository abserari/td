// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// TestVectorIntObject represents TL type `testVectorIntObject#f152999b`.
//
// See https://localhost:80/doc/constructor/testVectorIntObject for reference.
type TestVectorIntObject struct {
	// Vector of objects
	Value []TestInt
}

// TestVectorIntObjectTypeID is TL type id of TestVectorIntObject.
const TestVectorIntObjectTypeID = 0xf152999b

// String implements fmt.Stringer.
func (t *TestVectorIntObject) String() string {
	if t == nil {
		return "TestVectorIntObject(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TestVectorIntObject")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range t.Value {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TestVectorIntObject) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorIntObject#f152999b as nil")
	}
	b.PutID(TestVectorIntObjectTypeID)
	b.PutVectorHeader(len(t.Value))
	for idx, v := range t.Value {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode testVectorIntObject#f152999b: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorIntObject) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorIntObject#f152999b to nil")
	}
	if err := b.ConsumeID(TestVectorIntObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorIntObject#f152999b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorIntObject#f152999b: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestInt
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode testVectorIntObject#f152999b: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorIntObject.
var (
	_ bin.Encoder = &TestVectorIntObject{}
	_ bin.Decoder = &TestVectorIntObject{}
)
