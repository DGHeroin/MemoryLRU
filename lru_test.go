package MemoryLRU

import "testing"

func TestItemLimit(t *testing.T) {
	l := New(1, 20)

	l.OnEvicted = func(key Key, value []byte, fullType RemoveReason) {
		t.Log("LRU full, key:", key, " remove type:", fullType)
	}

	data1 := []byte("hello1")
	data2 := []byte("hello2")
	data3 := []byte("hello3")
	l.Add("key1", data1)
	l.Add("key2", data2)
	l.Add("key3", data3)
}

func TestMemoryLimit(t *testing.T) {
	l := New(3, 12)

	l.OnEvicted = func(key Key, value []byte, fullType RemoveReason) {
		t.Log("LRU full, key:", key, " remove type:", fullType)
	}

	data1 := []byte("hello1")
	data2 := []byte("hello2")
	data3 := []byte("hello3")
	l.Add("key1", data1)
	l.Add("key2", data2)
	l.Add("key3", data3)
}

func TestNoLimit(t *testing.T) {
	l := New(0, 0)

	l.OnEvicted = func(key Key, value []byte, fullType RemoveReason) {
		t.Log("LRU full, key:", key, " remove type:", fullType)
	}

	data1 := []byte("hello1")
	data2 := []byte("hello2")
	data3 := []byte("hello3")
	l.Add("key1", data1)
	l.Add("key2", data2)
	l.Add("key3", data3)
	l.RemoveOldest()
}