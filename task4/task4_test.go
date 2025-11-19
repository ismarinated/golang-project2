package cache

import (
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	var tests = []struct {
		name      string
		size      uint
		listK     []string
		listV     []int
		wantValue int
		wantBool  bool
	}{
		{
			"size 4",
			4,
			[]string{"a", "b", "c", "d"},
			[]int{1, 2, 3, 4},
			1,
			true,
		},
		{
			"size 1",
			1,
			[]string{"a"},
			[]int{1},
			1,
			true,
		},
		{
			"size 0",
			0,
			[]string{"a"},
			[]int{1},
			0,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCache[string, int](test.size)

			for i := range test.listK {
				c.Set(test.listK[i], test.listV[i])
			}

			gotValue, gotBool := c.Get(test.listK[0])

			if gotValue != test.wantValue && gotBool != test.wantBool {
				t.Errorf("%s: got value %v, want value %v\ngot bool %v, want bool %v", test.name, gotValue, test.wantValue, gotBool, test.wantBool)
			}
		})
	}
}

func TestCacheClear(t *testing.T) {
	var tests = []struct {
		name      string
		size      uint
		listK     []string
		listV     []int
		wantValue int
		wantBool  bool
	}{
		{
			"size 4",
			4,
			[]string{"a", "b", "c", "d"},
			[]int{1, 2, 3, 4},
			0,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCache[string, int](test.size)

			for i := range test.listK {
				c.Set(test.listK[i], test.listV[i])
			}

			c.Clear()

			for i := range test.listK[:test.size] {
				c.Set(test.listK[i], test.listV[i])
			}

			gotValue, gotBool := c.Get(test.listK[test.size-1])

			if gotValue != test.wantValue && gotBool != test.wantBool {
				t.Errorf("%s: got value %v, want value %v\ngot bool %v, want bool %v", test.name, gotValue, test.wantValue, gotBool, test.wantBool)
			}
		})
	}
}

func TestCacheConcurrency(t *testing.T) {
	var tests = []struct {
		name  string
		size  uint
		listK []string
		listV []int
	}{
		{
			"size 4",
			4,
			[]string{"a", "b", "c", "d"},
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCache[string, int](test.size)

			for i := range test.listK {
				c.Set(test.listK[i], test.listV[i])
			}

			var wg sync.WaitGroup
			wg.Go(func() {
				gotValue, gotBool := c.Get(test.listK[0])

				if gotValue != test.listV[0] && gotBool != false {
					t.Errorf("%s: got value %v, want value %v\ngot bool %v, want bool %v", test.name, gotValue, test.listV[0], gotBool, false)
				}

			})

			wg.Go(func() {
				gotValue, gotBool := c.Get(test.listK[test.size-1])

				if gotValue != test.listV[test.size-1] && gotBool != true {
					t.Errorf("%s: got value %v, want value %v\ngot bool %v, want bool %v", test.name, gotValue, test.listV[test.size-1], gotBool, true)
				}

			})

			wg.Wait()
		})
	}
}
