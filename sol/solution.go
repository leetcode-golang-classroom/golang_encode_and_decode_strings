package sol

import "strings"

type ListHandle struct {
	sep string
}

func Constructor(sep string) ListHandle {
	return ListHandle{
		sep: sep,
	}
}
func (h *ListHandle) encode(list []string) string {
	return strings.Join(list, h.sep)
}

func (h *ListHandle) decode(str string) []string {
	return strings.Split(str, h.sep)
}
