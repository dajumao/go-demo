package main

func parent(n int) int  {
	return (n-1)/2
}

func left(n int) int  {
	return 2*n+1
}

func right(n int) int  {
	return 2*n+2
}

func buildHeap(n int,data[] int)  {
	for i:=1;i<n ;i++  {
		t := i
		for t!= 0&& data[parent(t)]>data[t]  {
			temp := data[t]
			data[t] = data[parent(t)]
			data[parent(t)] = temp
			t = parent(t)
		}
	}
}

func adjust(i int,n int,data[] int)  {
	if data[i] <= data[0]{
		return
	}
	temp := data[i]
	data[i] =data[0]
	data[0] = temp
	t := 0
	for (left(t) < n && data[t] > data[left(t)])||(right(t)<n && data[t]>data[right(t)])  {
		if right(t)<n&&data[right(t)]<data[left(t)] {
			temp = data[t]
			data[t] = data[right(t)]
			data[right(t)] = temp
			t = right(t)
		} else {
			temp = data[t]
			data[t] = data[left(t)]
			data[left(t)] = temp
			t = left(t)
		}
	}
}

func main()  {

}