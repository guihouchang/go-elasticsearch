package field

import "fmt"

type Property map[string]interface{}

type Mapping struct {
	Name       string                 `json:"-"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Settings   map[string]interface{} `json:"settings,omitempty"`
}

func (m *Mapping) AlterBody() map[string]interface{} {
	return map[string]interface{}{
		"properties": m.Properties,
	}
}

func (m *Mapping) RecoveryBody() map[string]interface{} {
	return map[string]interface{}{
		"dest": map[string]string{
			"index": m.Name,
		},
		"source": map[string]string{
			"index": m.BackupName(),
		},
	}
}

func (m *Mapping) BackupName() string {
	return fmt.Sprintf("old-%s", m.Name)
}

func (m *Mapping) BackupBody() map[string]interface{} {
	return map[string]interface{}{
		"source": map[string]string{
			"index": m.Name,
		},
		"dest": map[string]string{
			"index": m.BackupName(),
		},
	}
}

func (m *Mapping) CreateBody() map[string]interface{} {
	data := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": m.Properties,
		},
		"settings": m.Settings,
	}

	return data
}
