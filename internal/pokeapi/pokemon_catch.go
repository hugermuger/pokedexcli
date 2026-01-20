package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) CatchPokemon(pageURL string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pageURL

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespPokemon{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	locationsResp := RespPokemon{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
