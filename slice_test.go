package gofin

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	for _, test := range []struct {
		Input    []int
		Function func(a int) int
		Output   []int
	}{
		{[]int{1, 2, 3}, func(i int) int { return i + 1 }, []int{2, 3, 4}},
		{[]int{1, 2, 3}, func(i int) int { return i - 1 }, []int{0, 1, 2}},
	} {
		Map(test.Input, test.Function)
		assert.Equal(t, test.Output, test.Input)
	}
}

func TestForEach(t *testing.T) {
	for _, test := range []struct {
		Input    []string
		Function func(k int, v string)
		Output   string
	}{
		{[]string{"hello", "world"}, func(k int, v string) { fmt.Printf("%d - %s\n", k, v) }, "0 - hello\n1 - world\n"},
		{[]string{"hello", "world"}, func(k int, v string) { fmt.Printf("%d / %s\n", k, v) }, "0 / hello\n1 / world\n"},
	} {
		r, w, _ := os.Pipe()
		os.Stdout = w
		ForEach(test.Input, test.Function)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		assert.Equal(t, test.Output, string(out))
	}
}

func TestFilter(t *testing.T) {
	for _, test := range []struct {
		Input    []int
		Function func(a int) bool
		Output   []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i > 2 }, []int{3}},
		{[]int{2, 3, 5}, func(i int) bool { return i%2 == 0 }, []int{2}},
	} {
		assert.Equal(t, test.Output, Filter(test.Input, test.Function))
	}
}

func TestRemoveIndex(t *testing.T) {
	for _, test := range []struct {
		Input  []int
		Index  int
		Output []int
	}{
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 0, []int{2, 3}},
	} {
		assert.Equal(t, test.Output, RemoveIndex(test.Input, test.Index))
	}
}

func TestRemoveMatching(t *testing.T) {
	for _, test := range []struct {
		Input    []int
		Function func(a int) bool
		Output   []int
	}{
		{[]int{1, 2, 3}, func(i int) bool { return i%2 == 0 }, []int{1, 3}},
	} {
		assert.Equal(t, test.Output, RemoveMatching(test.Input, test.Function))
	}
}

func TestContains(t *testing.T) {
	for _, test := range []struct {
		Input    []int
		Function int
		Output   bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 5, false},
	} {
		assert.Equal(t, test.Output, Contains(test.Input, test.Function))
	}
}

func TestIndexOf(t *testing.T) {
	for _, test := range []struct {
		Input   []int
		Element int
		Output  int
	}{
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 1, 0},
	} {
		assert.Equal(t, test.Output, IndexOf(test.Input, test.Element))
	}
}

func TestChunk(t *testing.T) {
	for _, test := range []struct {
		Input  []int
		Size   int
		Output [][]int
	}{
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
	} {
		assert.Equal(t, test.Output, Chunk(test.Input, test.Size))
	}
}
