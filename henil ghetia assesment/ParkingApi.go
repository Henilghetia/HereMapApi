package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	http.DefaultServeMux.HandleFunc("/x.json", jsonHandler)
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(l, nil)

	baseURL := "https://places.demo.api.here.com/places/v1/discover/explore?at=41.7219688%2C-87.6383633&cat=parking-facility&size=3&Accept-Language=en-GB%2Cen-US%3Bq%3D0.9%2Cen%3Bq%3D0.8&app_id=devportal-demo-20180625&app_code=9v2BkviRwi9Ot26kp2IysQ#" + l.Addr().String()
	type result struct {
		Foo int
	}

	tests := []struct {
		url    string
		result interface{}
	}{{
		url:    baseURL + "/",
		result: new(result),
	}, {
		url:    baseURL + "/x.json",
		result: nil,
	}, {
		url:    baseURL + "/x.json",
		result: new(result),
	}}
	for i, test := range tests {
		err := getJSON(test.url, test.result)
		if err != nil {
			fmt.Printf("test %d: error %v\n", i, err)
		} else {
			fmt.Printf("test %d: ok with result %#v\n", i, test.result)
		}
	}
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"Foo": 1234}`))
}

// getJSON fetches the contents of the given URL
// and decodes it as JSON into the given result,
// which should be a pointer to the expected data.
func getJSON(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot fetch URL %q: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http GET status: %s", resp.Status)
	}
	// We could also check the resulting content type
	// here too.
	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}
	fmt.Println(generic)

	
	
	return nil

	
}