package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"

	fullUrl := baseUrl + endpoint

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	locationAreasResp := LocationAreasResp{}

	//check the cache
	data, ok := c.cache.Get(fullUrl)

	if ok {
		//cache hit
		fmt.Println("cache hit")
		err := json.Unmarshal(data, &locationAreasResp)

		if err != nil {
			return locationAreasResp, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)

	if err != nil {
		return locationAreasResp, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return locationAreasResp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreasResp, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return locationAreasResp, err
	}
 
	c.cache.Add(fullUrl, data)

	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return locationAreasResp, err
	}

	return locationAreasResp, nil
}
