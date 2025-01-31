package llist

// Reversed linked list pointed by this head
// 1->2->3->4
func ReverseLinkedListHelper(head *llnode) *llnode {
    if head==nil || head.next == nil {
        return head
    }
    reversedllhead := ReverseLinkedListHelper(head.next)
    head.next.next=head
    head.next = nil
    return reversedllhead
}

func ReverseLinkedList(ll *Linkedlist) *Linkedlist {
    ll.head = ReverseLinkedListHelper(ll.head)
    return ll
}

// returns pointer to node where two linked lists intersects
func LinkedListIntersection(listA, listB *Linkedlist) *llnode {
    ptrA, ptrB := listA.head, listB.head
    for ptrA!=ptrB {
        // traverse A
        if ptrA!=nil {
            ptrA = ptrA.next
        } else {
            ptrA = listB.head
        }

        // traverseB
        if ptrB!=nil {
            ptrB = ptrB.next
        } else {
            ptrB = listA.head
        }
    }

    return ptrA
}