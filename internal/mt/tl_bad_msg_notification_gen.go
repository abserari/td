// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// BadMsgNotification represents TL type `bad_msg_notification#a7eff811`.
type BadMsgNotification struct {
	// BadMsgID field of BadMsgNotification.
	BadMsgID int64
	// BadMsgSeqno field of BadMsgNotification.
	BadMsgSeqno int
	// ErrorCode field of BadMsgNotification.
	ErrorCode int
}

// BadMsgNotificationTypeID is TL type id of BadMsgNotification.
const BadMsgNotificationTypeID = 0xa7eff811

// String implements fmt.Stringer.
func (b *BadMsgNotification) String() string {
	if b == nil {
		return "BadMsgNotification(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BadMsgNotification")
	sb.WriteString("{\n")
	sb.WriteString("\tBadMsgID: ")
	sb.WriteString(fmt.Sprint(b.BadMsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tBadMsgSeqno: ")
	sb.WriteString(fmt.Sprint(b.BadMsgSeqno))
	sb.WriteString(",\n")
	sb.WriteString("\tErrorCode: ")
	sb.WriteString(fmt.Sprint(b.ErrorCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BadMsgNotification) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bad_msg_notification#a7eff811 as nil")
	}
	buf.PutID(BadMsgNotificationTypeID)
	buf.PutLong(b.BadMsgID)
	buf.PutInt(b.BadMsgSeqno)
	buf.PutInt(b.ErrorCode)
	return nil
}

// Decode implements bin.Decoder.
func (b *BadMsgNotification) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bad_msg_notification#a7eff811 to nil")
	}
	if err := buf.ConsumeID(BadMsgNotificationTypeID); err != nil {
		return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: %w", err)
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field bad_msg_id: %w", err)
		}
		b.BadMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field bad_msg_seqno: %w", err)
		}
		b.BadMsgSeqno = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_msg_notification#a7eff811: field error_code: %w", err)
		}
		b.ErrorCode = value
	}
	return nil
}

// construct implements constructor of BadMsgNotificationClass.
func (b BadMsgNotification) construct() BadMsgNotificationClass { return &b }

// Ensuring interfaces in compile-time for BadMsgNotification.
var (
	_ bin.Encoder = &BadMsgNotification{}
	_ bin.Decoder = &BadMsgNotification{}

	_ BadMsgNotificationClass = &BadMsgNotification{}
)

// BadServerSalt represents TL type `bad_server_salt#edab447b`.
type BadServerSalt struct {
	// BadMsgID field of BadServerSalt.
	BadMsgID int64
	// BadMsgSeqno field of BadServerSalt.
	BadMsgSeqno int
	// ErrorCode field of BadServerSalt.
	ErrorCode int
	// NewServerSalt field of BadServerSalt.
	NewServerSalt int64
}

// BadServerSaltTypeID is TL type id of BadServerSalt.
const BadServerSaltTypeID = 0xedab447b

// String implements fmt.Stringer.
func (b *BadServerSalt) String() string {
	if b == nil {
		return "BadServerSalt(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BadServerSalt")
	sb.WriteString("{\n")
	sb.WriteString("\tBadMsgID: ")
	sb.WriteString(fmt.Sprint(b.BadMsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tBadMsgSeqno: ")
	sb.WriteString(fmt.Sprint(b.BadMsgSeqno))
	sb.WriteString(",\n")
	sb.WriteString("\tErrorCode: ")
	sb.WriteString(fmt.Sprint(b.ErrorCode))
	sb.WriteString(",\n")
	sb.WriteString("\tNewServerSalt: ")
	sb.WriteString(fmt.Sprint(b.NewServerSalt))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (b *BadServerSalt) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bad_server_salt#edab447b as nil")
	}
	buf.PutID(BadServerSaltTypeID)
	buf.PutLong(b.BadMsgID)
	buf.PutInt(b.BadMsgSeqno)
	buf.PutInt(b.ErrorCode)
	buf.PutLong(b.NewServerSalt)
	return nil
}

// Decode implements bin.Decoder.
func (b *BadServerSalt) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bad_server_salt#edab447b to nil")
	}
	if err := buf.ConsumeID(BadServerSaltTypeID); err != nil {
		return fmt.Errorf("unable to decode bad_server_salt#edab447b: %w", err)
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bad_server_salt#edab447b: field bad_msg_id: %w", err)
		}
		b.BadMsgID = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_server_salt#edab447b: field bad_msg_seqno: %w", err)
		}
		b.BadMsgSeqno = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode bad_server_salt#edab447b: field error_code: %w", err)
		}
		b.ErrorCode = value
	}
	{
		value, err := buf.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bad_server_salt#edab447b: field new_server_salt: %w", err)
		}
		b.NewServerSalt = value
	}
	return nil
}

// construct implements constructor of BadMsgNotificationClass.
func (b BadServerSalt) construct() BadMsgNotificationClass { return &b }

// Ensuring interfaces in compile-time for BadServerSalt.
var (
	_ bin.Encoder = &BadServerSalt{}
	_ bin.Decoder = &BadServerSalt{}

	_ BadMsgNotificationClass = &BadServerSalt{}
)

// BadMsgNotificationClass represents BadMsgNotification generic type.
//
// Example:
//  g, err := DecodeBadMsgNotification(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BadMsgNotification: // bad_msg_notification#a7eff811
//  case *BadServerSalt: // bad_server_salt#edab447b
//  default: panic(v)
//  }
type BadMsgNotificationClass interface {
	bin.Encoder
	bin.Decoder
	construct() BadMsgNotificationClass
	fmt.Stringer
}

// DecodeBadMsgNotification implements binary de-serialization for BadMsgNotificationClass.
func DecodeBadMsgNotification(buf *bin.Buffer) (BadMsgNotificationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BadMsgNotificationTypeID:
		// Decoding bad_msg_notification#a7eff811.
		v := BadMsgNotification{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BadMsgNotificationClass: %w", err)
		}
		return &v, nil
	case BadServerSaltTypeID:
		// Decoding bad_server_salt#edab447b.
		v := BadServerSalt{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BadMsgNotificationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BadMsgNotificationClass: %w", bin.NewUnexpectedID(id))
	}
}

// BadMsgNotification boxes the BadMsgNotificationClass providing a helper.
type BadMsgNotificationBox struct {
	BadMsgNotification BadMsgNotificationClass
}

// Decode implements bin.Decoder for BadMsgNotificationBox.
func (b *BadMsgNotificationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BadMsgNotificationBox to nil")
	}
	v, err := DecodeBadMsgNotification(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BadMsgNotification = v
	return nil
}

// Encode implements bin.Encode for BadMsgNotificationBox.
func (b *BadMsgNotificationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BadMsgNotification == nil {
		return fmt.Errorf("unable to encode BadMsgNotificationClass as nil")
	}
	return b.BadMsgNotification.Encode(buf)
}
