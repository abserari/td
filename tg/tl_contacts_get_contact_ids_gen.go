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

// ContactsGetContactIDsRequest represents TL type `contacts.getContactIDs#2caa4a42`.
// Get contact by telegram IDs
//
// See https://core.telegram.org/method/contacts.getContactIDs for reference.
type ContactsGetContactIDsRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// ContactsGetContactIDsRequestTypeID is TL type id of ContactsGetContactIDsRequest.
const ContactsGetContactIDsRequestTypeID = 0x2caa4a42

// String implements fmt.Stringer.
func (g *ContactsGetContactIDsRequest) String() string {
	if g == nil {
		return "ContactsGetContactIDsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsGetContactIDsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(g.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ContactsGetContactIDsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContactIDs#2caa4a42 as nil")
	}
	b.PutID(ContactsGetContactIDsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetContactIDsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContactIDs#2caa4a42 to nil")
	}
	if err := b.ConsumeID(ContactsGetContactIDsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetContactIDsRequest.
var (
	_ bin.Encoder = &ContactsGetContactIDsRequest{}
	_ bin.Decoder = &ContactsGetContactIDsRequest{}
)

// ContactsGetContactIDs invokes method contacts.getContactIDs#2caa4a42 returning error if any.
// Get contact by telegram IDs
//
// See https://core.telegram.org/method/contacts.getContactIDs for reference.
func (c *Client) ContactsGetContactIDs(ctx context.Context, hash int) ([]int, error) {
	var result IntVector

	request := &ContactsGetContactIDsRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
