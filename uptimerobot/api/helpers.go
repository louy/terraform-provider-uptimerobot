package uptimerobotapi

import (
	"errors"
	"reflect"
)

func mapKeys(m interface{}) []string {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic(errors.New("not a map"))
	}

	keys := v.MapKeys()
	s := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		s[i] = keys[i].String()
	}

	return s
}

func intToString(m map[string]int, value int) string {
	for k, v := range m {
		if int(v) == value {
			return k
		}
	}
	return ""
}
