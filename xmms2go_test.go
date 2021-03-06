package xmms2go

import (
	"errors"
	"os"
	"testing"
)

func TestValue(t *testing.T) {
	// Start from None.
	vn := NewValueFromNone()
	defer vn.Unref()

	ve := NewValueFromError(errors.New("ValueError Test"))
	defer ve.Unref()
	t.Log("ve is error:", ve.IsError())
	veo, err := ve.GetError()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test error value:", veo)

	vi64 := NewValueFromInt64(17)
	defer vi64.Unref()
	vi64o, err := vi64.GetInt64()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test int64 value:", vi64o)

	vi32 := NewValueFromInt32(23)
	defer vi32.Unref()
	vi32o, err := vi32.GetInt32()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test int32 value:", vi32o)

	vf64 := NewValueFromFloat64(1.4)
	defer vf64.Unref()
	vf64o, err := vf64.GetFloat64()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test float64 value:", vf64o)

	vf32 := NewValueFromFloat32(1.5)
	defer vf32.Unref()
	vf32o, err := vf32.GetFloat32()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test float32 value:", vf32o)

	vs := NewValueFromString("Test string")
	defer vs.Unref()
	vso, err := vs.GetString()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test string value:", vso)

	vb := NewValueFromBytes([]byte("Test\tTest"))
	defer vb.Unref()
	vbo, err := vb.GetBytes()
	if err != nil {
		t.Error(err)
	}
	t.Log("Got test bytes value:", vbo, string(vbo))

	va := NewValueFromAny(func() string { return "Okay" })
	defer va.Unref()
	vao, err := va.GetAny()
	if err != nil {
		t.Error(err)
	}
	// Holy! Go type interface{} is Okay!
	t.Log("Got test anytype value: func() ->", vao.(func() string)())

}

func TestList(t *testing.T) {
	var slice1 []interface{}
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, int64(i))
	}
	t.Log("Slice1=", slice1)

	li64 := NewList()
	defer li64.Unref()
	err := li64.FromSlice(slice1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", li64.GetSize())

	slice2, err := li64.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice2=", slice2)

	var slice3 []interface{}
	var empty interface{}
	t.Log("empty=", empty)
	for i := 0; i < 10; i++ {
		slice3 = append(slice3, empty)
	}
	t.Log("Slice3=", slice3)
	lempty := NewList()
	defer lempty.Unref()
	err = lempty.FromSlice(slice3)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", lempty.GetSize())

	slice4, err := lempty.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice4=", slice4)

	var slice5 []interface{}
	bytes := ([]byte)("Test")
	for i := 0; i < 10; i++ {
		slice5 = append(slice5, bytes)
	}
	t.Log("Slice5=", slice5)
	lb := NewList()
	defer lb.Unref()
	err = lb.FromSlice(slice5)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", lb.GetSize())

	slice6, err := lb.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice6=", slice6)

	var slice7 []interface{}
	b := byte('a')
	for i := 0; i < 10; i++ {
		slice7 = append(slice7, b)
	}
	t.Log("Slice7=", slice7)
	lbb := NewList()
	defer lbb.Unref()
	err = lbb.FromSlice(slice7)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", lbb.GetSize())

	slice8, err := lbb.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice8=", slice8)

	var slice9 []interface{}
	bi := make([]int, 2)
	bi[0] = 1
	bi[1] = 2
	for i := 0; i < 10; i++ {
		slice9 = append(slice9, bi)
	}
	t.Log("Slice9=", slice9)
	lbbi := NewList()
	defer lbbi.Unref()
	err = lbbi.FromSlice(slice9)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", lbbi.GetSize())

	slice10, err := lbbi.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice10=", slice10)

	var slice11 []interface{}
	for i := 10; i > 0; i-- {
		slice11 = append(slice11, float32(i)*3.14)
	}
	t.Log("Slice11=", slice11)
	ls := NewList()
	defer ls.Unref()
	err = ls.FromSlice(slice11)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", ls.GetSize())

	// Do sort
	t.Log("Now test Sort(func)")
	ls.RestrictType(TypeFloat)
	f := func(a *Value, b *Value) int {
		va, _ := a.GetFloat32()
		vb, _ := b.GetFloat32()
		return int(va - vb)
	}
	err = ls.Sort(f)
	if err != nil {
		t.Error(err)
	}
	slice12, err := ls.ToSlice()
	if err != nil {
		t.Error(err)
	}
	t.Log("Slice12=", slice12)
}

func TestDict(t *testing.T) {
	map1 := make(map[string]interface{})
	map1["A"] = 1
	map1["B"] = 2
	map1["C"] = 3
	t.Log("map1=", map1)
	di1 := NewDict()
	defer di1.Unref()
	err := di1.FromMap(map1)
	if err != nil {
		t.Error(err)
	}
	t.Log("Size=", di1.GetSize())

	map2, err := di1.ToMap()
	if err != nil {
		t.Error(err)
	}
	t.Log("map2=", map2)
}

func TestClient(t *testing.T) {
	X, err := NewXmms2Client("xmms2go-test")
	defer X.Unref()
	if err != nil {
		t.Error(err)
	}

	err = X.Connect(os.Getenv("XMMS_PATH"))
	if err != nil {
		t.Error(err)
	}

	err = X.Play()
	if err != nil {
		t.Error(err)
	}

	i, err := X.CurrentID()
	if err != nil {
		t.Error(err)
	}
	t.Log("Current Playing ID:", i)
}
