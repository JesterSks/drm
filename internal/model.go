package internal

type SysCapability struct {
	ID  uint64
	Val uint64
}

type SysVersion struct {
	Major   int32
	Minor   int32
	Patch   int32
	NameLen int64
	Name    uintptr
	DateLen int64
	Date    uintptr
	DescLen int64
	Desc    uintptr
}
