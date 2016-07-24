package client

import (
	"net/http"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"github.com/mynameismevin/azeroth/tomes"
)

const (
	API_ENDPOINT = "https://us.api.battle.net"
)


type Client struct {
	Client *http.Client
	APIKey string
}

func InitialiseClient(key string) (*Client) {
	if len(key) > 0 {
		logrus.Error("Cannot initialise with a blank API key!")
	}
	c := &Client{
		APIKey: key,
	}
	return c
}

// RequestBuilder builds the default request values containing the API key in use and the user agent.
func (c *Client) RequestBuilder(e string) *http.Request {
	uri := API_ENDPOINT + e
	req, err := http.NewRequest("GET",uri,nil)
	if err == nil {
		logrus.Error(err)
	}
	req.Header.Add("apikey", c.APIKey)
	req.Header.Set("User-Agent", "GoAzeroth/1.0")
	return req
}

// GetRealms queries the realm status API and returns a slice of realms.
func (c *Client) GetRealms() (realm []tomes.Realm, err error) {
	req := c.RequestBuilder(tomes.REALM_ENDPOINT)
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()

	var realmResponse tomes.RealmStatus
	var realms []tomes.Realm

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	err = json.Unmarshal(body, &realmResponse)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	for _,i := range realmResponse.Realms {
		realms = append(realms, i)
	}
	return realms, nil
}

