package boil

import "testing"

func TestGetStructPointers(t *testing.T) {
	t.Parallel()

	o := struct {
		Title string
		ID    *int
	}{
		Title: "patrick",
	}

	ptrs := GetStructPointers(&o, "title", "id")
	*ptrs[0].(*string) = "test"
	if o.Title != "test" {
		t.Errorf("Expected test, got %s", o.Title)
	}
	x := 5
	*ptrs[1].(**int) = &x
	if *o.ID != 5 {
		t.Errorf("Expected 5, got %d", *o.ID)
	}
}

func TestCheckType(t *testing.T) {
	t.Parallel()

	type Thing struct {
	}

	validTest := []struct {
		Input    interface{}
		IsSlice  bool
		TypeName string
	}{
		{&[]*Thing{}, true, "boil.Thing"},
		{[]Thing{}, false, ""},
		{&[]Thing{}, false, ""},
		{Thing{}, false, ""},
		{new(int), false, ""},
		{5, false, ""},
		{&Thing{}, false, "boil.Thing"},
	}

	for i, test := range validTest {
		typ, isSlice, err := checkType(test.Input)
		if err != nil {
			if len(test.TypeName) > 0 {
				t.Errorf("%d) Type: %T %#v - should have succeded but got err: %v", i, test.Input, test.Input, err)
			}
			continue
		}

		if isSlice != test.IsSlice {
			t.Errorf("%d) Type: %T %#v - succeded but wrong isSlice value: %t, want %t", i, test.Input, test.Input, isSlice, test.IsSlice)
		}

		if got := typ.String(); got != test.TypeName {
			t.Errorf("%d) Type: %T %#v - succeded but wrong type name: %s, want: %s", i, test.Input, test.Input, got, test.TypeName)
		}
	}
}