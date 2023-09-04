package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(location string) (LocationAreaResp, error) {
	endpoint := "/location-area/" + location

	fullUrl := baseUrl + endpoint

	locationAreaResp := LocationAreaResp{}

	//check the cache
	data, ok := c.cache.Get(fullUrl)

	if ok {
		//cache hit
		fmt.Println("cache hit")
		err := json.Unmarshal(data, &locationAreaResp)

		if err != nil {
			return locationAreaResp, err
		}

		return locationAreaResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return locationAreaResp, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return locationAreaResp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreaResp, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return locationAreaResp, err
	}
 
	c.cache.Add(fullUrl, data)

	err = json.Unmarshal(data, &locationAreaResp)

	if err != nil {
		return locationAreaResp, err
	}

	return locationAreaResp, nil
}
