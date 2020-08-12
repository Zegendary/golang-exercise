package strain

type Lister interface {
	Keep()
	Discard()
}
type Strings []string
type Ints []int
type Lists [][]int

func (inputs Ints) Keep(pred func(int) bool) Ints {
	var result = Ints{}
	if inputs == nil {
		return nil
	}
	for _, item := range inputs {
		if pred(item) {
			result = append(result, item)
		}
	}
	return result
}

func (inputs Ints) Discard(pred func(int) bool) Ints {
	if inputs == nil {
		return nil
	}
	var result = Ints{}
	for _, item := range inputs {
		if !pred(item) {
			result = append(result, item)
		}
	}
	return result
}

func (inputs Strings) Keep(pred func(string) bool) Strings {
	if inputs == nil {
		return nil
	}
	var result = Strings{}
	for _, item := range inputs {
		if pred(item) {
			result = append(result, item)
		}
	}
	return result
}

func (inputs Strings) Discard(pred func(string) bool) Strings {
	if inputs == nil {
		return nil
	}
	var result = Strings{}
	for _, item := range inputs {
		if !pred(item) {
			result = append(result, item)
		}
	}
	return result
}

func (inputs Lists) Keep(pred func([]int) bool) Lists {
	if inputs == nil {
		return nil
	}
	var result = Lists{}
	for _, item := range inputs {
		if pred(item) {
			result = append(result, item)
		}
	}
	return result
}

func (inputs Lists) Discard(pred func([]int) bool) Lists {
	if inputs == nil {
		return nil
	}
	var result = Lists{}
	for _, item := range inputs {
		if !pred(item) {
			result = append(result, item)
		}
	}
	return result
}
