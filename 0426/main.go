package main

type SnapshotArray struct {
	vList [][2]int // value 和 snapId
	changeList [][]int // idx 发生变化的所有版本号
	vcList map[int]map[int]int // 版本值
	snapId int // 快照编号
}


func Constructor(length int) SnapshotArray {
	res := SnapshotArray{
		make([][2]int, length),
		make([][]int, length),
		make(map[int]map[int]int, 0),
		-1,
	}
	return res
}


func (this *SnapshotArray) Set(index int, val int)  {
	this.vList[index] = [2]int{val, this.snapId+1}
	this.changeList[index] = append(this.changeList[index], this.snapId+1)
	if len(this.vcList[index]) == 0 {
		this.vcList[index] = make(map[int]int, 0)
	}
	this.vcList[index][this.snapId+1] = val
}


func (this *SnapshotArray) Snap() int {
	this.snapId++
	return this.snapId
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
	// 如果快照版本大于当前值的版本直接返回
	if snap_id >= this.vList[index][1] {
		return this.vList[index][0]
	} else {
		// 查找
		l := 0
		r := len(this.changeList[index])
		find := -1
		for l<r {
			mid := (l+r)/2
			if this.changeList[index][mid] == snap_id {
				find = mid
				break
			}
			if this.changeList[index][mid] < snap_id {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if find == -1 {
			find = l-1
		}
		if find < 0 {
			return 0
		}
		if len(this.vcList[index]) == 0 {
			return this.vList[index][1]
		} else {
			return this.vcList[index][this.changeList[index][find]]
		}
	}
}

func main() {
	// s := Constructor(3)
	// s.Set(1, 18)
	// s.Set(1, 4)
	// s.Snap()
	// s.Get(0, 0)
	// s.Set(0, 20)
	// s.Snap()
	// s.Set(0, 2)
	// s.Set(1, 1)
	// s.Get(1, 1)
	// s.Get(1, 0)

	s := Constructor(2)
	s.Snap()
	s.Get(1,0)
	s.Get(0,0)
	s.Set(1,8)
	s.Get(1,0)
	s.Set(0,20)
	s.Get(0,0)
	s.Set(0,7)
	
	println(123)
}