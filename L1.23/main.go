package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removePtrElement(slice []*int, i int) []*int {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = nil // Предотвращаем утечку памяти
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("Исходный слайс:", nums)
	nums = remove(nums, 2) // удалим элемент с индексом 2 (значение 30)
	fmt.Println("После удаления:", nums)

	a, b, c, d := 1, 2, 3, 4
	ptrs := []*int{&a, &b, &c, &d}
	ptrs = removePtrElement(ptrs, 1) // Удаляем элемент с индексом 1 (указатель на b)
	fmt.Println(ptrs)
}
