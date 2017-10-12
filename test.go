package main

import (
	"fmt"
	"log"

	"github.com/sidmishraw/gunc/gunc"
)

type a = gunc.A
type b = gunc.B
type c = gunc.C

func addints(x a, y b) c {
	return x.(int) + y.(int)
}

// curried addints for 2 -- partially added
var add2 = gunc.Curry(addints, a(2))

// TestCurryAdd tests currying for add function
func TestCurryAdd() {
	// testing currying
	// the currying is too --- tacky because of Go and my utter noobness
	// this should be fixed in coming revisions
	if add2to1 := add2(1); add2to1 != 3 {
		log.Fatalf("Currying failed, expected %d got %d", 3, add2to1)
	}
	if add2to2 := add2(2); add2to2 != 4 {
		log.Fatalf("Currying failed, expected %d got %d", 4, add2to2)
	}
	if add2to99 := add2(99); add2to99 != 101 {
		log.Fatalf("Currying failed, expected %d got %d", 101, add2to99)
	}
	log.Println("Curry succeeded...")
}

// TestMap tests the Map function
func TestMap() {
	xs := []a{1, 2, 3, 4, 5}
	bs := gunc.Map(func(x a) b { return x.(int) + 2 }, xs)
	fmt.Printf("mapped:: %s", bs)
	for i := range bs {
		if bs[i].(int) != xs[i].(int)+2 {
			log.Fatalf("mapping failed:: expected %d got %d", (xs[i].(int) + 2), bs[i].(int))
		}
	}
	log.Println("Map succeeded...")
}

// TestFoldl tests the Foldl function
func TestFoldl() {
	xs := []a{2, 4, 8}
	bs := gunc.Foldl(func(x a, y b) b { return x.(int) / y.(int) }, b(1), xs)
	fmt.Printf("foldl'ed:: %s\n", bs)
	if bs != 4 {
		log.Fatalf("foldl failed:: expected 4 got %d\n", bs)
	} else {
		fmt.Printf("Foldl succeeded...\n")
	}
}

// TestFoldr tests the Foldr function
// here the way foldr operates, the data type needs to be double since fractions are
// involved
func TestFoldr() {
	xs := []a{2.0, 4.0, 8.0}
	bs := gunc.Foldr(func(x a, y b) b { return x.(float64) / y.(float64) }, b(float64(1.0)), xs)
	fmt.Printf("foldr'ed:: %s\n", bs)
	if bs != 4.0 {
		log.Fatalf("foldr failed:: expected 4 got %d\n", bs)
	} else {
		fmt.Printf("Foldr succeeded...\n")
	}
}

func main() {
	TestCurryAdd()
	TestMap()
	TestFoldl()
	TestFoldr()
}
