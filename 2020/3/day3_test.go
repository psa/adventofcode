package main

import (
	"reflect"
	"testing"
)

func TestParseTreeLines(t *testing.T) {
	data := []string{
		".#.#.#",
		".....#",
	}
	expected := map[int][]int{
		0: {1, 3, 5},
		1: {5},
	}
	length, result := parseTreeLines(data)

	if !reflect.DeepEqual(result, expected) {
		t.Log("Error, tree map wrong, got", result)
		t.Fail()
	}
	if length != 6 {
		t.Log("Error, expect length 5, got", length)
		t.Fail()
	}
}

func TestHitTreeHit(t *testing.T) {
	position := 2
	trees := []int{2, 4, 6}

	if !hitTree(position, trees) {
		t.Log("Error, expect to hit tree, did not")
		t.Fail()
	}
}

func TestHitTreeMiss(t *testing.T) {
	position := 3
	trees := []int{2, 4, 6}

	if hitTree(position, trees) {
		t.Log("Error, expect to miss tree, did not")
		t.Fail()
	}
}

func TestExtendTrees(t *testing.T) {
	length := 6
	trees := []int{1, 3, 5}
	position := 15
	expected := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	result := extendTrees(length, trees, position)
	if !reflect.DeepEqual(result, expected) {
		t.Log("Error, unexpected extension, got", result)
		t.Fail()
	}
}

func TestScanTrees(t *testing.T) {
	length := 6
	trees := map[int][]int{
		0: {1, 3, 5},
		1: {5},
	}
	right := 5
	down := 1

	hits := scanTrees(length, trees, right, down)
	if hits != 1 {
		t.Log("Error, expect 1 hits, got", hits)
		t.Fail()
	}
}

func TestScanTreesExtension(t *testing.T) {
	length := 6
	/*
		.#.#.#.#.#.#.#.#.#.#.#.#
		....#.....#.....#.....#.
		.#..#..#..#..#..#..#..#.
		....#.....#.....#.....#.
	*/
	trees := map[int][]int{
		0: {1, 3, 5},
		1: {4},
		2: {1, 4},
		3: {4},
	}
	right := 5
	down := 1

	hits := scanTrees(length, trees, right, down)
	if hits != 1 {
		t.Log("Error, expect 1 hits, got", hits)
		t.Fail()
	}
}

func TestScanTreesMultipleExtensionNoHit(t *testing.T) {
	length := 3
	/*
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
	*/
	trees := map[int][]int{
		0: {2},
		1: {1},
		2: {0},
		3: {2},
		4: {1},
		5: {0},
		6: {2},
		7: {1},
		8: {0},
	}
	right := 2
	down := 1

	hits := scanTrees(length, trees, right, down)
	if hits != 0 {
		t.Log("Error, expect 0 hits, got", hits)
		t.Fail()
	}
}

func TestScanTreesMultipleExtensionHit(t *testing.T) {
	length := 3
	/*
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
		..#..#..#..#..#..#..#..#
		.#..#..#..#..#..#..#..#.
		#..#..#..#..#..#..#..#..
	*/
	trees := map[int][]int{
		0: {2},
		1: {1},
		2: {0},
		3: {2},
		4: {1},
		5: {0},
		6: {2},
		7: {1},
		8: {0},
	}
	right := 3
	down := 1

	hits := scanTrees(length, trees, right, down)
	if hits != 3 {
		t.Log("Error, expect 3 hits, got", hits)
		t.Fail()
	}
}
