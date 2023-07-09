package object

import "encoding/json"

func ToString(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}
