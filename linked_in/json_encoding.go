package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(manager)
	return &buf, err
}

func main() {
	manager := Manager{
		FullName:       "Firman Agam",
		Position:       "CEO",
		Age:            25,
		YearsInCompany: 3,
	}
	fmt.Println(EncodeManager(&manager))
}
