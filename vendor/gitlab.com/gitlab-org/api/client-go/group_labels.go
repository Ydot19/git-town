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
)

type (
	GroupLabelsServiceInterface interface {
		ListGroupLabels(gid interface{}, opt *ListGroupLabelsOptions, options ...RequestOptionFunc) ([]*GroupLabel, *Response, error)
		GetGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
		CreateGroupLabel(gid interface{}, opt *CreateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
		DeleteGroupLabel(gid interface{}, lid interface{}, opt *DeleteGroupLabelOptions, options ...RequestOptionFunc) (*Response, error)
		UpdateGroupLabel(gid interface{}, lid interface{}, opt *UpdateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
		SubscribeToGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error)
		UnsubscribeFromGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*Response, error)
	}

	// GroupLabelsService handles communication with the label related methods of the
	// GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/group_labels/
	GroupLabelsService struct {
		client *Client
	}
)

var _ GroupLabelsServiceInterface = (*GroupLabelsService)(nil)

// GroupLabel represents a GitLab group label.
//
// GitLab API docs: https://docs.gitlab.com/api/group_labels/
type GroupLabel Label

func (l GroupLabel) String() string {
	return Stringify(l)
}

// ListGroupLabelsOptions represents the available ListGroupLabels() options.
//
// GitLab API docs: https://docs.gitlab.com/api/group_labels/#list-group-labels
type ListGroupLabelsOptions struct {
	ListOptions
	WithCounts               *bool   `url:"with_counts,omitempty" json:"with_counts,omitempty"`
	IncludeAncestorGroups    *bool   `url:"include_ancestor_groups,omitempty" json:"include_ancestor_groups,omitempty"`
	IncludeDescendantGrouops *bool   `url:"include_descendant_groups,omitempty" json:"include_descendant_groups,omitempty"`
	OnlyGroupLabels          *bool   `url:"only_group_labels,omitempty" json:"only_group_labels,omitempty"`
	Search                   *string `url:"search,omitempty" json:"search,omitempty"`
}

// ListGroupLabels gets all labels for given group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#list-group-labels
func (s *GroupLabelsService) ListGroupLabels(gid interface{}, opt *ListGroupLabelsOptions, options ...RequestOptionFunc) ([]*GroupLabel, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/labels", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var l []*GroupLabel
	resp, err := s.client.Do(req, &l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// GetGroupLabel get a single label for a given group.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#get-a-single-group-label
func (s *GroupLabelsService) GetGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	label, err := parseID(lid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/labels/%s", PathEscape(group), PathEscape(label))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var l *GroupLabel
	resp, err := s.client.Do(req, &l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// CreateGroupLabelOptions represents the available CreateGroupLabel() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#create-a-new-group-label
type CreateGroupLabelOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	Color       *string `url:"color,omitempty" json:"color,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	Priority    *int    `url:"priority,omitempty" json:"priority,omitempty"`
}

// CreateGroupLabel creates a new label for given group with given name and
// color.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#create-a-new-group-label
func (s *GroupLabelsService) CreateGroupLabel(gid interface{}, opt *CreateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/labels", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	l := new(GroupLabel)
	resp, err := s.client.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// DeleteGroupLabelOptions represents the available DeleteGroupLabel() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#delete-a-group-label
type DeleteGroupLabelOptions struct {
	Name *string `url:"name,omitempty" json:"name,omitempty"`
}

// DeleteGroupLabel deletes a group label given by its name or ID.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#delete-a-group-label
func (s *GroupLabelsService) DeleteGroupLabel(gid interface{}, lid interface{}, opt *DeleteGroupLabelOptions, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/labels", PathEscape(group))

	if lid != nil {
		label, err := parseID(lid)
		if err != nil {
			return nil, err
		}
		u = fmt.Sprintf("groups/%s/labels/%s", PathEscape(group), PathEscape(label))
	}

	req, err := s.client.NewRequest(http.MethodDelete, u, opt, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UpdateGroupLabelOptions represents the available UpdateGroupLabel() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#update-a-group-label
type UpdateGroupLabelOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	NewName     *string `url:"new_name,omitempty" json:"new_name,omitempty"`
	Color       *string `url:"color,omitempty" json:"color,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
	Priority    *int    `url:"priority,omitempty" json:"priority,omitempty"`
}

// UpdateGroupLabel updates an existing label with new name or now color. At least
// one parameter is required, to update the label.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#update-a-group-label
func (s *GroupLabelsService) UpdateGroupLabel(gid interface{}, lid interface{}, opt *UpdateGroupLabelOptions, options ...RequestOptionFunc) (*GroupLabel, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/labels", PathEscape(group))

	if lid != nil {
		label, err := parseID(lid)
		if err != nil {
			return nil, nil, err
		}
		u = fmt.Sprintf("groups/%s/labels/%s", PathEscape(group), PathEscape(label))
	}

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	l := new(GroupLabel)
	resp, err := s.client.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// SubscribeToGroupLabel subscribes the authenticated user to a label to receive
// notifications. If the user is already subscribed to the label, the status
// code 304 is returned.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#subscribe-to-a-group-label
func (s *GroupLabelsService) SubscribeToGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*GroupLabel, *Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, nil, err
	}
	label, err := parseID(lid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/labels/%s/subscribe", PathEscape(group), PathEscape(label))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	l := new(GroupLabel)
	resp, err := s.client.Do(req, l)
	if err != nil {
		return nil, resp, err
	}

	return l, resp, nil
}

// UnsubscribeFromGroupLabel unsubscribes the authenticated user from a label to not
// receive notifications from it. If the user is not subscribed to the label, the
// status code 304 is returned.
//
// GitLab API docs:
// https://docs.gitlab.com/api/group_labels/#unsubscribe-from-a-group-label
func (s *GroupLabelsService) UnsubscribeFromGroupLabel(gid interface{}, lid interface{}, options ...RequestOptionFunc) (*Response, error) {
	group, err := parseID(gid)
	if err != nil {
		return nil, err
	}
	label, err := parseID(lid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("groups/%s/labels/%s/unsubscribe", PathEscape(group), PathEscape(label))

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
