package langDetect

// represents language
type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// retrieves list of supported languages
func (c *Client) Languages() (out []*Language, err error) {
	err = c.get(nil, "languages", &out)
	return
}