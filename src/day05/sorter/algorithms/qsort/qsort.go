package qsort

func quickSort(values []int , left, right int){
	temp := values[left]
	p := left
	i, j := left, right
	for i <= j{
		for j >= p && values[j] >= temp{
			j--
		}
		//从后往前找到比values[p]小的放在p所在位置
		if j >= p{
			values[p] = values[j]
			//更新p为j所在位置
			p = j
		}

		for i <= p && values[i] <= temp{
			i++
		}
		//从前往后找到比values[p]大的放在p所在位置
		 if i <= p{
			 values[p] = values[i]
			 //更新p为i所在位置
		 	p = i
		 }
	}
	values[p] = temp
	if p - left > 1{
		quickSort(values, left, p-1)
	}
	if right - p > 1{
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int){
	if len(values) <= 1{
		return 
	}
	quickSort(values, 0, len(values)-1)
}