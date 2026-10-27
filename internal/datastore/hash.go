package datastore

type Hash struct {
	fields map[string]string
}

func NewHash() *Hash {
	return &Hash{
		fields: make(map[string]string),
	}
}

func (h *Hash) Set(field, value string) {
	h.fields[field] = value
}

func (h *Hash) Get(field string) (string, bool) {

	value, ok := h.fields[field]

	return value, ok
}

func (h *Hash) Delete(field string) bool {

	if _, ok := h.fields[field]; !ok {
		return false
	}

	delete(h.fields, field)

	return true
}

func (h *Hash) Exists(field string) bool {

	_, ok := h.fields[field]

	return ok
}

func (h *Hash) Len() int {

	return len(h.fields)
}

func (h *Hash) GetAll() map[string]string {

	result := make(map[string]string)

	for k, v := range h.fields {
		result[k] = v
	}

	return result
}
