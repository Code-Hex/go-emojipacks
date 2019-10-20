package emojipacks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const emojiAddEndpoint = "https://slack.com/api/emoji.add"

const useragent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Mobile Safari/537.36"

var _ error = (*Response)(nil)

type Response struct {
	OK       bool   `json:"ok"`
	Err      string `json:"error,omitempty"`
	Needed   string `json:"needed,omitempty"`
	Provided string `json:"provided,omitempty"`
}

func (r *Response) Error() string {
	return fmt.Sprintf("slack error: %s, needed: %s, provided: %s", r.Err, r.Needed, r.Provided)
}

// imgSize == resp.ContentLength
func (u *uploader) uploadEmoji(imgBody io.Reader, imgSize int64, emojiName string) error {

	data := map[string]string{
		"mode":  "data",
		"name":  emojiName,
		"token": u.accessToken,
	}

	form, err := multiPartFormStream(data, imgBody, imgSize)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", emojiAddEndpoint, form)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", form.contentType)
	req.Header.Set("User-Agent", useragent)
	req.ContentLength = form.contentLength

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("rejected (code=%d)", resp.StatusCode)
	}

	var ir Response
	if err := json.NewDecoder(resp.Body).Decode(&ir); err != nil {
		return err
	}

	if !ir.OK {
		return &ir // treats as an error
	}

	return nil
}

var _ io.Reader = (*formData)(nil)

type formData struct {
	io.Reader
	contentType   string
	contentLength int64
}

func multiPartFormStream(data map[string]string, imgBody io.Reader, contentLength int64) (*formData, error) {
	var form bytes.Buffer
	mpWriter := multipart.NewWriter(&form)
	for k, v := range data {
		if err := mpWriter.WriteField(k, v); err != nil {
			return nil, err
		}
	}

	if _, err := mpWriter.CreateFormFile("image", "tmp_img.gif"); err != nil {
		return nil, err
	}
	contentType := mpWriter.FormDataContentType()

	var mpData bytes.Buffer
	if _, err := io.Copy(&mpData, &form); err != nil {
		return nil, err
	}

	mpWriter.Close()

	return &formData{
		Reader:        io.MultiReader(&mpData, imgBody, &form),
		contentType:   contentType,
		contentLength: contentLength + int64(mpData.Len()+form.Len()),
	}, nil
}
