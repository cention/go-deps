package gohubspot

type CMSFilesService service

type File struct {
	Objects []Objects `json:"objects"`
}
type Objects struct {
	Id int `json:"id"`
}

func (s *CMSFilesService) Upload(body interface{}, contentType string) (*File, error) {
	url := "/filemanager/api/v3/files/upload"
	s.client.ContentType = contentType
	res := new(File)
	err := s.client.RunPost(url, body, res)
	return res, err
}
