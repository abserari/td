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

// TestBytes represents TL type `testBytes#a422c4de`.
//
// See https://localhost:80/doc/constructor/testBytes for reference.
type TestBytes struct {
	// Bytes
	Value []byte
}

// TestBytesTypeID is TL type id of TestBytes.
const TestBytesTypeID = 0xa422c4de

// String implements fmt.Stringer.
func (t *TestBytes) String() string {
	if t == nil {
		return "TestBytes(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TestBytes")
	sb.WriteString("{\n")
	sb.WriteString("\tValue: ")
	sb.WriteString(fmt.Sprint(t.Value))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *TestBytes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testBytes#a422c4de as nil")
	}
	b.PutID(TestBytesTypeID)
	b.PutBytes(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestBytes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testBytes#a422c4de to nil")
	}
	if err := b.ConsumeID(TestBytesTypeID); err != nil {
		return fmt.Errorf("unable to decode testBytes#a422c4de: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode testBytes#a422c4de: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TestBytes.
var (
	_ bin.Encoder = &TestBytes{}
	_ bin.Decoder = &TestBytes{}
)
