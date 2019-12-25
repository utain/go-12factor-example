package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Product model
type Product struct {
	Model
	Name  string         `json:"name"`
	Code  string         `json:"code"`
	Price uint           `json:"price"`
	Attr  AttrType       `json:"attr" gorm:"type:text"`
	Props []ProductProps `json:"props" gorm:"foreignkey:ProductRef;"`
}

// AttrType example for custom sql data type using build in interface
type AttrType map[string]string

// Value convert golang datatype (struct)  to json-string (sql datatype) for save to database.
func (attr AttrType) Value() (driver.Value, error) {
	fmt.Println("Value::", attr)
	attrVl, err := json.Marshal(attr)
	return string(attrVl), err
}

//Scan convert database value (sql datatype) to golang datatype (struct).
func (attr *AttrType) Scan(src interface{}) error {
	fmt.Println("Scane:", src)
	err := json.Unmarshal([]byte(src.(string)), &attr)
	return err
}

func (p Product) String() string {
	str, _ := json.Marshal(p)
	return string(str)
}
