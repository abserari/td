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

// MaskCoords represents TL type `maskCoords#aed6dbb2`.
// Position on a photo where a mask should be placed
// The n position indicates where the mask should be placed:
//
// See https://core.telegram.org/constructor/maskCoords for reference.
type MaskCoords struct {
	// Part of the face, relative to which the mask should be placed
	N int
	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	X float64
	// Shift by Y-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
	Y float64
	// Mask scaling coefficient. (For example, 2.0 means a doubled size)
	Zoom float64
}

// MaskCoordsTypeID is TL type id of MaskCoords.
const MaskCoordsTypeID = 0xaed6dbb2

// String implements fmt.Stringer.
func (m *MaskCoords) String() string {
	if m == nil {
		return "MaskCoords(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MaskCoords")
	sb.WriteString("{\n")
	sb.WriteString("\tN: ")
	sb.WriteString(fmt.Sprint(m.N))
	sb.WriteString(",\n")
	sb.WriteString("\tX: ")
	sb.WriteString(fmt.Sprint(m.X))
	sb.WriteString(",\n")
	sb.WriteString("\tY: ")
	sb.WriteString(fmt.Sprint(m.Y))
	sb.WriteString(",\n")
	sb.WriteString("\tZoom: ")
	sb.WriteString(fmt.Sprint(m.Zoom))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *MaskCoords) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode maskCoords#aed6dbb2 as nil")
	}
	b.PutID(MaskCoordsTypeID)
	b.PutInt(m.N)
	b.PutDouble(m.X)
	b.PutDouble(m.Y)
	b.PutDouble(m.Zoom)
	return nil
}

// Decode implements bin.Decoder.
func (m *MaskCoords) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode maskCoords#aed6dbb2 to nil")
	}
	if err := b.ConsumeID(MaskCoordsTypeID); err != nil {
		return fmt.Errorf("unable to decode maskCoords#aed6dbb2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field n: %w", err)
		}
		m.N = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field x: %w", err)
		}
		m.X = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field y: %w", err)
		}
		m.Y = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode maskCoords#aed6dbb2: field zoom: %w", err)
		}
		m.Zoom = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MaskCoords.
var (
	_ bin.Encoder = &MaskCoords{}
	_ bin.Decoder = &MaskCoords{}
)
