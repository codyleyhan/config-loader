package config

import "strings"

type loader struct {
	fileName string
	data     map[string]interface{}
}

// traverse the data map depending on the keys
func (config *loader) traverse(keys []string) (map[string]interface{}, bool) {
	currentMap := config.data
	maxDepth := len(keys) - 1

	for i := 0; i < maxDepth; i++ {
		current, ok := currentMap[keys[i]].(map[string]interface{})
		if !ok {
			return nil, false
		}

		currentMap = current
	}

	return currentMap, true
}

func (config *loader) FileName() string {
	return config.fileName
}

//Get gets the value out of a key from the json data, if not found value is ""
func (config *loader) Get(key string) string {
	splitKey := strings.Split(key, ".")

	if len(splitKey) == 1 {
		genericValue := config.data[key]

		value, ok := genericValue.(string)

		if !ok {
			return ""
		}

		return value
	}

	finalKey := splitKey[len(splitKey)-1]
	finalMap, ok := config.traverse(splitKey)

	value, ok := finalMap[finalKey].(string)
	if !ok {
		return ""
	}

	return value
}

//GetBool gets the value out of a key from the json data, if not found value is false
func (config *loader) GetBool(key string) bool {
	splitKey := strings.Split(key, ".")

	if len(splitKey) == 1 {
		genericValue := config.data[key]

		value, ok := genericValue.(bool)

		if !ok {
			return false
		}

		return value
	}

	finalKey := splitKey[len(splitKey)-1]
	finalMap, ok := config.traverse(splitKey)

	value, ok := finalMap[finalKey].(bool)
	if !ok {
		return false
	}

	return value
}

//GetFloat gets the value out of a key from the json data, if not found value is 0.0
func (config *loader) GetFloat(key string) float64 {
	splitKey := strings.Split(key, ".")

	if len(splitKey) == 1 {
		genericValue := config.data[key]

		value, ok := genericValue.(float64)

		if !ok {
			return 0.0
		}

		return value
	}

	finalKey := splitKey[len(splitKey)-1]
	finalMap, ok := config.traverse(splitKey)

	value, ok := finalMap[finalKey].(float64)
	if !ok {
		return 0.0
	}

	return value
}

//GetInt gets the value out of a key from the json data, if not found value is 0
func (config *loader) GetInt(key string) int64 {
	floatValue := config.GetFloat(key)

	return int64(floatValue)
}

//GetUint gets the value out of a key from the json data, if not found value is 0
func (config *loader) GetUint(key string) uint64 {
	floatValue := config.GetFloat(key)

	return uint64(floatValue)
}
