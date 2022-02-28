package valorant

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type respBody struct {
	LastUpdated string `json:"last_updated"`
	Weaponskins struct {
		W1 `json:"1"`
		W2 `json:"2"`
		W3 `json:"3"`
		W4 `json:"4"`
	} `json:"weaponskins"`
	Expires      string `json:"expires"`
	Status       string `json:"status"`
	RitoUsername string `json:"rito_username"`
	RitoId       string `json:"rito_id"`
}

type W1 struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Video string `json:"video"`
}
type W2 struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Video string `json:"video"`
}

type W3 struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Video string `json:"video"`
}

type W4 struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Video string `json:"video"`
}

func GetWeapons() {
	c := &http.Client{}
	u, _ := url.Parse("https://api.checkvalorant.com/store/store/octane0411")
	req := &http.Request{
		Method: "GET",
		URL:    u,
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	b := &respBody{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	log.Printf("%v", b.Weaponskins.W1.Name)
	weapon := b.Weaponskins

	t, err := template.New("weapon").Parse("hi there,\nyour weapons today:\n" +
		"{{.W1.Name}}\n{{.W1.Image}}\n{{.W2.Name}}\n{{.W2.Image}}\n" +
		"{{.W3.Name}}\n{{.W3.Image}}\n{{.W4.Name}}\n{{.W4.Image}}\n")
	err = t.Execute(os.Stdout, weapon)
}
