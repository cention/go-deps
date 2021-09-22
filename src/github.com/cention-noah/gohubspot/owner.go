package gohubspot

import (
	"fmt"
)

type OwnersService service

type Owner struct {
	PortalID                int          `json:"portalId"`
	OwnerId                 int          `json:"ownerId"`
	ActiveUserId            int          `json:"activeUserId"`
	UserIdIncludingInactive int          `json:"userIdIncludingInactive"`
	IsActive                bool         `json:"isActive"`
	HasContactsAccess       bool         `json:"hasContactsAccess"`
	Type                    string       `json:"type"`
	FirstName               string       `json:"firstName"`
	LastName                string       `json:"lastName"`
	Email                   string       `json:"email"`
	CreatedAt               UnixTime     `json:"createdAt"`
	UpdatedAt               UnixTime     `json:"updatedAt"`
	RemoteList              []RemoteList `json:"remoteList"`
}

type RemoteList struct {
	ID         int    `json:"id"`
	PortalID   int    `json:"portalId"`
	OwnerId    int    `json:"ownerId"`
	RemoteId   string `json:"remoteId"`
	RemoteType string `json:"remoteType"`
	Active     bool   `json:"active"`
}

func (s *OwnersService) GetOwnerById(id int) (*Owner, error) {
	url := fmt.Sprintf("/owners/v2/owners/%d", id)

	res := new(Owner)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *OwnersService) GetOwnersByEmail(email string) ([]*Owner, error) {
	url := fmt.Sprintf("/owners/v2/owners?email=%s", email)
	var res []*Owner
	err := s.client.RunGet(url, &res)
	return res, err
}
