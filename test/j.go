package main

type Element int
type Node struct {
	data Element
	next *Node
}

func main()  {
	g := &Node{
		data:1,
		next:&Node{
			data:8,
			next:&Node{
				data:5,
				next:&Node{
					
				},
			},
		},
	}
}
