package publicdefine

import (
	"encoding/json"
	// "fmt"
	"errors"
	"github.com/jojopoper/ConsoleColor"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type UseProxyInfo struct {
	Enabled  bool
	IP       string
	Port     string
	UserName string
	Password string
}

var CurrProxyInfo = &UseProxyInfo{
	Enabled: false,
}

func HttpGet(geturl string) (map[string]interface{}, error) {
	resp, err := http.Get(geturl)
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, "ERROR [ HTTPGet() ]\r\n\t", "Get: ", err)
		return nil, err
	}
	return getResponseDecode(resp)
}

func HttpProxyGet(pgUrl, proxyIP, proxyPort string, auth ...*proxy.Auth) (map[string]interface{}, error) {

	client := getProxyClient(proxyIP, proxyPort, auth...)

	resp, err := client.Get(pgUrl)
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** HTTPProxyGet() Error\r\n\t", "Get: ", err)
		return nil, err
	}

	return getResponseDecode(resp)
}

func HttpPost_json(address string, data []byte) (map[string]interface{}, error) {

	resp, err := http.Post(address,
		"application/json", strings.NewReader(string(data)))
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, "ERROR [ HTTPPOST.json() ]\r\n\t", "Post: ", err)
		return nil, err
	}

	return getResponseDecode(resp)
}

func HttpPost_form(address, data string) (map[string]interface{}, error) {

	resp, err := http.Post(address,
		"application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** HTTPPost() Error\r\n\t", "Post: ", err)
		return nil, err
	}

	return getResponseDecode(resp)
}

func HttpProxyPost_form(address, data, proxyIP, proxyPort string, auth ...*proxy.Auth) (map[string]interface{}, error) {

	client := getProxyClient(proxyIP, proxyPort, auth...)

	resp, err := client.Post(address,
		"application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** HTTPProxyPost() Error\r\n\t", "Post: ", err)
		return nil, err
	}

	return getResponseDecode(resp)
}

func getResponseDecode(resp *http.Response) (map[string]interface{}, error) {

	if resp == nil {
		return nil, errors.New("http.Response is nil!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** HTTP Response ERROR\r\n\t", "ReadAll: ", err)
		return nil, err
	}

	// fmt.Println("[ HTTPGet().body ]\r\n\t", string(body))

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** HTTP Response ERROR\r\n\t", "Unmarshal: ", err)
		return nil, err
	}

	return result, nil
}

func getProxyClient(proxyIP, proxyPort string, auth ...*proxy.Auth) *http.Client {

	proxyurl := proxyIP + ":" + proxyPort
	var author *proxy.Auth = nil
	if auth != nil {
		author = auth[0]
	}
	dialer, err := proxy.SOCKS5("tcp", proxyurl, author,
		&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	)

	if err != nil {
		ConsoleColor.Println(ConsoleColor.C_RED, " ** Socket5ProxyClient() Error\r\n\t", "proxy.SOCKS5: ", err)
		return nil
	}

	transport := &http.Transport{
		Proxy:               nil,
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	return &http.Client{Transport: transport}
}
