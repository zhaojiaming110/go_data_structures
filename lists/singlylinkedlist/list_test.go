package singlylinkedlist

import (
	"github.com/zhaojiaming110/go_data_structures/utils"
	"testing"
)

func TestNew(t *testing.T) {
	list1 := New()
	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New(1, "b")
	t.Log(list2.Values())

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}

	if actualValue, ok := list2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestList_Add(t *testing.T) {
	list := New()
	list.Add(1, 2, 3)
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestList_Remove(t *testing.T) {
	list := New(1, 2, 3)
	t.Log(list.Remove(2))
	t.Log(list.Values())
	t.Log(list.Size())
	list.Add(3, 4)
	t.Log(list.Size())
	for i := 0; i < list.size; i++ {
		t.Log(list.Get(i))
	}
	t.Log(list.Values())
}

func TestList_Insert(t *testing.T) {
	list1 := New(1, 2, 3)
	list1.Insert(1, 4,5,6)
	t.Log(list1.Values())

	list2 := New(1, 2, 3)
	list2.Insert(2, 4,5,6)
	t.Log(list2.Values())

	list3 := New(1,2,3)
	list3.Insert(3,4,5,6)
	t.Log(list3.Values())

	list4 := New(1, 2, 3)
	list4.Insert(0,4,5,6)
	t.Log(list4.Values())
}

func TestList_Contains(t *testing.T) {
	list := New(1, 2, 3)
	t.Log(list.Contains(2, 3, 4))
}

func TestList_Swap(t *testing.T) {
	list := New(1, 2, 3)
	t.Log(list.Values())
	list.Swap(1,2)
	t.Log(list.Values())
}

func TestList_Set(t *testing.T) {
	list := New(1,2,3)
	list.Set(3, 4)
	t.Log(list.Values())
	list.Set(1,9)
	t.Log(list.Values())
}

func TestList_Sort(t *testing.T) {
	list := New(9, 5, 3, 8, 7)
	t.Log(list.Values())
	list.Sort(utils.IntComparator)
	t.Log(list.Values())
}

func TestList_Disordered(t *testing.T) {
	list1 := New(1,3,5,2,8)
	t.Log(list1.Disordered(utils.IntComparator))
	list2 := New(1,2,2,4,4)
	t.Log(list2.Disordered(utils.IntComparator))
	list3 := New(1,2)
	t.Log(list3.Disordered(utils.IntComparator))
}