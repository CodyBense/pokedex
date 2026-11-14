package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon ...
func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemon

	data, found := c.cache.Get(url)
	if found {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	// if resp.StatusCode < 299 || resp.StatusCode > 300 {
	// 	return Pokemon{}, fmt.Errorf("unexpected status: %s", resp.Status)
	// }

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	// fmt.Println("From network request")
	return pokemonResp, nil
}
