package mode

const DisplayModeLen = 32

type SysResources struct {
	FBIDPtr         uint64
	CrtcIDPtr       uint64
	ConnectorIDPtr  uint64
	EncoderIDPtr    uint64
	CountFBs        uint32
	CountCrtcs      uint32
	CountConnectors uint32
	CountEncoders   uint32
	MinWidth        uint32
	MaxWidth        uint32
	MinHeight       uint32
	MaxHeight       uint32
}

type SysGetConnector struct {
	EncodersPtr   uint64
	ModesPtr      uint64
	PropsPtr      uint64
	PropValuesPtr uint64

	CountModes    uint32
	CountProps    uint32
	CountEncoders uint32

	EncoderID       uint32 // current encoder
	ID              uint32
	ConnectorType   uint32
	ConnectorTypeID uint32

	Connection uint32

	// HxW in millimeters
	MMWidth  uint32
	MMHeight uint32
	Subpixel uint32
}

type SysGetEncoder struct {
	ID   uint32
	Type uint32

	CrtcID uint32

	PossibleCrtcs  uint32
	PossibleClones uint32
}

type SysCreateDumb struct {
	Height uint32
	Width  uint32
	BPP    uint32
	Flags  uint32

	// returned values
	Handle uint32
	Pitch  uint32
	Size   uint64
}

type SysMapDumb struct {
	Handle uint32 // Handle for the object being mapped
	Pad    uint32

	// Fake offset to use for subsequent mmap call
	// This is a fixed-size type for 32/64 compatibility.
	Offset uint64
}

type SysFBCmd struct {
	FBID   uint32
	Width  uint32
	Height uint32
	Pitch  uint32
	BPP    uint32
	Depth  uint32

	/* driver specific handle */
	Handle uint32
}

type SysRmFB struct {
	Handle uint32
}

type SysInfo struct {
	Clock uint32

	Hdisplay   uint16
	HsyncStart uint16
	HsyncEnd   uint16
	Htotal     uint16
	Hskew      uint16

	Vdisplay   uint16
	VsyncStart uint16
	VsyncEnd   uint16
	Vtotal     uint16
	Vscan      uint16

	Vrefresh uint32

	Flags uint32
	Type  uint32
	Name  [DisplayModeLen]uint8
}

type SysCrtc struct {
	SetConnectorsPtr uint64
	CountConnectors  uint32

	ID   uint32
	FBID uint32 // Id of framebuffer

	// Position on the frameuffer
	X uint32
	Y uint32

	GammaSize uint32
	ModeValid uint32
	Mode      SysInfo
}

type SysDestroyDumb struct {
	Handle uint32
}
