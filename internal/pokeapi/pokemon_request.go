package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon *string) (PokemonResp, error) {
	endpoint := "/pokemon/" + *pokemon

	fullUrl := baseUrl + endpoint
	
	pokemonResp := PokemonResp{}

	//check the cache
	data, ok := c.cache.Get(fullUrl)

	if ok {
		//cache hit
		fmt.Println("cache hit")
		err := json.Unmarshal(data, &pokemonResp)

		if err != nil {
			return pokemonResp, err
		}

		return pokemonResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return pokemonResp, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return pokemonResp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return pokemonResp, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return pokemonResp, err
	}
 
	c.cache.Add(fullUrl, data)

	err = json.Unmarshal(data, &pokemonResp)

	if err != nil {
		return pokemonResp, err
	}

	return pokemonResp, nil
}
