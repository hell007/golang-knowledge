package hashTables

import (
	"fmt"
	"testing"
)

//go test -v ht.go ht_test.go
func TestHt(t *testing.T) {

	// Test Append
	ht := New(1000)
	ht.Append("a", "bar")
	ht.Append("b", "buzz")
	ht.Append("c", "wayne")
	ht.Append("d", "parker")
	ht.Append("e", "kent")

	// Test simple get
	val, err := ht.Get("a")
	fmt.Println("ht.Get(a)==", val, err)

	// Test updated
	ht.Append("d", "bob")
	val, err = ht.Get("d")
	fmt.Println("ht.Get(d) updated==", val, err)
	//if err != nil && val == "bob" || ht.Size != 5 {
	//	fmt.Println(val, err)
	//	t.Error()
	//}

	// Test Remove
	ht.Remove("d")
	val, err = ht.Get("d")
	fmt.Println("ht.Get(d) Remove==", val, err, ht.Size)
	//val, err = ht.Get("d")
	//if val != "" || err == nil || ht.Size != 4 {
	//	t.Error()
	//}

	// Test Each
	counter := 0
	f := func(a *item) {
		counter++
		fmt.Println("Each==", a)
	}

	ht.Each(f)
	if counter != 4 {
		t.Error()
	}

}

func TestHash(t *testing.T) {
	test := "Hello World!"
	fmt.Println("hashCode(test)==", hashCode(test))

	if hashCode(test) != 969099747 {
		t.Error()
	}
}
