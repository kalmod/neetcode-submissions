/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    
    nodeMap := make(map[*Node]*Node)
    curr := head
    for curr != nil {
        cpy, exists := nodeMap[curr]
        if !exists { // curr node doesn't exist in map, create entry
            cpy = &Node{curr.Val, nil, nil}
            nodeMap[curr] = cpy
        }

        cpyNext, exists := nodeMap[curr.Next]
        if !exists && curr.Next != nil{
            cpyNext = &Node{curr.Next.Val, nil, nil}
            nodeMap[curr.Next] = cpyNext
        }
        cpy.Next = cpyNext

        cpyRand, exists := nodeMap[curr.Random]
        if !exists && curr.Random != nil {
            cpyRand = &Node{curr.Random.Val, nil, nil}
            nodeMap[curr.Random] = cpyRand
        }
        cpy.Random = cpyRand

        curr = curr.Next
    }

    // print check
    // fmt.Println()
    // cpyHead := nodeMap[head]
    // for cpyHead != nil {
    //     fmt.Println(cpyHead)
    //     cpyHead = cpyHead.Next
    // }

    return nodeMap[head]
}
