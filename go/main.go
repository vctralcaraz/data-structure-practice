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

func (l *LinkedList) add(data string) {

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

func (l LinkedList) get(data string) *Node {

	i := l.head
	for ; i != nil; i = i.next {
		if i.data == data {
			return i
		}
	}
	return nil
}

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

func main() {
	states := new(LinkedList)

	states.add("California")
	states.add("Arizona")
	states.add("Nevada")
	states.add("Washington")
	states.add("Oregon")
	states.add("Colorado")

	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")

	states.remove("California")

	states.printAll()
	fmt.Println("size:", states.size)
	fmt.Println("")

	state := states.get("Texas")
	if state == nil {
		fmt.Println("The state you are looking for doesn't exist")
	} else {
		fmt.Println("The state you are looking for is at ", state)
	}

	state = states.get("Colorado")
	if state == nil {
		fmt.Println("The state you are looking for doesn't exist")
	} else {
		fmt.Println("The state you are looking for is at ", state)
	}
}
