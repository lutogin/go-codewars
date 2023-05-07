package reverseLinkedList

import "fmt"

/**
Need to revers linked list.
*/

type LinkedList struct {
	val  interface{}
	next *LinkedList
}

func GetExampleData() *LinkedList {
	return &LinkedList{
		val: 1,
		next: &LinkedList{
			val: 2,
			next: &LinkedList{
				val:  "3",
				next: nil,
			},
		}}
}

func PrintList(list *LinkedList) {
	current := list
	i := 1

	for current != nil {
		fmt.Printf("Node: %d; Value: %v \n", i, current.val)
		current = current.next
		i++
	}
}

func Run(list *LinkedList) *LinkedList {
	fmt.Println("Initial list:")
	PrintList(list)
	fmt.Println("======================")

	var prevNode *LinkedList
	var nextNode *LinkedList
	currentNode := list

	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	fmt.Println("Reversed list:")
	PrintList(prevNode)
	fmt.Println("======================")

	return prevNode
}
