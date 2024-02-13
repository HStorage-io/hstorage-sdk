package hstorage_sdk

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hstorage-io/hstorage_common"
)

type UploadFinishRequest struct {
	FileName string `json:"file_name"`
}

func (c *Client) GetUploadConfig(request hstorage_common.PreSignedReq) (*hstorage_common.PreSignedResp, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/upload/config", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var configResponse hstorage_common.PreSignedResp
	if err := json.Unmarshal(body, &configResponse); err != nil {
		return nil, err
	}

	return &configResponse, nil
}

func (c *Client) UploadFinish(fileName string) (*hstorage_common.Upload, error) {
	requestBody, err := json.Marshal(UploadFinishRequest{FileName: fileName})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", c.BaseURL+"/file?update_type=uploaded", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var upload hstorage_common.Upload
	err = json.Unmarshal(body, &upload)

	return &upload, err
}
