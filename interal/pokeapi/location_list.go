package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations ...
func (c *Client) ListLocations(pageUrl *string) (RespShallowLocation, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	data, found := c.cache.Get(url)
	if found {
		locationsResp := RespShallowLocation{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocation{}, err
		}
		// fmt.Println("From cache")
		return locationsResp, nil
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

	if resp.StatusCode < 299 || resp.StatusCode > 300 {
		return RespShallowLocation{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationsResp := RespShallowLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	c.cache.Add(url, dat)

	// fmt.Println("From network request")
	return locationsResp, nil
}
