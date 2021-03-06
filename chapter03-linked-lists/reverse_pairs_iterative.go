// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

func reversePairs(head *ListNode) *ListNode {
    if head == nil || head.next == nil {
        return head
    }
    result := head.next
    var previousNode *ListNode
    for head != nil && head.next != nil {
        nextNode := head.next
        head.next = nextNode.next
        nextNode.next = head
        if previousNode != nil {
            previousNode.next = nextNode
        }
        previousNode = head
        head = head.next
    }
    return result
}
