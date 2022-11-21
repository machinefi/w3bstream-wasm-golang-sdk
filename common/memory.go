package common

func NewMemory(size uint32) *Mem {
	if size == 0 {
		return &Mem{}
	}
	data := make([]byte, size)
	addr, _ := BytesToPointer(data)
	return &Mem{Size: size, Addr: addr, Data: data}
}

type Mem struct {
	Addr uint32
	Size uint32
	Data []byte
}

type MemMgr struct {
	memories  map[uint32]*Mem // memories: key is addr(uint32)
	resources map[uint32]*Mem // resources: key is resource id, be consistent with host resources
}

var Allocations = &MemMgr{
	memories:  make(map[uint32]*Mem),
	resources: make(map[uint32]*Mem),
}

//export alloc
func Alloc(size uint32) uint32 { return Allocations.AddAllocation(size).Addr }

func FreeResource(rid uint32) bool { return Allocations.FreeResource(rid) }

func (mm *MemMgr) AddAllocation(size uint32) *Mem {
	data := make([]byte, size)
	addr, _ := BytesToPointer(data)
	m := &Mem{
		Addr: addr,
		Size: size,
		Data: data,
	}
	Allocations.memories[addr] = m
	return m
}

func (mm *MemMgr) AddResource(rid, size uint32) *Mem {
	m := mm.AddAllocation(size)
	mm.resources[rid] = m
	return m
}

func (mm *MemMgr) AddResourceWithMem(rid uint32, m *Mem) {
	mm.resources[rid] = m
}

func (mm *MemMgr) GetByAddr(ptr uint32) *Mem {
	m, ok := mm.memories[ptr]
	if ok {
		return m
	}
	return nil
}

func (mm *MemMgr) GetByRID(rid uint32) *Mem {
	m, ok := mm.resources[rid]
	if ok {
		return m
	}
	return nil
}

func (mm *MemMgr) FreeResource(rid uint32) bool {
	m := Allocations.GetByRID(rid)
	delete(Allocations.resources, rid)
	if m != nil {
		delete(Allocations.memories, m.Addr)
		return true
	}
	return false
}
