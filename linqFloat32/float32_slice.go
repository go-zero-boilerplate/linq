// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Float32

package linqFloat32

import "errors"

// Sort implementation is a modification of http://golang.org/pkg/sort/#Sort
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found at http://golang.org/LICENSE.

// Float32Slice is a slice of type Float32. Use it where you would use []Float32.
type Float32Slice []Float32

// All verifies that all elements of Float32Slice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv Float32Slice) All(fn func(Float32) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of Float32Slice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv Float32Slice) Any(fn func(Float32) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// Count gives the number elements of Float32Slice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv Float32Slice) Count(fn func(Float32) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// Distinct returns a new Float32Slice whose elements are unique. See: http://clipperhouse.github.io/gen/#Distinct
func (rcv Float32Slice) Distinct() (result Float32Slice) {
	appended := make(map[Float32]bool)
	for _, v := range rcv {
		if !appended[v] {
			result = append(result, v)
			appended[v] = true
		}
	}
	return result
}

// DistinctBy returns a new Float32Slice whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv Float32Slice) DistinctBy(equal func(Float32, Float32) bool) (result Float32Slice) {
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
func (rcv Float32Slice) First(fn func(Float32) bool) (result Float32, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no Float32Slice elements return true for passed func")
	return
}

// SortBy returns a new ordered Float32Slice, determined by a func defining ‘less’. See: http://clipperhouse.github.io/gen/#SortBy
func (rcv Float32Slice) SortBy(less func(Float32, Float32) bool) Float32Slice {
	result := make(Float32Slice, len(rcv))
	copy(result, rcv)
	// Switch to heapsort if depth of 2*ceil(lg(n+1)) is reached.
	n := len(result)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSortFloat32Slice(result, less, 0, n, maxDepth)
	return result
}

// Where returns a new Float32Slice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv Float32Slice) Where(fn func(Float32) bool) (result Float32Slice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Sort implementation based on http://golang.org/pkg/sort/#Sort, see top of this file

func swapFloat32Slice(rcv Float32Slice, a, b int) {
	rcv[a], rcv[b] = rcv[b], rcv[a]
}

// Insertion sort
func insertionSortFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(rcv[j], rcv[j-1]); j-- {
			swapFloat32Slice(rcv, j, j-1)
		}
	}
}

// siftDown implements the heap property on rcv[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDownFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, lo, hi, first int) {
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
		swapFloat32Slice(rcv, first+root, first+child)
		root = child
	}
}

func heapSortFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDownFloat32Slice(rcv, less, i, hi, first)
	}

	// Pop elements, largest first, into end of rcv.
	for i := hi - 1; i >= 0; i-- {
		swapFloat32Slice(rcv, first, first+i)
		siftDownFloat32Slice(rcv, less, lo, i, first)
	}
}

// Quicksort, following Bentley and McIlroy,
// Engineering a Sort Function, SP&E November 1993.

// medianOfThree moves the median of the three values rcv[a], rcv[b], rcv[c] into rcv[a].
func medianOfThreeFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, a, b, c int) {
	m0 := b
	m1 := a
	m2 := c
	// bubble sort on 3 elements
	if less(rcv[m1], rcv[m0]) {
		swapFloat32Slice(rcv, m1, m0)
	}
	if less(rcv[m2], rcv[m1]) {
		swapFloat32Slice(rcv, m2, m1)
	}
	if less(rcv[m1], rcv[m0]) {
		swapFloat32Slice(rcv, m1, m0)
	}
	// now rcv[m0] <= rcv[m1] <= rcv[m2]
}

func swapRangeFloat32Slice(rcv Float32Slice, a, b, n int) {
	for i := 0; i < n; i++ {
		swapFloat32Slice(rcv, a+i, b+i)
	}
}

func doPivotFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's Ninther, median of three medians of three.
		s := (hi - lo) / 8
		medianOfThreeFloat32Slice(rcv, less, lo, lo+s, lo+2*s)
		medianOfThreeFloat32Slice(rcv, less, m, m-s, m+s)
		medianOfThreeFloat32Slice(rcv, less, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThreeFloat32Slice(rcv, less, lo, m, hi-1)

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
				swapFloat32Slice(rcv, a, b)
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
				swapFloat32Slice(rcv, c-1, d-1)
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
		swapFloat32Slice(rcv, b, c-1)
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
	swapRangeFloat32Slice(rcv, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRangeFloat32Slice(rcv, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

func quickSortFloat32Slice(rcv Float32Slice, less func(Float32, Float32) bool, a, b, maxDepth int) {
	for b-a > 7 {
		if maxDepth == 0 {
			heapSortFloat32Slice(rcv, less, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivotFloat32Slice(rcv, less, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSortFloat32Slice(rcv, less, a, mlo, maxDepth)
			a = mhi // i.e., quickSortFloat32Slice(rcv, mhi, b)
		} else {
			quickSortFloat32Slice(rcv, less, mhi, b, maxDepth)
			b = mlo // i.e., quickSortFloat32Slice(rcv, a, mlo)
		}
	}
	if b-a > 1 {
		insertionSortFloat32Slice(rcv, less, a, b)
	}
}
