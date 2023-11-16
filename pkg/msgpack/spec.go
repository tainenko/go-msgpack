package msgpack

// Formats Table https://github.com/msgpack/msgpack/blob/master/spec.md#overview
// format name / first byte (in binary)
const (
	NilFormat byte = 0xc0 + iota
	NeverUsedFormat
	FalseFormat
	TrueFormat
	Bin8Format
	Bin16Format
	Bin32Format
	Ext8Format
	Ext16Format
	Ext32Format
	Float32Format
	Float64Format
	Uint8Format
	Uint16Format
	Uint32Format
	Uint64Format
	Int8Format
	Int16Format
	Int32Format
	Int64Format
	FixExt1Format
	FixExt2Format
	FixExt4Format
	FixExt8Format
	FixExt16Format
	Str8Format
	Str16Format
	Str32Format
	Array16Format
	Array32Format
	Map16Format
	Map32Format
)
