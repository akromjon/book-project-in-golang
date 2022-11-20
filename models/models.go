package models

import "net/url"

type Model interface {
	All() interface{}
	Find(id string) interface{}
	Create(c url.Values) bool
	Update(id string, c url.Values) bool
	Delete(id string) bool
}

func All(m Model) interface{} {
	return m.All()
}

func Find(m Model, id string) interface{} {
	return m.Find(id)
}

func Create(m Model, c url.Values) bool {
	return m.Create(c)
}
func Update(m Model, id string, c url.Values) bool {
	return m.Update(id, c)
}

func Delete(m Model, id string) bool {
	return m.Delete(id)
}
