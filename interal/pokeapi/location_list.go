package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations ...
func (c *Client) ListLocations(pageUrl *string) (RespShallowLocation, error) {
	url := baseUrl + "/lcoation-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationsResp := RespShallowLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return locationsResp, nil
}
