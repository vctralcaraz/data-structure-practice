package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

/*
	I will need to find more methods here - https://www.w3schools.com/java/java_ref_linkedlist.asp
*/

/*	**Insertion Methods**
	addFirst(data)
	addLast(data)
	addAtIndex(data, index)
*/
func (l *LinkedList) addFirst(data string) {
	
	temp := new(Node)
	temp.data = data

	if l.head == nil {
		l.head = temp
	} else {
		temp.next = l.head
		l.head = temp
	}
	l.size++
}

func (l *LinkedList) addLast(data string) {

	temp := new(Node)
	temp.data = data

	if l.head == nil {
		l.head = temp
	} else {
		i := l.head
		for ; i.next != nil; i = i.next {
		}
		i.next = temp
	}
	l.size++
}

func (l *LinkedList) addAtIndex(data string, index int) {
	
	temp := new(Node)
	temp.data = data

	if index > l.size || index < 0 {
		fmt.Println("Index out of range\n")
		return
	}

	if l.head == nil {
		l.head = temp
		l.size++
		return
	}

	i := l.head
	for k := 0; k < index - 1; k++ {
		i = i.next
	}

	j := i.next
	temp.next = j
	i.next = temp
	l.size++
}

/*	**Deletion Methods**
	removeFirst()
	removeLast()
	removeAtIndex(index)
	remove(data) - removes the first node containing the specified data
*/
func (l *LinkedList) removeFirst() {
	
	if l.head != nil {
		i := l.head
		l.head = i.next
		i = nil
		l.size--
		return
	}

	fmt.Println("Nothing to remove")
}
func (l *LinkedList) removeLast() {
	
	if l.head != nil {
		i := l.head
		for ; i.next.next != nil; i = i.next {}
		j := i.next
		i.next = nil
		fmt.Println("Deleted", j.data)
		j = nil
		l.size--
		return
	}

	fmt.Println("Nothing to remove")
}
func (l *LinkedList) removeAtIndex(index int) {

	if index > l.size || index < 0 {
		fmt.Println("Index out of range\n")
		return
	}

	if l.size == 0 {
		fmt.Println("This list is empty")
		return
	}

	i := l.head
	for k := 0; k < index - 1; k++ {
		i = i.next
	}

	j := i.next
	i.next = i.next.next
	fmt.Println("Deleted", j.data)
	j = nil
	l.size--
}

func (l *LinkedList) remove(data string) {

	i := l.head
	if l.head.data == data {
		l.head = i.next
		i = nil
		l.size--
		return
	} else {
		for ; i.next != nil; i = i.next {
			if i.next.data == data {
				j := i.next
				i.next = j.next
				j = nil
				l.size--
				return
			}
		}
	}
}

/*	**Search Methods**
	contains(data)		checks if the list contains a node with the specified data - returns true or false
	indexOf(data)		Returns the index of the first node containing the specified data
	get(index)		returns the data stored in the node at the specified index
	search(data)		returns the data of the first node containing the specified data
*/
func (l LinkedList) contains(data string) bool { return false }
func (l LinkedList) indexOf(data string) int { return -1 }
func (l LinkedList) get(index int) *Node { return nil }

func (l LinkedList) search(data string) *Node {

	i := l.head
	for ; i != nil; i = i.next {
		if i.data == data {
			return i
		}
	}
	return nil
}

/*	**Traversal Methods**
	printAll()
*/
func (l LinkedList) printAll() {
	if l.size > 0 {
		i := l.head
		for ; i != nil; i = i.next {
			fmt.Println(i.data)
		}
	} else {
		fmt.Println("This list is empty")
	}
}

/*	**Misc Methods**
	isEmpty()		checks if the list is empty - returns true or false
	getSize()		returns the number of nodes in the list
	reverse()		Reverses the order of the nodes in the list - returns a new list
*/
func (l LinkedList) isEmpty() bool { return false }
func (l LinkedList) getSize() int { return 0 }
func (l *LinkedList) reverse() {}

func main() {
	states := new(LinkedList)


	states.addLast("California")
	states.addLast("Arizona")
	states.addLast("Nevada")
	states.addLast("Washington")
	states.addLast("Oregon")
	states.addLast("Colorado")

	states.addFirst("Florida")

	states.addAtIndex("Georgia", 5)

	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")

	states.remove("California")

	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")

	states.removeLast()
	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")

	states.removeAtIndex(2)
	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")
	
	state := states.search("Texas")
	if state == nil {
		fmt.Println("The state you are looking for doesn't exist")
	} else {
		fmt.Println("The state you are looking for is at ", state)
	}

	state = states.search("Colorado")
	if state == nil {
		fmt.Println("The state you are looking for doesn't exist")
	} else {
		fmt.Println("The state you are looking for is at ", state)
	}
}
