package viber

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// PostData to Viber API
func (v *Viber) PostData(url string, i interface{}) ([]byte, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	if v.Debug {
		Log.Println("Post data:", string(b))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Viber-Auth-Token", v.AppKey)

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
