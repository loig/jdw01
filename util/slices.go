// Package util implements utility functions
/*
Copyright (C) 2020  Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package util

// Overlap checks if two sorted arrays of consecutive
// int overlap or are consecutive
func Overlap(s1, s2 []int) bool {
	return (s1[0] <= s2[0] && s1[len(s1)-1] >= s2[0]) ||
		(s2[0] <= s1[0] && s2[len(s2)-1] >= s1[0]) ||
		(s1[len(s1)-1]+1 == s2[0]) ||
		(s2[len(s2)-1]+1 == s1[0])
}

// Merge merges two sorted arrays of consecutive int that
// are known to overlap
func Merge(s1, s2 []int) []int {
	res := make([]int, 0)
	for len(s1) > 0 && len(s2) > 0 {
		if s1[0] < s2[0] {
			res = append(res, s1[0])
			s1 = s1[1:]
		} else if s1[0] > s2[0] {
			res = append(res, s2[0])
			s2 = s2[1:]
		} else {
			res = append(res, s1[0])
			s1 = s1[1:]
			s2 = s2[1:]
		}
	}
	res = append(res, s1...)
	res = append(res, s2...)
	return res
}

// RegisterSort sorts an array of non overlaping sorted
// arrays of consecutive int, in the order of the int
func RegisterSort(register [][]int) [][]int {
	sorted := make([][]int, 0)
	for len(register) > 0 {
		indiceSmallest := 0
		for i, s := range register {
			if register[indiceSmallest][0] > s[0] {
				indiceSmallest = i
			}
		}
		sorted = append(sorted, register[indiceSmallest])
		register = append(register[:indiceSmallest], register[indiceSmallest+1:]...)
	}
	return sorted
}
