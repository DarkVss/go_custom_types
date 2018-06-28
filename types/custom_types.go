package custom_types

type Array []interface{}

func (a Array)Contains(search interface{}) int {
	for index, item := range a {
		if item == search {
			return index
		}
	}
	return -1
}

func (a *Array) Add(s interface{}) {
	(*a) = append(*a, s)
}

func (a *Array) DeleteKey(index int) {
	(*a) = append((*a)[:index], (*a)[index + 1:]...)
}


//Delete - is proxy function for DeleteKey
func (a *Array) Delete(index int) {
	a.DeleteKey(index)
}
