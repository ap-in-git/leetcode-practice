package main

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	productExceptSelf([]int{1, 2, 3, 4})
	isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}})
}

func removeDuplicates(items []int) []int {

	var returnItems []int
	//itemsMap := make(map[int]int, len(items))

	i := 0

	for i < len(items) {
		currentValue := items[i]
		if i == len(items)-1 {
			returnItems = append(returnItems, currentValue)
			break
		}
		nextValue := items[i+1]
		if currentValue == nextValue {
			i = i + 2
		} else {
			i = i + 1
			returnItems = append(returnItems, currentValue)
		}
	}

	//for i := 0; i < len(items); i++ {
	//
	//	if currentValue == nextValue {
	//
	//	}
	//}
	//for index, item := range items {
	//	currentValue := items[index]
	//	nextValue := items[index+1]
	//	if currentValue != nextValue {
	//		returnItems = append(returnItems, item)
	//	}
	//	//if index  {
	//	//
	//	//}
	//	//val, ok := itemsMap[item]
	//	//if ok {
	//	//delete()
	//	//itemsMap[item] = val + 1
	//	//} else {
	//	//	itemsMap[item] = 1
	//	//}
	//}

	//for key, val := range itemsMap {
	//	if val == 1 {
	//		returnItems = append(returnItems, key)
	//	}
	//}

	return returnItems
}
