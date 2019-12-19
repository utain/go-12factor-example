package models

import (
	"database/sql/driver"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

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
	log.Debug("Value::", attr)
	attrVl, err := json.Marshal(attr)
	return string(attrVl), err
}

//Scan convert database value (sql datatype) to golang datatype (struct).
func (attr *AttrType) Scan(src interface{}) error {
	log.Debug("Scane:", src)
	err := json.Unmarshal([]byte(src.(string)), &attr)
	return err
}

func (p Product) String() string {
	str, _ := json.Marshal(p)
	return string(str)
}
