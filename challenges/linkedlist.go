package challenges

import "fmt"

type Element struct {
	Value string
	Next  *Element
}

type SinglyLinkedList struct {
	Count int
	Head  *Element
	Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	element := Element{
		Value: value,
		Next:  nil,
	}
	if l.Count == 0 {
		l.Head = &element
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &element
	}
	l.Tail = &element
	l.Count++
}

// You will have to ensure when you add new elements
// that this method still returns the correct value
func (l *SinglyLinkedList) Size() int {
	return l.Count
}

func (l *SinglyLinkedList) Print() {
	current := l.Head
	for current.Next != nil {
		fmt.Printf("%+v\n", current.Value)
		current = current.Next
	}
	fmt.Printf("%+v\n", l.Tail.Value)
}

func SinglyLinkedListMain() {
	fmt.Println("Singly Linked List Challenge")

	var llist SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		llist.Append(value)
	}
	llist.Print()
}
