package handler

import "encoding/json"

func UnmarshalPerson(data []byte) (Person, error) {
	var r Person
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Person) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Person struct {
	Name    string  `json:"name,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	Area    string `json:"area,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}
