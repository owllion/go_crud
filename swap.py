def swapPairs(head):
    if not head: return 
    dummy = ListNode(0)
    dummy.next = head
    cur = dummy

    while cur.next:
        fn = cur.next
        sn = cur.next.next

        cur.next = sn
        cur.next.next = fn

        fn.next = sn.next

        cur = cur.next.next
        #就是fn