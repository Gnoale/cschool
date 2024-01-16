package main

import "sort"

type By func(s1, s2 *seatCoor) bool

// seatSorter joins a By function and a slice of seatCoor to be sorted.
type seatSorter struct {
	seats []seatCoor
	by    func(s1, s2 *seatCoor) bool
}

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(seats []seatCoor) {
	ss := &seatSorter{
		seats,
		by,
	}
	sort.Sort(ss)

}

// Len is part of sort.Interface.
func (s *seatSorter) Len() int {
	return len(s.seats)
}

// Swap is part of sort.Interface.
func (s *seatSorter) Swap(i, j int) {
	s.seats[i], s.seats[j] = s.seats[j], s.seats[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *seatSorter) Less(i, j int) bool {
	return s.by(&s.seats[i], &s.seats[j])
}
