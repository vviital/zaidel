package database

import (
	"sort"
	"sync"
)

type waveLengthBasedFastFinder struct {
	table TablePerElement
	once  sync.Once
}

func newWaveLengthBasedFastFinder(table TablePerElement) Zaidel {
	finder := waveLengthBasedFastFinder{}
	finder.init(table)

	return &finder
}

func (f *waveLengthBasedFastFinder) init(table TablePerElement) {
	f.once.Do(func() {
		var wg sync.WaitGroup

		elNum := len(table)
		f.table = make(map[string]Elements, elNum)
		wg.Add(elNum)

		for name := range table {
			elements := make(Elements, len(table[name]))
			f.table[name] = elements
			go func(name string) {
				defer wg.Done()
				copy(elements, table[name])
				sort.Sort(elements)
			}(name)
		}

		wg.Wait()
	})
}

func (f *waveLengthBasedFastFinder) processElement(name string, criteria SearchCriteria, ch chan Element, wg *sync.WaitGroup) {
	defer wg.Done()
	elements := f.table[name]
	l := len(elements)

	index := sort.Search(l, func(i int) bool {
		if elements[i].WaveLength < criteria.MinWaveLength {
			return false
		}
		return true
	})

	for i := index; i < l; i++ {
		if elements[i].WaveLength > criteria.MaxWaveLength {
			break
		}
		ch <- elements[i]
	}
}

func (f *waveLengthBasedFastFinder) FindElements(criteria SearchCriteria) chan Element {
	ch := make(chan Element, 2*len(f.table)) // allow each goroutine to write into channel
	var wg sync.WaitGroup

	has := func(name string) bool {
		if criteria.Elements == nil {
			return true
		}
		_, ok := criteria.Elements[name]
		return ok
	}

	for name := range f.table {
		if !has(name) {
			continue
		}

		wg.Add(1)
		go f.processElement(name, criteria, ch, &wg)
	}

	go func() {
		defer close(ch)
		defer wg.Wait()
	}()

	return ch
}
