package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) Append(newValue int) {
	newNode := Node{value: newValue, next: nil}

	if list.head == nil {
		list.head = &newNode
		list.tail = &newNode
		list.length++
	} else {
		list.tail.next = &newNode
		list.tail = &newNode
		list.length++
	}

}

func (list *LinkedList) InsertAfter(target int, newValue int) {
	newNode := Node{value: newValue, next: nil}
	current_node := list.head
	for current_node != nil {
		if current_node.value == target {
			if current_node == list.tail {
				list.tail = &newNode
			}
			newNode.next = current_node.next
			current_node.next = &newNode
			list.length++
			break
		}

		current_node = current_node.next
	}
}

func (list *LinkedList) InsertBefore(target int, newValue int) {
	newNode := Node{value: newValue, next: nil}
	var previous_node *Node
	current_node := list.head
	for current_node != nil {
		if current_node.value == target {
			if current_node == list.head {
				list.head = &newNode
				newNode.next = current_node
			} else {
				previous_node.next = &newNode
				newNode.next = current_node
			}
			list.length++
			break
		}
		previous_node = current_node
		current_node = current_node.next
	}
}

func (list *LinkedList) Print() {
	current_node := list.head
	for current_node != nil {
		fmt.Println(current_node.value)
		current_node = current_node.next
	}
}

func (list *LinkedList) Delete(value int) {
	current_node := list.head
	var previous_node *Node
	// var next_node *Node

	for current_node != nil {
		if current_node.value == value {
			if list.length == 1 {
				list.head = current_node.next
				list.tail = current_node.next
				list.length--
				break
			}
			if current_node == list.head {
				list.head = current_node.next
				list.length--
				break
			}
			if current_node == list.tail {
				list.tail = previous_node
				previous_node.next = nil
				list.length--
				break
			}
			previous_node.next = current_node.next
			list.length--
			break
		}
		previous_node = current_node
		current_node = current_node.next
		// next_node = current_node.next
	}
}

func (list *LinkedList) IsExist(target int) bool {
	current_node := list.head
	for current_node != nil {
		if current_node.value == target {
			return true
		}
		current_node = current_node.next
	}
	return false
}

func (list *LinkedList) Head() int {
	if list.length == 0 {
		return 0
	}
	return list.head.value
}

func (list *LinkedList) Tail() int {
	if list.length == 0 {
		return 0
	}
	return list.tail.value
}

func (list *LinkedList) Length() int {
	return list.length
}

func (list *LinkedList) Reverse() {
	temp := list.head
	list.head = list.tail
	list.tail = temp
	current_node := list.tail
	var previous_node *Node
	previous_node = nil

	for range list.length {
		temp := current_node.next
		current_node.next = previous_node
		previous_node = current_node
		current_node = temp
	}
}

func main() {
	list := LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}

	list.Append(6)
	list.Append(5)
	list.Append(3)
	list.Append(22)
	list.InsertAfter(22, 4)
	list.InsertBefore(6, 1)
	list.Append(65)
	list.Delete(65)
	list.Append(33)

	fmt.Printf("Head: %v | Tail: %v | Length: %v\n", list.Head(), list.Tail(), list.Length())
	list.Print()
	fmt.Println(list.IsExist(23))

	list.Reverse()

	list.Append(16)
	list.Append(15)
	list.Append(13)
	list.Append(122)
	list.InsertAfter(122, 4)
	list.InsertBefore(16, 1)
	list.Append(165)
	list.Delete(165)
	list.Append(133)

	fmt.Printf("Head: %v | Tail: %v | Length: %v\n", list.Head(), list.Tail(), list.Length())
	list.Print()
}
