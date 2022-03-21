package gohubspot

import "fmt"

type TicketsService service

type Ticket struct {
	ObjectType string     `json:"objectType"`
	PortalID   int        `json:"portalId"`
	ObjectID   int        `json:"objectId"`
	IsDeleted  bool       `json:"isDeleted"`
	Properties Properties `json:"properties"`
}

func (s *TicketsService) Create(body interface{}) (*Ticket, error) {
	url := "/crm-objects/v1/objects/tickets"
	res := new(Ticket)
	err := s.client.RunPost(url, body, res)
	return res, err
}

func (s *TicketsService) Update(id int, body interface{}) (*Ticket, error) {
	url := fmt.Sprintf("/crm-objects/v1/objects/tickets/%d", id)
	res := new(Ticket)
	err := s.client.RunPut(url, body, res)
	return res, err
}
