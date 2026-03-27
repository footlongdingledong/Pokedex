package pokeapi

import(
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func Get(url string) ([]byte, error){
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Print("Invalid status code")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) GetLocations(url string) (LocationArea, error){
	if val, ok := c.Cache.Get(url); ok {
		location := LocationArea{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}else {
			body, _ := Get(url)
			c.Cache.Add(url, body)
			location := LocationArea{}
			err := json.Unmarshal(body, &location)
			if err != nil {
				return LocationArea{}, err
			}
			return location, nil
	}
}

func (c *Client) GetLocationInfo(url string) (LocationAreaFull, error) {
	if val, ok := c.Cache.Get(url); ok {
		location := LocationAreaFull{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationAreaFull{}, err
		}
		return location, nil
	}else {
			body, _ := Get(url)
			c.Cache.Add(url, body)
			location := LocationAreaFull{}
			err := json.Unmarshal(body, &location)
			if err != nil {
				return LocationAreaFull{}, err
			}
			return location, nil
	}
}

func (c *Client) GetPokemon(pokemon string) (Pokemon, error){
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	val, ok := c.Cache.Get(url)
	if ok {
		pokeinfo := Pokemon{}
		err := json.Unmarshal(val, &pokeinfo)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeinfo, nil
	}
	body, err := Get(url)
	if err != nil {
		fmt.Println(err)
	}
	c.Cache.Add(url, body)
	pokeinfo := Pokemon{}
	err = json.Unmarshal(body, &pokeinfo)
	if err != nil {
		return Pokemon{}, err
	}
	return pokeinfo, nil
}
