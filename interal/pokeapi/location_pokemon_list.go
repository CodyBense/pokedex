package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemonLocation ...
func (c *Client) ListPokemonLocation(locationArea string) (RespDeepLocation, error) {
	url := baseUrl + "/location-area/" + locationArea

	data, found := c.cache.Get(url)
	if found {
		locationResp := RespDeepLocation{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}

	defer resp.Body.Close()

	// if resp.StatusCode < 299 || resp.StatusCode > 300 {
	// 	return RespDeepLocation{}, fmt.Errorf("unexpected status: %s", resp.Status)
	// }

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	locationsResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	c.cache.Add(url, dat)

	// fmt.Println("From network request")
	return locationsResp, nil

}
