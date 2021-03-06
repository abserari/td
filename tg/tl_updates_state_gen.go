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

// UpdatesState represents TL type `updates.state#a56c2a3e`.
// Updates state.
//
// See https://core.telegram.org/constructor/updates.state for reference.
type UpdatesState struct {
	// Number of events occured in a text box
	Pts int
	// Position in a sequence of updates in secret chats. For further detailes refer to article secret chats¹Parameter was added in eigth layer².
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end
	//  2) https://core.telegram.org/api/layers#layer-8
	Qts int
	// Date of condition
	Date int
	// Number of sent updates
	Seq int
	// Number of unread messages
	UnreadCount int
}

// UpdatesStateTypeID is TL type id of UpdatesState.
const UpdatesStateTypeID = 0xa56c2a3e

// String implements fmt.Stringer.
func (s *UpdatesState) String() string {
	if s == nil {
		return "UpdatesState(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UpdatesState")
	sb.WriteString("{\n")
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(s.Pts))
	sb.WriteString(",\n")
	sb.WriteString("\tQts: ")
	sb.WriteString(fmt.Sprint(s.Qts))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(s.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tSeq: ")
	sb.WriteString(fmt.Sprint(s.Seq))
	sb.WriteString(",\n")
	sb.WriteString("\tUnreadCount: ")
	sb.WriteString(fmt.Sprint(s.UnreadCount))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *UpdatesState) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode updates.state#a56c2a3e as nil")
	}
	b.PutID(UpdatesStateTypeID)
	b.PutInt(s.Pts)
	b.PutInt(s.Qts)
	b.PutInt(s.Date)
	b.PutInt(s.Seq)
	b.PutInt(s.UnreadCount)
	return nil
}

// Decode implements bin.Decoder.
func (s *UpdatesState) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode updates.state#a56c2a3e to nil")
	}
	if err := b.ConsumeID(UpdatesStateTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.state#a56c2a3e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field pts: %w", err)
		}
		s.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field qts: %w", err)
		}
		s.Qts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field seq: %w", err)
		}
		s.Seq = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.state#a56c2a3e: field unread_count: %w", err)
		}
		s.UnreadCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesState.
var (
	_ bin.Encoder = &UpdatesState{}
	_ bin.Decoder = &UpdatesState{}
)
