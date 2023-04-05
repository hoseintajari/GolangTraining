package IntSet

import (
	"bytes"
	"strconv"
)

type Intset struct {
	data map[int]bool
}

func (i *Intset) Add(value ...int) {
	for _, v := range value {
		i.data[v] = true
	}
}

func NewInset(value ...int) Intset {
	newdata := Intset{data: make(map[int]bool)}
	newdata.Add(value...)
	return newdata
}

func (i *Intset) Contains(value int) bool {
	_, ok := i.data[value]
	return ok
}

func (i *Intset) Remove(value ...int) {
	for _, v := range value {
		if i.Contains(v) {
			delete(i.data, v)
		}
	}

}

func Union(val ...Intset) Intset {
	ret := NewInset()
	for _, v := range val {
		for index, key := range v.data {
			ret.data[index] = key
		}
	}
	return ret
}

func Intersect(sets ...Intset) Intset {
	newset := Intset{sets[0].data}
	for i := 1; i < len(sets); i++ {
		for index, _ := range newset.data {
			if !sets[i].Contains(index) {
				newset.Remove(index)
			}
		}
	}
	return newset
}

func (i *Intset) Len() int {

	return len(i.data)
}

func (i *Intset) String() string {
	count := 1
	var buf bytes.Buffer
	buf.WriteString("{")

	for v := range i.data {
		if count != len(i.data) {
			buf.WriteString(strconv.Itoa(v))
			buf.WriteString(", ")
			count++
		} else if count == i.Len() {
			buf.WriteString(strconv.Itoa(v))
		}
	}
	buf.WriteString("}")

	return buf.String()
}
