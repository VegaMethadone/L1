package main

import "fmt"

type USB interface {
	connectWithUsbCable()
}

type MemoryCard struct {
	name string
	size int
}
type memoryCardAdapter struct {
	memoryCardAdapter *MemoryCard
}

func (m *MemoryCard) insert()   { fmt.Printf("%s is inserted\n", m.name) }
func (m *MemoryCard) copyData() { fmt.Printf("Successfully copied %d Bits\n", m.size) }

func NewMemoryCardAdapter(memoryCard *MemoryCard) *memoryCardAdapter {
	return &memoryCardAdapter{memoryCardAdapter: memoryCard}
}
func (m *memoryCardAdapter) connectWithUsbCable() {
	m.memoryCardAdapter.insert()
	m.memoryCardAdapter.copyData()
}

func main() {
	memoryCard := &MemoryCard{name: "SSD", size: 1024}
	adaprder := NewMemoryCardAdapter(memoryCard)

	var usb USB = adaprder

	usb.connectWithUsbCable()
}
