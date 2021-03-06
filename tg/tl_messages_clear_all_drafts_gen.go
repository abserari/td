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

// MessagesClearAllDraftsRequest represents TL type `messages.clearAllDrafts#7e58ee9c`.
// Clear all drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.clearAllDrafts for reference.
type MessagesClearAllDraftsRequest struct {
}

// MessagesClearAllDraftsRequestTypeID is TL type id of MessagesClearAllDraftsRequest.
const MessagesClearAllDraftsRequestTypeID = 0x7e58ee9c

// String implements fmt.Stringer.
func (c *MessagesClearAllDraftsRequest) String() string {
	if c == nil {
		return "MessagesClearAllDraftsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesClearAllDraftsRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *MessagesClearAllDraftsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearAllDrafts#7e58ee9c as nil")
	}
	b.PutID(MessagesClearAllDraftsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesClearAllDraftsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearAllDrafts#7e58ee9c to nil")
	}
	if err := b.ConsumeID(MessagesClearAllDraftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.clearAllDrafts#7e58ee9c: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesClearAllDraftsRequest.
var (
	_ bin.Encoder = &MessagesClearAllDraftsRequest{}
	_ bin.Decoder = &MessagesClearAllDraftsRequest{}
)

// MessagesClearAllDrafts invokes method messages.clearAllDrafts#7e58ee9c returning error if any.
// Clear all drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.clearAllDrafts for reference.
func (c *Client) MessagesClearAllDrafts(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &MessagesClearAllDraftsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
