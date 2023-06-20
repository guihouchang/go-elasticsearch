package annotation

type Setting struct {
	Settings map[string]interface{}
}

func (s Setting) Name() string {
	return "Settings"
}
