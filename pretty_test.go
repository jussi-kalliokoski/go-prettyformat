package prettyformat

import (
	"testing"
)

func TestPrimitives(t *testing.T) {
	expectResultToMatch(t, "asdfasda\nasd", `"asdfasda\nasd"`)
	expectResultToMatch(t, true, "true")
	expectResultToMatch(t, false, "false")
	expectResultToMatch(t, int(1234), "1234")
	expectResultToMatch(t, int(-4321), "-4321")
	expectResultToMatch(t, uint(59912), "59912")
	expectResultToMatch(t, int8(112), "112")
	expectResultToMatch(t, int8(-15), "-15")
	expectResultToMatch(t, uint8(123), "123")
	expectResultToMatch(t, int16(1234), "1234")
	expectResultToMatch(t, int16(-4321), "-4321")
	expectResultToMatch(t, uint16(30123), "30123")
	expectResultToMatch(t, int32(1234), "1234")
	expectResultToMatch(t, int32(-4321), "-4321")
	expectResultToMatch(t, uint32(59912), "59912")
	expectResultToMatch(t, int64(1234), "1234")
	expectResultToMatch(t, int64(-4321), "-4321")
	expectResultToMatch(t, uint64(59912), "59912")
	expectResultToMatch(t, float32(1234.5677), "1234.5677")
	expectResultToMatch(t, float32(-4321.1235), "-4321.1235")
	expectResultToMatch(t, float64(1234.56789), "1234.56789")
	expectResultToMatch(t, float64(-4321.12345), "-4321.12345")
	expectResultToMatch(t, complex(float32(1234.5677), float32(-4321.1235)), "(1234.5677-4321.1235i)")
	expectResultToMatch(t, complex(float32(-4321.1235), float32(1234.5677)), "(-4321.1235+1234.5677i)")
	expectResultToMatch(t, complex(float64(1234.56789), float64(-4321.12345)), "(1234.56789-4321.12345i)")
	expectResultToMatch(t, complex(float64(-4321.12345), float64(1234.56789)), "(-4321.12345+1234.56789i)")
}

func TestSlices(t *testing.T) {
	expectResultToMatch(t, []int{}, "[]int{}")
	expectResultToMatch(t,
		[]string{
			"hello",
			"world",
		},
		`[]string{
  "hello",
  "world",
}`)
	expectResultToMatch(t,
		[]interface{}{
			"hello",
			1234,
		},
		`[]interface{}{
  (string)"hello",
  (int)1234,
}`)
}

func TestMaps(t *testing.T) {
	expectResultToMatch(t, map[string]int{}, "map[string]int{}")
	expectResultToMatch(t,
		map[int]string{
			0: "asd",
			3: "bar",
		},
		`map[int]string{
  0: "asd",
  3: "bar",
}`)
	expectResultToMatch(t,
		map[interface{}]string{
			0:     "asd",
			"foo": "bar",
		},
		`map[interface{}]string{
  (string)"foo": "bar",
  (int)0: "asd",
}`)
	expectResultToMatch(t,
		map[string]interface{}{
			"foo": "asd",
			"bar": 123,
		},
		`map[string]interface{}{
  "bar": (int)123,
  "foo": (string)"asd",
}`)
}

func TestArrays(t *testing.T) {
	expectResultToMatch(t,
		[2]string{
			"hello",
			"world",
		},
		`[2]string{
  "hello",
  "world",
}`)
	expectResultToMatch(t,
		[3]interface{}{
			"hello",
			1234,
			int64(12345),
		},
		`[3]interface{}{
  (string)"hello",
  (int)1234,
  (int64)12345,
}`)
}

func TestPointers(t *testing.T) {
	v := 1
	expectResultToMatch(t, &v, "&1")
}

func TestStructs(t *testing.T) {
	type foo1 struct {
	}

	expectResultToMatch(t, foo1{}, `foo1{}`)

	type foo2 struct {
		dsa int
	}

	expectResultToMatch(t, foo2{123}, `foo2{}`)

	type foo3 struct {
		Asd string
		bar int
		Foo int
	}

	expectResultToMatch(t,
		foo3{
			"foo",
			123,
			321,
		}, `foo3{
  Asd: "foo",
  Foo: 321,
}`)

	type foo4 struct {
		Asd interface{}
	}

	expectResultToMatch(t,
		foo4{
			"foo",
		}, `foo4{
  Asd: (string)"foo",
}`)
}

func TestNested(t *testing.T) {
	type foo1 struct {
		Foo interface{}
	}

	expectResultToMatch(t,
		[1]foo1{
			foo1{
				Foo: []*foo1{
					&foo1{
						Foo: map[string]interface{}{
							"foo": foo1{
								Foo: "123",
							},
						},
					},
				},
			},
		},
		`[1]foo1{
  foo1{
    Foo: []*foo1{
      &foo1{
        Foo: map[string]interface{}{
          "foo": foo1{
            Foo: (string)"123",
          },
        },
      },
    },
  },
}`)
}

func expectResultToMatch(t *testing.T, value interface{}, expected string) {
	t.Helper()

	received, err := Format(value)

	if err != nil {
		t.Error(err)
		return
	}

	if expected != received {
		t.Errorf("Expected `%s`, got `%s`", expected, received)
	}
}
