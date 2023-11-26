package jServCore

import (
	"encoding/json"
)

type Query struct {
	Collection         []string               `json:"collection"`
	Equals             map[string]interface{} `json:"equals"`
	NotEquals          map[string]interface{} `json:"not-equals"`
	LessThan           map[string]interface{} `json:"lt"`
	LessThanEqualTo    map[string]interface{} `json:"lte"`
	GreaterThanEqualTo map[string]interface{} `json:"gte"`
	GreaterThan        map[string]interface{} `json:"gt"`
	Between            map[string]interface{} `json:"between"`
}

func (q Query) ToJson() string {
	data := make(map[string]interface{})
	if len(q.Collection) > 0 {
		if len(q.Collection) == 1 {
			data["collection"] = q.Collection[0]
		} else {
			data["collection"] = q.Collection
		}
	}
	if len(q.Equals) > 0 {
		data["equals"] = q.Equals
	}
	if len(q.NotEquals) > 0 {
		data["not-equals"] = q.NotEquals
	}
	if len(q.LessThan) > 0 {
		data["lt"] = q.LessThan
	}
	if len(q.LessThanEqualTo) > 0 {
		data["lte"] = q.LessThanEqualTo
	}
	if len(q.GreaterThanEqualTo) > 0 {
		data["gte"] = q.GreaterThanEqualTo
	}
	if len(q.GreaterThan) > 0 {
		data["gt"] = q.GreaterThan
	}
	if len(q.Between) > 0 {
		data["between"] = q.Between
	}
	js, _ := json.Marshal(data)
	return string(js)
}

func (q *Query) FromJson(s string) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		panic(err)
	}
	if col, ok := data["collection"]; ok {
		if _, ok := col.(string); ok {
			q.Collection = []string{col.(string)}
		} else {
			q.Collection = col.([]string)
		}
	}
	if eq, ok := data["equals"]; ok {
		q.Equals = eq.(map[string]interface{})
	}
	if neq, ok := data["not-equals"]; ok {
		q.NotEquals = neq.(map[string]interface{})
	}
	if lt, ok := data["lt"]; ok {
		q.LessThan = lt.(map[string]interface{})
	}
	if lte, ok := data["lte"]; ok {
		q.LessThanEqualTo = lte.(map[string]interface{})
	}
	if gte, ok := data["gte"]; ok {
		q.GreaterThanEqualTo = gte.(map[string]interface{})
	}
	if gt, ok := data["gt"]; ok {
		q.GreaterThan = gt.(map[string]interface{})
	}
	if bet, ok := data["between"]; ok {
		q.Between = bet.(map[string]interface{})
	}
}
