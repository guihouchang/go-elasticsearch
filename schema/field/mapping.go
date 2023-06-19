package field

type Property map[string]interface{}

type Mapping struct {
	Name       string
	Properties map[string]Property
	Setting    map[string]Property
}
