package idsSource

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Rules struct{
	infects *Infects
}

func (c * Rules) NewRules() *Infects{
	if c.infects!=nil{
		return c.infects
	}
	jsonFile, err := os.Open("filters.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	c.infects=new(Infects)

	json.Unmarshal(byteValue, c.infects)
	return c.infects

}