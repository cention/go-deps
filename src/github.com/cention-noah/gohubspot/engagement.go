package gohubspot

type EngagementsService service

type Engagements struct {
	Engagement   *Engagement   `json:"engagement"`
	Associations *Association  `json:"associations"`
	Attachments  []*Attachment `json:"attachments"`
	MetaData     *MetaData     `json:"metadata"`
}
type Engagement struct {
	ID          int      `json:"id"`
	PortalID    int      `json:"portalId"`
	Active      bool     `json:"active"`
	CreatedAt   UnixTime `json:"createdAt"`
	LastUpdated UnixTime `json:"lastUpdated"`
	OwnerId     int      `json:"ownerId"`
	Type        string   `json:"type"`
	TimeStamp   UnixTime `json:"timestamp"`
}
type Association struct {
	ContactIds  []int `json:"contactIds"`
	CompanyIds  []int `json:"companyIds"`
	DealIds     []int `json:"dealIds"`
	OwnerIds    []int `json:"ownerIds"`
	WorkflowIds []int `json:"workflowIds"`
	TicketIds   []int `json:"ticketIds"`
}
type Attachment struct {
	ID int `json:"id"`
}
type MetaData struct {
	Body string `json:"body"`
}

func (s *EngagementsService) Create(body interface{}) (*Engagements, error) {
	url := "/engagements/v1/engagements"
	res := new(Engagements)
	err := s.client.RunPost(url, body, res)
	return res, err
}
