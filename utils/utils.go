package utils

import (
	"encoding/json"
	"net/http"
)

func PerformJSONRequest(client *http.Client, r *http.Request, target interface{}) error {
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	/*
		b, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(b), res.StatusCode)
	*/

	if target != nil {
		return json.NewDecoder(res.Body).Decode(target)
	}

	return nil
}
