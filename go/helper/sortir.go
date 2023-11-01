package helper

import "sort"

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func SortingMapIntAsc(population map[string]int) PairList {
	p := make(PairList, len(population))
	i := 0
	for k, v := range population {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func SortingMapIntDesc(population map[string]int) PairList {
	p := make(PairList, len(population))
	i := 0
	for k, v := range population {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	sort.Slice(p, func(i, j int) bool {
		return p[i].Value > p[j].Value || p[i].Key < p[j].Key
	})
	return p
}
