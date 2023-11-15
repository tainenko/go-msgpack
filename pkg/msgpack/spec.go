package msgpack

// Formats Table https://github.com/msgpack/msgpack/blob/master/spec.md#overview
// format name / first byte (in binary)
const (
	Nil byte = 0xc0 + iota
	NeverUsed
	False
	True
	Bin8
	Bin16
	Bin32
	Ext8
	Ext16
	Ext32
	Float32
	Float64
	Uint8
	Uint16
	Uint32
	Uint64
	Int8
	Int16
	Int32
	Int64
	FixExt1
	FixExt2
	FixExt4
	FixExt8
	FixExt16
	Str8
	Str16
	Str32
	Array16
	Array32
	Map16
	Map32
)
