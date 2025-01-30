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

