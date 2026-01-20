package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListPokemon(pageURL string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + pageURL

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespDeepLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespDeepLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}

	locationsResp := RespDeepLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
