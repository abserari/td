// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// SecureValueVector is a box for Vector<SecureValue>
type SecureValueVector struct {
	// Elements of Vector<SecureValue>
	Elems []SecureValue
}

// String implements fmt.Stringer.
func (vec *SecureValueVector) String() string {
	if vec == nil {
		return "SecureValueVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("SecureValueVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// Encode implements bin.Encoder.
func (vec *SecureValueVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<SecureValue> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<SecureValue>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *SecureValueVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<SecureValue> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<SecureValue>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SecureValue
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<SecureValue>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValueVector.
var (
	_ bin.Encoder = &SecureValueVector{}
	_ bin.Decoder = &SecureValueVector{}
)
