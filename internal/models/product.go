package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type StringArray []string

// Scan implementa a interface sql.Scanner
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
		return nil
	}
	
	// Converte de []byte para string
	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return errors.New("invalid type for StringArray")
	}
	
	// Remove os caracteres '{' e '}' e divide a string
	str = str[1 : len(str)-1]
	if str == "" {
		*a = StringArray{}
		return nil
	}
	
	// Parse a string como JSON array
	var arr []string
	err := json.Unmarshal([]byte("["+str+"]"), &arr)
	if err != nil {
		return err
	}
	
	*a = arr
	return nil
}

// Value implementa a interface driver.Valuer
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "{}", nil
	}
	
	// Converte para formato PostgreSQL array
	bytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	
	return "{" + string(bytes[1:len(bytes)-1]) + "}", nil
}

type Product struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	CreatedAt   string     `json:"created_at"`
	Rating      float64    `json:"rating"`
	Size        string     `json:"size"`
	Reviews     StringArray `json:"reviews"`
	Image       string     `json:"image"`
}
