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

// HelpRecentMeUrls represents TL type `help.recentMeUrls#e0310d7`.
// Recent t.me URLs
//
// See https://core.telegram.org/constructor/help.recentMeUrls for reference.
type HelpRecentMeUrls struct {
	// URLs
	Urls []RecentMeUrlClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
}

// HelpRecentMeUrlsTypeID is TL type id of HelpRecentMeUrls.
const HelpRecentMeUrlsTypeID = 0xe0310d7

// String implements fmt.Stringer.
func (r *HelpRecentMeUrls) String() string {
	if r == nil {
		return "HelpRecentMeUrls(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpRecentMeUrls")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range r.Urls {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range r.Chats {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range r.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *HelpRecentMeUrls) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode help.recentMeUrls#e0310d7 as nil")
	}
	b.PutID(HelpRecentMeUrlsTypeID)
	b.PutVectorHeader(len(r.Urls))
	for idx, v := range r.Urls {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field urls element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field urls element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Chats))
	for idx, v := range r.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(r.Users))
	for idx, v := range r.Users {
		if v == nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.recentMeUrls#e0310d7: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *HelpRecentMeUrls) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode help.recentMeUrls#e0310d7 to nil")
	}
	if err := b.ConsumeID(HelpRecentMeUrlsTypeID); err != nil {
		return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field urls: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeRecentMeUrl(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field urls: %w", err)
			}
			r.Urls = append(r.Urls, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field chats: %w", err)
			}
			r.Chats = append(r.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.recentMeUrls#e0310d7: field users: %w", err)
			}
			r.Users = append(r.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpRecentMeUrls.
var (
	_ bin.Encoder = &HelpRecentMeUrls{}
	_ bin.Decoder = &HelpRecentMeUrls{}
)
