package sol

func RunTest(input []string) []string {
	h := ListHandle{sep: ":;"}
	return h.decode(h.encode(input))
}
