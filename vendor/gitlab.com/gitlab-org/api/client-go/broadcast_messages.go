//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

type (
	BroadcastMessagesServiceInterface interface {
		ListBroadcastMessages(opt *ListBroadcastMessagesOptions, options ...RequestOptionFunc) ([]*BroadcastMessage, *Response, error)
		GetBroadcastMessage(broadcast int, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error)
		CreateBroadcastMessage(opt *CreateBroadcastMessageOptions, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error)
		UpdateBroadcastMessage(broadcast int, opt *UpdateBroadcastMessageOptions, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error)
		DeleteBroadcastMessage(broadcast int, options ...RequestOptionFunc) (*Response, error)
	}

	// BroadcastMessagesService handles communication with the broadcast
	// messages methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/broadcast_messages/
	BroadcastMessagesService struct {
		client *Client
	}
)

var _ BroadcastMessagesServiceInterface = (*BroadcastMessagesService)(nil)

// BroadcastMessage represents a GitLab broadcast message.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#get-all-broadcast-messages
type BroadcastMessage struct {
	Message            string             `json:"message"`
	StartsAt           *time.Time         `json:"starts_at"`
	EndsAt             *time.Time         `json:"ends_at"`
	Font               string             `json:"font"`
	ID                 int                `json:"id"`
	Active             bool               `json:"active"`
	TargetAccessLevels []AccessLevelValue `json:"target_access_levels"`
	TargetPath         string             `json:"target_path"`
	BroadcastType      string             `json:"broadcast_type"`
	Dismissable        bool               `json:"dismissable"`
	Theme              string             `json:"theme"`

	// Deprecated: This parameter was removed in GitLab 15.6.
	Color string `json:"color"`
}

// ListBroadcastMessagesOptions represents the available ListBroadcastMessages()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#get-all-broadcast-messages
type ListBroadcastMessagesOptions ListOptions

// ListBroadcastMessages gets a list of all broadcasted messages.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#get-all-broadcast-messages
func (s *BroadcastMessagesService) ListBroadcastMessages(opt *ListBroadcastMessagesOptions, options ...RequestOptionFunc) ([]*BroadcastMessage, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "broadcast_messages", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var bs []*BroadcastMessage
	resp, err := s.client.Do(req, &bs)
	if err != nil {
		return nil, resp, err
	}

	return bs, resp, nil
}

// GetBroadcastMessage gets a single broadcast message.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#get-a-specific-broadcast-message
func (s *BroadcastMessagesService) GetBroadcastMessage(broadcast int, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error) {
	u := fmt.Sprintf("broadcast_messages/%d", broadcast)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	b := new(BroadcastMessage)
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}

// CreateBroadcastMessageOptions represents the available CreateBroadcastMessage()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#create-a-broadcast-message
type CreateBroadcastMessageOptions struct {
	Message            *string            `url:"message" json:"message"`
	StartsAt           *time.Time         `url:"starts_at,omitempty" json:"starts_at,omitempty"`
	EndsAt             *time.Time         `url:"ends_at,omitempty" json:"ends_at,omitempty"`
	Font               *string            `url:"font,omitempty" json:"font,omitempty"`
	TargetAccessLevels []AccessLevelValue `url:"target_access_levels,omitempty" json:"target_access_levels,omitempty"`
	TargetPath         *string            `url:"target_path,omitempty" json:"target_path,omitempty"`
	BroadcastType      *string            `url:"broadcast_type,omitempty" json:"broadcast_type,omitempty"`
	Dismissable        *bool              `url:"dismissable,omitempty" json:"dismissable,omitempty"`
	Theme              *string            `url:"theme,omitempty" json:"theme,omitempty"`

	// Deprecated: This parameter was removed in GitLab 15.6.
	Color *string `url:"color,omitempty" json:"color,omitempty"`
}

// CreateBroadcastMessage creates a message to broadcast.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#create-a-broadcast-message
func (s *BroadcastMessagesService) CreateBroadcastMessage(opt *CreateBroadcastMessageOptions, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "broadcast_messages", opt, options)
	if err != nil {
		return nil, nil, err
	}

	b := new(BroadcastMessage)
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}

// UpdateBroadcastMessageOptions represents the available CreateBroadcastMessage()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#update-a-broadcast-message
type UpdateBroadcastMessageOptions struct {
	Message            *string            `url:"message,omitempty" json:"message,omitempty"`
	StartsAt           *time.Time         `url:"starts_at,omitempty" json:"starts_at,omitempty"`
	EndsAt             *time.Time         `url:"ends_at,omitempty" json:"ends_at,omitempty"`
	Font               *string            `url:"font,omitempty" json:"font,omitempty"`
	TargetAccessLevels []AccessLevelValue `url:"target_access_levels,omitempty" json:"target_access_levels,omitempty"`
	TargetPath         *string            `url:"target_path,omitempty" json:"target_path,omitempty"`
	BroadcastType      *string            `url:"broadcast_type,omitempty" json:"broadcast_type,omitempty"`
	Dismissable        *bool              `url:"dismissable,omitempty" json:"dismissable,omitempty"`
	Theme              *string            `url:"theme,omitempty" json:"theme,omitempty"`

	// Deprecated: This parameter was removed in GitLab 15.6.
	Color *string `url:"color,omitempty" json:"color,omitempty"`
}

// UpdateBroadcastMessage update a broadcasted message.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#update-a-broadcast-message
func (s *BroadcastMessagesService) UpdateBroadcastMessage(broadcast int, opt *UpdateBroadcastMessageOptions, options ...RequestOptionFunc) (*BroadcastMessage, *Response, error) {
	u := fmt.Sprintf("broadcast_messages/%d", broadcast)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	b := new(BroadcastMessage)
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}

// DeleteBroadcastMessage deletes a broadcasted message.
//
// GitLab API docs:
// https://docs.gitlab.com/api/broadcast_messages/#delete-a-broadcast-message
func (s *BroadcastMessagesService) DeleteBroadcastMessage(broadcast int, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("broadcast_messages/%d", broadcast)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
