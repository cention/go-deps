package gohubspot

import "fmt"

type ContactsService service

type Contact struct {
	Vid          int        `json:"vid"`
	CanonicalVid int        `json:"canonical-vid"`
	MergedVids   []int      `json:"merged-vids"`
	PortalID     int        `json:"portal-id"`
	IsContact    bool       `json:"is-contact"`
	ProfileToken string     `json:"profile-token"`
	ProfileURL   string     `json:"profile-url"`
	Properties   Properties `json:"properties"`
}

type Contacts struct {
	Vid          int    `json:"vid"`
	CanonicalVid int    `json:"canonical-vid"`
	MergedVids   []int  `json:"merged-vids"`
	PortalID     int    `json:"portal-id"`
	IsContact    bool   `json:"is-contact"`
	ProfileToken string `json:"profile-token"`
	ProfileURL   string `json:"profile-url"`
	Property     Fields `json:"properties"`
}

type Fields struct {
	FirstName map[string]interface{} `json:"firstname"`
	LastName  map[string]interface{} `json:"lastname"`
	Website   map[string]interface{} `json:"website"`
	Company   map[string]interface{} `json:"company"`
	Phone     map[string]interface{} `json:"phone"`
	Address   map[string]interface{} `json:"address"`
	City      map[string]interface{} `json:"city"`
	State     map[string]interface{} `json:"state"`
	Zip       map[string]interface{} `json:"zip"`
}

func (s *ContactsService) Create(properties Properties) (*IdentityProfile, error) {
	url := "/contacts/v1/contact"
	res := new(IdentityProfile)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) Update(contactID int, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", contactID)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) UpdateByEmail(email string, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s/profile", email)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) CreateOrUpdateByEmail(email string, properties Properties) (*Vid, error) {
	url := fmt.Sprintf("/contacts/v1/contact/createOrUpdate/email/%s", email)

	res := new(Vid)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) DeleteById(id int) (*ContactDeleteResult, error) {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d", id)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(url, res)
	return res, err
}

func (s *ContactsService) DeleteByEmail(email string) (*ContactDeleteResult, error) {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s", email)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(url, res)
	return res, err
}

func (s *ContactsService) Merge(primaryID, secondaryID int) error {
	url := fmt.Sprintf("/contacts/v1/contact/merge-vids/%d/", primaryID)
	secondary := struct {
		SecondaryID int `json:"vidToMerge"`
	}{
		SecondaryID: secondaryID,
	}

	return s.client.RunPost(url, secondary, nil)
}

func (s *ContactsService) GetByToken(token string) (*Contact, error) {
	url := fmt.Sprintf("/contacts/v1/contact/utk/%s/profile", token)
	res := new(Contact)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *ContactsService) GetByEmail(email string) (*Contacts, error) {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s/profile", email)
	res := new(Contacts)
	err := s.client.RunGet(url, res)
	return res, err
}
