package viber

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// PostData to viber API
func (v *Viber) PostData(url string, i interface{}) ([]byte, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	Log.Println("Post data:", string(b))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add("X-Viber-Auth-Token", v.AppKey)
	req.Close = true

	if v.Client == nil {
		v.Client = &http.Client{}
	}

	resp, err := v.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
