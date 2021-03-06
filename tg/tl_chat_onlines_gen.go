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

// ChatOnlines represents TL type `chatOnlines#f041e250`.
// Number of online users in a chat
//
// See https://core.telegram.org/constructor/chatOnlines for reference.
type ChatOnlines struct {
	// Number of online users
	Onlines int
}

// ChatOnlinesTypeID is TL type id of ChatOnlines.
const ChatOnlinesTypeID = 0xf041e250

// String implements fmt.Stringer.
func (c *ChatOnlines) String() string {
	if c == nil {
		return "ChatOnlines(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatOnlines")
	sb.WriteString("{\n")
	sb.WriteString("\tOnlines: ")
	sb.WriteString(fmt.Sprint(c.Onlines))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatOnlines) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatOnlines#f041e250 as nil")
	}
	b.PutID(ChatOnlinesTypeID)
	b.PutInt(c.Onlines)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatOnlines) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatOnlines#f041e250 to nil")
	}
	if err := b.ConsumeID(ChatOnlinesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatOnlines#f041e250: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatOnlines#f041e250: field onlines: %w", err)
		}
		c.Onlines = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatOnlines.
var (
	_ bin.Encoder = &ChatOnlines{}
	_ bin.Decoder = &ChatOnlines{}
)
