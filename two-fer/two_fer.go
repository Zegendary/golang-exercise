package twofer

func ShareWith(name string) string {
	var value string
	if name != "" {
		value = name
	} else {
		value = "you"
	}
	return "One for " + value + ", one for me."
}
