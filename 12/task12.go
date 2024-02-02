// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.
package main

import "fmt"

func removeDuplicate(slice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range slice {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func main() {
	newSlice := []string{"cat", "cat", "dog", "cat", "tree"}
	result := removeDuplicate(newSlice)
	fmt.Println(result)

}
