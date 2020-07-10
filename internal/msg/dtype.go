package msg

// These functions convert the block types (VtRoot, VtDir, VtData)
// from the library usage to the wire and disk, and vice-versa.

// Please note that in the wire and disk there is no distinction
// between dir+N and data+N packets, they are all "PointerTypeN".

const (
	OVtErrType = iota // illegal
	OVtRootType
	OVtDirType
	OVtPointerType0
	OVtPointerType1
	OVtPointerType2
	OVtPointerType3
	OVtPointerType4
	OVtPointerType5
	OVtPointerType6
	OVtPointerType7 // not used
	OVtPointerType8 // not used
	OVtPointerType9 // not used
	OVtDataType
)

const (
	VtDataType = iota << 3
	VtDirType
	VtRootType
	VtCorruptType = 0xFF
)

func toDiskType(b byte) byte {
	todisk := []byte{
		OVtDataType,
		OVtPointerType0,
		OVtPointerType1,
		OVtPointerType2,
		OVtPointerType3,
		OVtPointerType4,
		OVtPointerType5,
		OVtPointerType6,
		OVtDirType,
		OVtPointerType0,
		OVtPointerType1,
		OVtPointerType2,
		OVtPointerType3,
		OVtPointerType4,
		OVtPointerType5,
		OVtPointerType6,
		OVtRootType,
	}
	if int(b) >= len(todisk) {
		return VtCorruptType
	}
	return todisk[b]
}

func fromDiskType(b byte) byte {
	fromdisk := []byte{
		VtCorruptType,
		VtRootType,
		VtDirType,
		VtDirType + 1,
		VtDirType + 2,
		VtDirType + 3,
		VtDirType + 4,
		VtDirType + 5,
		VtDirType + 6,
		VtDirType + 7,
		VtCorruptType,
		VtCorruptType,
		VtCorruptType,
		VtDataType,
	}

	if int(b) >= len(fromdisk) {
		return VtCorruptType
	}
	return fromdisk[b]
}
