package main

import "fmt"

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

/*
   https://www.golangdev.in/2021/09/extra-long-factorials-hackerank.html
   Above question is basically simple factorial of a number but the real problem of this question is
   to how to store the output of the factorial in a int variable, because after 25th,
   the output is not going to store in int64 variable.

   Solution: For solving this question I used linked-list data structure for storing the value and apply simple math and re-create the linked list.
*/

type LL struct {
	List *longInt
}

type longInt struct {
	val  int32
	next *longInt
}

func createNode(n int32) *longInt {
	return &longInt{
		val:  n,
		next: nil,
	}
}

func (lst *LL) insertAtBeginning(data int32) {
	if nil == lst.List {
		lst.List = createNode(data)
		return
	}

	tempNode := createNode(data)
	head := lst.List
	tempNode.next = head
	lst.List = tempNode
}

func (lst *LL) deleteFromBegin() int32 {
	var val int32
	if nil != lst && nil != lst.List {
		head := lst.List
		val = head.val
		lst.List = head.next
		head = nil
	}
	return val
}

func (lst *LL) reverseList() {
	if nil != lst && nil != lst.List {
		head := lst.List
		firstPtr := head
		temp := new(longInt)
		temp = nil

		for head != nil && nil != firstPtr {
			firstPtr = head
			head = firstPtr.next
			firstPtr.next = nil
			firstPtr.next = temp
			temp = firstPtr
		}
		lst.List = temp
	}
}

func (lst *LL) readFromLast(data int32) {
	tempLst := new(LL)
	if nil != lst && nil != lst.List {
		lst.reverseList()
		var remain int32
		for nil != lst.List {
			val := lst.deleteFromBegin()
			newMultVal := (val * data) + remain
			//fmt.Println(newMultVal)
			remain = 0
			if newMultVal <= 9 {
				tempLst.insertAtBeginning(newMultVal)
				continue
			}
			if newMultVal > 9 {
				tempLst.insertAtBeginning(newMultVal % 10)
				remain = newMultVal / 10
			}
		}
		if remain > 0 {
			for remain > 0 {
				tempLst.insertAtBeginning(remain % 10)
				remain /= 10
			}
		}
	} else {
		lst.insertAtBeginning(data)
		return
	}
	lst.List = tempLst.List
}

func (lst *LL) printList() {
	if nil != lst && nil != lst.List {
		head := lst.List
		for nil != head {
			fmt.Printf("%d", head.val)
			head = head.next
		}
		fmt.Println()
	}
}

func (lst *LL) factMultiPly(val int32) {
	lst.readFromLast(val)
}

func extraLongFactorials(n int32) {
	bigLst := new(LL)
	for n > 0 {
		bigLst.factMultiPly(n)
		n--
	}
	bigLst.printList()
}

func main() {
	extraLongFactorials(25)
}
