package loki

import "fmt"

type Label map[string]string

func (l Label) AddLabel(key string, value string) error {
	if _, ok := l[key]; ok {
		return fmt.Errorf("Key is already exist")
	} else {
		l[key] = value
		return nil
	}
}
