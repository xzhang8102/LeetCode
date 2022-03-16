package golang

import "container/list"

/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] All O`one Data Structure
 */

// @lc code=start
type lc432Node struct {
	count int
	data  map[string]bool
}

type AllOne struct {
	*list.List
	nodeMap map[string]*list.Element
}

func lc432Constructor() AllOne {
	return AllOne{list.New(), map[string]*list.Element{}}
}

func (this *AllOne) Inc(key string) {
	if curr := this.nodeMap[key]; curr == nil {
		if this.List.Len() == 0 || this.List.Front().Value.(lc432Node).count > 1 {
			newNode := lc432Node{count: 1, data: map[string]bool{key: true}}
			this.nodeMap[key] = this.List.PushFront(newNode)
		} else {
			this.nodeMap[key] = this.List.Front()
			this.nodeMap[key].Value.(lc432Node).data[key] = true
		}
	} else {
		if next := curr.Next(); next == nil || next.Value.(lc432Node).count > curr.Value.(lc432Node).count+1 {
			newNode := lc432Node{count: curr.Value.(lc432Node).count + 1, data: map[string]bool{key: true}}
			this.nodeMap[key] = this.List.InsertAfter(newNode, curr)
		} else {
			this.nodeMap[key] = next
			next.Value.(lc432Node).data[key] = true
		}
		delete(curr.Value.(lc432Node).data, key)
		if len(curr.Value.(lc432Node).data) == 0 {
			this.List.Remove(curr)
		}
	}
}

func (this *AllOne) Dec(key string) {
	if curr := this.nodeMap[key]; curr != nil {
		currNode := curr.Value.(lc432Node)
		if currNode.count > 1 {
			if prev := curr.Prev(); prev == nil || prev.Value.(lc432Node).count < currNode.count-1 {
				this.nodeMap[key] = this.List.InsertBefore(lc432Node{count: currNode.count - 1, data: map[string]bool{key: true}}, curr)
			} else {
				this.nodeMap[key] = prev
				prev.Value.(lc432Node).data[key] = true
			}
		} else {
			delete(this.nodeMap, key)
		}
		delete(curr.Value.(lc432Node).data, key)
		if len(curr.Value.(lc432Node).data) == 0 {
			this.List.Remove(curr)
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	if back := this.List.Back(); back != nil {
		for k := range back.Value.(lc432Node).data {
			return k
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if front := this.List.Front(); front != nil {
		for k := range front.Value.(lc432Node).data {
			return k
		}
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end
