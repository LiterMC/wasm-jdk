package properties

var properties = map[string]string{
	"java.home": "/java",
}

func SetProp(key string, value string) {
	properties[key] = value
}

func GetProp(key string) string {
	return properties[key]
}

func RemoveProp(key string) {
	delete(properties, key)
}

func ForEachProp(cb func(k string, v string) bool) {
	for k, v := range properties {
		if !cb(k, v) {
			return
		}
	}
}

func GetPropKVArray() []string {
	arr := make([]string, 0, len(properties) * 2)
	for k, v := range properties {
		arr = append(arr, k, v)
	}
	return arr
}
