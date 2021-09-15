package endpoint

type Endpoint struct {
	Protocol   string
	Domain     string
	URI        string
	Method     string
	Header     []string
	Parameters []string
}

func (e *Endpoint) Endpoint() string {

	return ""
}
