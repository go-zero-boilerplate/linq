// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Float64

package linqFloat64

import "errors"

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// Float64Slice is a slice of type Float64. Use it where you would use []Float64.
type Float64Slice []Float64

// All verifies that all elements of Float64Slice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv Float64Slice) All(fn func(Float64) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of Float64Slice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv Float64Slice) Any(fn func(Float64) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// Count gives the number elements of Float64Slice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv Float64Slice) Count(fn func(Float64) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// Distinct returns a new Float64Slice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv Float64Slice) Distinct() (result Float64Slice) {
	appended := make(map[Float64]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// DistinctBy returns a new Float64Slice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv Float64Slice) DistinctBy(equal func(Float64, Float64) bool) (result Float64Slice) {
Outer:
	for _, v := range rcv {
		for _, r := range result {
			if equal(v, r) {
				continue Outer
			}
		}
		result = append(result, v)
	}
	return result
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv Float64Slice) First(fn func(Float64) bool) (result Float64, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no Float64Slice elements return true for passed func")
	return
}

// SortBy returns a new ordered Float64Slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Float64Slice) SortBy(less func(Float64, Float64) bool) Float64Slice {
	result := make(Float64Slice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortFloat64Slice(result, less, 0, n, maxDepth)
	return result
}

// Where returns a new Float64Slice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv Float64Slice) Where(fn func(Float64) bool) (result Float64Slice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapFloat64Slice(rcv Float64Slice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapFloat64Slice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(rcv[first+child], rcv[first+child+1]) {
			child++
		}
		if !less(rcv[first+root], rcv[first+child]) {
			return
		}
		swapFloat64Slice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownFloat64Slice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapFloat64Slice(rcv, first, first+i)
		siftDownFloat64Slice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapFloat64Slice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapFloat64Slice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapFloat64Slice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeFloat64Slice(rcv Float64Slice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapFloat64Slice(rcv, a+i, b+i)
	}
}

func doPivotFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeFloat64Slice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeFloat64Slice(rcv, less, m, m-s, m+s)
		medianOfThreeFloat64Slice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeFloat64Slice(rcv, less, lo, m, hi-1)

	// Invariants are:
	//	rcv[lo] = pivot (set up by ChoosePivot)
	//	rcv[lo <= i < a] = pivot
	//	rcv[a <= i < b] < pivot
	//	rcv[b <= i < c] is unexamined
	//	rcv[c <= i < d] > pivot
	//	rcv[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.
	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi
	for {
		for b < c {
			if less(rcv[b], rcv[pivot]) { // rcv[b] < pivot
				b++
			} else if !less(rcv[pivot], rcv[b]) { // rcv[b] = pivot
				swapFloat64Slice(rcv, a, b)
				a++
				b++
			} else {
				break
			}
		}
		for b < c {
			if less(rcv[pivot], rcv[c-1]) { // rcv[c-1] > pivot
				c--
			} else if !less(rcv[c-1], rcv[pivot]) { // rcv[c-1] = pivot
				swapFloat64Slice(rcv, c-1, d-1)
				c--
				d--
			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// rcv[b] > pivot; rcv[c-1] < pivot
		swapFloat64Slice(rcv, b, c-1)
		b++
		c--
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := min(b-a, a-lo)
	swapRangeFloat64Slice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeFloat64Slice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortFloat64Slice(rcv Float64Slice, less func(Float64, Float64) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortFloat64Slice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotFloat64Slice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortFloat64Slice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortFloat64Slice(rcv, mhi, b)
		} else {
			quickSortFloat64Slice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortFloat64Slice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortFloat64Slice(rcv, less, a, b)
	}
}
