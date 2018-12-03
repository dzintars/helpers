package discovery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Service struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Services struct {
	Service Service `json:"service"`
}

func GetConfig(service string) *Service {

	url := fmt.Sprintf("http://127.0.0.1:9000/read/%s", service)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}
	defer resp.Body.Close()

	var record Services

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	return &record.Service
}
