package stnet

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type City struct {
	ID       string
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (c City) toJson() string {
	return fmt.Sprintf(`{"name":"%s", "location":"%s"}`, c.Name, c.Location)
}

func Rest() {
	s := createServerRest(addr)
	go s.ListenAndServe()

	cities, err := getCities()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Retrived cities: %v\n", cities)
	city, err := saveCity(City{"", "Paris", "France"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Save city: %v\n", city)
}

func saveCity(city City) (City, error) {

	r, err := http.Post(strings.Join([]string{"http://", addr, "/cities"}, ""), "application/json", strings.NewReader(city.toJson()))
	if err != nil {
		return City{}, nil
	}

	defer r.Body.Close()

	return decodeCity(r.Body)
}

func getCities() ([]City, error) {
	r, err := http.Get(strings.Join([]string{"http://", addr, "/cities"}, ""))
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()
	return decodeCitiess(r.Body)
}

func decodeCity(r io.Reader) (City, error) {
	city := City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&city)
	return city, err
}

func decodeCitiess(r io.Reader) ([]City, error) {
	cities := []City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&cities)
	return cities, err
}

func createServerRest(addr string) http.Server {
	cities := []City{
		City{"1", "Prague", "Czechia"},
		City{"2", "Bratislava", "Slovakia"},
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		if r.Method == http.MethodGet {
			enc.Encode(cities)
		} else if r.Method == http.MethodPost {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			r.Body.Close()

			city := City{}
			json.Unmarshal(data, &city)
			city.ID = strconv.Itoa(len(cities) + 1)
			cities = append(cities, city)
			enc.Encode(city)

		}
	})

	return http.Server{
		Addr:    addr,
		Handler: mux,
	}

}