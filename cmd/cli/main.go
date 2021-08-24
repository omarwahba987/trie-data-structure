package main

import (
	"fmt"
	"va_test_d/internal"
)

func main() {
	trie := internal.InitTrie()
	fmt.Println("construct a Trie for how many persons ?")
	numPersons := 0
	_, err := fmt.Scan(&numPersons)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < numPersons; i++ {
		var name, address string
		fmt.Printf("enter name of person number %v name \n ", i+1)
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("enter address of person number %v name \n ", i+1)
		_, err = fmt.Scan(&address)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = trie.Insert(name, address)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
entering:
	fmt.Println("\n if you want to visualize the trie enter 1 \n" +
		" if you want to search for a person address press 2 \n" +
		"if you want to exit press CTRL+c")
	var input int
	_, err = fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if input == 1 {
		fmt.Println(trie)
	} else if input == 2 {
		fmt.Println("enter the person name")
		var name string
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println(err)
			return
		}
		address, err := trie.Search(name)

		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(address)
		}
	}
	goto entering

}
