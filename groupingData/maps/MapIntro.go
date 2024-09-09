package maps

import "fmt"

func MapsIntro() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	an := make(map[string]int) // again defines a map rather than populate
	an["Lucas"] = 28
	an["Steph"] = 37
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	fmt.Println(len(an))

	//  --- or ---
	m := make(map[string]int)
	m["todd"] = 42
	m["henry"] = 16

	fmt.Println("Henry's age is ", m["henry"])

	for k := range m {
		fmt.Printf("just the keys: %s\n", k)
	}

	for k, v := range m {
		fmt.Printf("%s is %d years old\n", k, v)
	}

	for _, v := range m {
		fmt.Printf("just the values: %d\n", v)
	}

	// comma ok idiom
	if v, ok := m["henry"]; ok {
		fmt.Printf("the value prints", v)
	} else {
		fmt.Printf("Key didn't exist")
	}

	// idiomatic code in go: deal error soon as it occurs
	if v, ok := m["henry"]; !ok {
		fmt.Printf("Key didn't exist")
	} else {
		fmt.Printf("the value prints", v)
	}

	//delete
	fmt.Println("delete")
	m["Shakespeare"] = 459
	fmt.Printf("%#v\n", m)
	delete(m, "Shakespeare")
	fmt.Printf("%#v\n", m)
	delete(m, "Shakespeare") // no panic -- if nil/not present then delete is a no-op

	// len
	fmt.Println("len: ", len(m))

	// for range with a SLICE
	xi := []int{42, 43, 44}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	for _, v := range xi {
		fmt.Println(v)
	}

	for i := range xi {
		fmt.Println(i)
	}
}
