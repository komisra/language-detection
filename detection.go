package langDetect

// request parameters for lang detection
type DetectReq struct {
	Query string `json:"q"`
}

// pointer to lang detection response
type DetectResp struct {
	Data *DetectResponseData `json:"data"`
}

// pointer to array of data for lang detection response
type DetectResponseData struct {
	Detections []*DetectionResult `json:"detections"`
}

// result of a single lang detection test
type DetectionResult struct {
	Language  string  `json:"language"`	  /* which language is detected */
	Valid  string  `json:"reliable"` /* Y/N */
	Confidence float32 `json:"confidence`
}

// contains batch language detection request parameters
type DetectBatchRequest struct {
	Query []string `json:"q"`
}

// contains language detection response
type DetectBatchResponse struct {
	Data *DetectBatchResponseData `json:"data"`
}

// contains batch language detection response data
type DetectBatchResponseData struct {
	Detections [][]*DetectionResult `json:"detections"`
}

// for a single text, executes the actual language detectiona
func (c *Client) Detect(in string) (out []*DetectionResult, err error) {
	var response DetectResponse
	err = c.post(nil, "detect", &DetectReq{Query: in}, &response)

	if err != nil {
		return nil, err
	}

	return response.Data.Detections, err
}

// returns detected language code based on language detection of a single text
func (c *Client) DetectCode(in string) (out string, err error) {
	detections, err := c.Detect(in)

	if err != nil {
		return "", err
	}

	if len(detections) == 0 {
		return "", &DetectionError{"Language not detected"}
	}

	return detections[0].Language, err
}

// uses multiple texts to execute language detection
func (c *Client) DetectBatch(in []string) (out [][]*DetectionResult, err error) {
	var response DetectBatchResponse
	err = c.post(nil, "detect", &DetectBatchRequest{Query: in}, &response)

	if err != nil {
		return nil, err
	}

	return response.Data.Detections, err
}