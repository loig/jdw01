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

// SmoothStop2 gives a position between 0 and 1 in function
// of a time t beetween 0 and 1, so that the corresponding
// move stops smoothly
func SmoothStop2(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

// SmoothStop3 is even smoother than SmoothStop2
func SmoothStop3(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)
}

// SmoothStop4 is even smoother than SmoothStop3
func SmoothStop4(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)
}

// SmoothStop5 is even smoother than SmoothStop4
func SmoothStop5(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)*(1-t)
}

// SmoothStop6 is even smoother than SmoothStop5
func SmoothStop6(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)*(1-t)*(1-t)
}

// SmoothStart2 gives a position between 0 and 1 in function
// of a time t beetween 0 and 1, so that the corresponding
// move starts smoothly
func SmoothStart2(t float64) float64 {
	return t * t
}

// SmoothStart3 is even smoother than SmoothStart2
func SmoothStart3(t float64) float64 {
	return t * t * t
}

// SmoothStart4 is even smoother than SmoothStart3
func SmoothStart4(t float64) float64 {
	return t * t * t * t
}

// SmoothStart5 is even smoother than SmoothStart4
func SmoothStart5(t float64) float64 {
	return t * t * t * t * t
}

// SmoothStart6 is even smoother than SmoothStart5
func SmoothStart6(t float64) float64 {
	return t * t * t * t * t * t
}
