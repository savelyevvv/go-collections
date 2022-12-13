package list

import (
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	t.Run("can append one element", func(t *testing.T) {
		var ss []string
		want := []string{"A"}
		Append(&ss, "A")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can append several elements", func(t *testing.T) {
		var ss []string
		want := []string{"A", "B", "C"}
		Append(&ss, "A")
		Append(&ss, "B")
		Append(&ss, "C")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestCount(t *testing.T) {
	t.Run("no elements equal to search key", func(t *testing.T) {
		ss := []string{"A", "B", "A", "C", "A"}
		want := 0
		got := Count(&ss, "D", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("one element equal to search key", func(t *testing.T) {
		ss := []string{"A", "B", "A", "C", "A"}
		want := 1
		got := Count(&ss, "B", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("several elements equal to search key", func(t *testing.T) {
		ss := []string{"A", "B", "A", "C", "A"}
		want := 3
		got := Count(&ss, "A", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("can insert to empty then cap(s) == 0", func(t *testing.T) {
		var ss []string
		want := []string{"A"}
		Insert(&ss, 0, "A")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert to empty then cap(s) > 0", func(t *testing.T) {
		ss := make([]string, 0, 2)
		want := []string{"A"}
		Insert(&ss, 0, "A")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert to the beginning then len(s) == cap(s)", func(t *testing.T) {
		ss := []string{"B"}
		want := []string{"A", "B"}
		Insert(&ss, 0, "A")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert to the beginning then len(s) < cap(s)", func(t *testing.T) {
		ss := make([]string, 1, 2)
		ss[0] = "B"
		want := []string{"A", "B"}
		Insert(&ss, 0, "A")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert to the end then len(s) == cap(s)", func(t *testing.T) {
		ss := []string{"B"}
		want := []string{"B", "C"}
		Insert(&ss, 1, "C")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert to the end then len(s) < cap(s)", func(t *testing.T) {
		ss := make([]string, 1, 2)
		ss[0] = "B"
		want := []string{"B", "C"}
		Insert(&ss, 1, "C")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert in the middle then len(s) == cap(s)", func(t *testing.T) {
		ss := []string{"A", "C"}
		want := []string{"A", "B", "C"}
		Insert(&ss, 1, "B")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can insert in the middle then len(s) < cap(s)", func(t *testing.T) {
		ss := make([]string, 2, 3)
		ss[0] = "A"
		ss[1] = "C"
		want := []string{"A", "B", "C"}
		Insert(&ss, 1, "B")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestAlignIndex(t *testing.T) {
	t.Run("index is positive and within range", func(t *testing.T) {
		i := 2
		n := 5
		want := 2
		got, _ := alignIndex(i, n)
		if want != got {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("index is positive and out of range", func(t *testing.T) {
		i := 6
		n := 5
		want := 5
		got, _ := alignIndex(i, n)
		if want != got {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("index is negative and within range", func(t *testing.T) {
		i := -4
		n := 5
		want := 1
		got, _ := alignIndex(i, n)
		if want != got {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("index is negative and out of range", func(t *testing.T) {
		i := -6
		n := 5
		want := 0
		got, _ := alignIndex(i, n)
		if want != got {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("then slice is empty", func(t *testing.T) {
		ss := []string{}
		want := []string{}
		Reverse(&ss)
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("then slice has one element", func(t *testing.T) {
		ss := []string{"A"}
		want := []string{"A"}
		Reverse(&ss)
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("then slice has several elements", func(t *testing.T) {
		ss := []string{"A", "B", "C", "D"}
		want := []string{"D", "C", "B", "A"}
		Reverse(&ss)
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestExtend(t *testing.T) {
	t.Run("cat extend a nil slice", func(t *testing.T) {
		var ss []string
		want := []string{"A", "B"}
		Extend(&ss, "A", "B")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("cat extend an empty slice", func(t *testing.T) {
		ss := []string{}
		want := []string{"A", "B"}
		Extend(&ss, "A", "B")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can extend with zero arguments", func(t *testing.T) {
		ss := []string{"A", "B"}
		want := []string{"A", "B"}
		Extend(&ss)
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can extend with one argument", func(t *testing.T) {
		ss := []string{"A", "B"}
		want := []string{"A", "B", "C"}
		Extend(&ss, "C")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
	t.Run("can extend with several arguments", func(t *testing.T) {
		ss := []string{"A", "B"}
		want := []string{"A", "B", "C", "D", "E"}
		Extend(&ss, "C", "D", "E")
		got := ss
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, but want %#v\n", got, want)
		}
	})
}

func TestPop(t *testing.T) {
	// # TODO: test panic
	t.Run("one-element slice", func(t *testing.T) {
		ss := []string{"A"}
		want := "A"
		got := Pop(&ss, 0)
		if got != want || !reflect.DeepEqual(ss, []string{}) {
			t.Errorf("%v != %v or %v != %v\n", got, want, ss, []string{})
		}
	})
	t.Run("pop first", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := "A"
		got := Pop(&ss, 0)
		if got != want || !reflect.DeepEqual(ss, []string{"B", "C"}) {
			t.Errorf("%v != %v or %v != %v\n", got, want, ss, []string{})
		}
	})
	t.Run("pop last", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := "B"
		got := Pop(&ss, 1)
		if got != want || !reflect.DeepEqual(ss, []string{"A", "C"}) {
			t.Errorf("%v != %v or %v != %v\n", got, want, ss, []string{})
		}
	})
	t.Run("pop in the middle", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := "C"
		got := Pop(&ss, 2)
		if got != want || !reflect.DeepEqual(ss, []string{"A", "B"}) {
			t.Errorf("%v != %v or %v != %v\n", got, want, ss, []string{})
		}
	})
}

func TestCopy(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var ss []string
		got := Copy(&ss)
		if got != nil {
			t.Errorf("got %v, but want %v", got, nil)
		}
	})
	t.Run("empty slice", func(t *testing.T) {
		ss := []string{}
		want := ss
		got := Copy(&ss)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	})
	t.Run("non-empty slice", func(t *testing.T) {
		ss := []string{"A", "B"}
		want := ss
		got := Copy(&ss)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove first", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := []string{"B", "C"}
		Remove(&ss, "A", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(ss, want) {
			t.Errorf("got %v, want %v", ss, want)
		}
	})
	t.Run("remove last", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := []string{"A", "B"}
		Remove(&ss, "C", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(ss, want) {
			t.Errorf("got %v, want %v", ss, want)
		}
	})
	t.Run("remove in the middle", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := []string{"A", "C"}
		Remove(&ss, "B", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(ss, want) {
			t.Errorf("got %v, want %v", ss, want)
		}
	})
	t.Run("remove if a slice has no such element", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		want := []string{"A", "B", "C"}
		Remove(&ss, "D", func(a, b string) bool { return a == b })
		if !reflect.DeepEqual(ss, want) {
			t.Errorf("got %v, want %v", ss, want)
		}
	})
}

func TestIndex(t *testing.T) {
	t.Run("slice is empty", func(t *testing.T) {
		var ss []string
		got := Index(&ss, "B", func(a, b string) bool { return a == b })
		want := -1
		if got != want {
			t.Errorf("got %d, want %d\n", got, want)
		}
	})
	t.Run("slice contains an element", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		got := Index(&ss, "B", func(a, b string) bool { return a == b })
		want := 1
		if got != want {
			t.Errorf("got %d, want %d\n", got, want)
		}
	})
	t.Run("slice doesn't contain any element", func(t *testing.T) {
		ss := []string{"A", "B", "C"}
		got := Index(&ss, "D", func(a, b string) bool { return a == b })
		want := -1
		if got != want {
			t.Errorf("got %d, want %d\n", got, want)
		}
	})
}
