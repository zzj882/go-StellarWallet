package publicdefine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("[ HTTPGet() ]\r\n\t", "Get: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ HTTPGet() ]\r\n\t", "ReadAll: ", err)
		return nil, err
	}

	// fmt.Println("[ HTTPGet().body ]\r\n\t", string(body))

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("[ HTTPGet() ]\r\n\t", "Unmarshal: ", err)
		return nil, err
	}

	return result, nil
}

func HttpPost_json(address string, data []byte) (map[string]interface{}, error) {

	resp, err := http.Post(address,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		fmt.Println("[ HTTPPOST.json() ]\r\n\t", "Post: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ HTTPPOST.json() ]\r\n\t", "ReadAll: ", err)
		return nil, err
	}

	// logdebug.DebugPrintLn("StellarWebsocketFuncs:callUrl.Body", string(body))

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("[ HTTPPOST.json() ]\r\n\t", "Unmarshal: ", err)
		return nil, err
	}

	// fmt.Println("[ HTTPPOST() ]\r\n\t", string(body))

	return result, nil
}

func HttpPost_form(address, data string) (map[string]interface{}, error) {

	resp, err := http.Post(address,
		"application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println("[ HTTPPOST.form() ]\r\n\t", "Post: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ HTTPPOST.form() ]\r\n\t", "ReadAll: ", err)
		return nil, err
	}

	// logdebug.DebugPrintLn("StellarWebsocketFuncs:callUrl.Body", string(body))

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("[ HTTPPOST.form()) ]\r\n\t", "Unmarshal: ", err)
		return nil, err
	}

	// fmt.Println("[ HTTPPOST.form() ]\r\n\t", string(body))

	return result, nil
}
