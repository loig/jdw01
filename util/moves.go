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

func SmoothStop2(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

func SmoothStop3(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)
}

func SmoothStop4(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)
}

func SmoothStop5(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)*(1-t)
}

func SmoothStop6(t float64) float64 {
	return 1 - (1-t)*(1-t)*(1-t)*(1-t)*(1-t)*(1-t)
}

func SmoothStart2(t float64) float64 {
	return t * t
}

func SmoothStart3(t float64) float64 {
	return t * t * t
}

func SmoothStart4(t float64) float64 {
	return t * t * t * t
}

func SmoothStart5(t float64) float64 {
	return t * t * t * t * t
}

func SmoothStart6(t float64) float64 {
	return t * t * t * t * t * t
}
