package gohubspot

type CRMAssociationsService service

func (s *CRMAssociationsService) Create(body interface{}) error {
	url := "/crm-associations/v1/associations"
	return s.client.RunPut(url, body, nil)
}
