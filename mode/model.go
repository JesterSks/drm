package mode

import "github.com/JesterSks/drm/internal/mode"

type ModeInfo struct {
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

type Resources struct {
	mode.SysResources

	Fbs        []uint32
	Crtcs      []uint32
	Connectors []uint32
	Encoders   []uint32
}

type Connector struct {
	mode.SysGetConnector

	ID            uint32
	EncoderID     uint32
	Type          uint32
	TypeID        uint32
	Connection    uint8
	Width, Height uint32
	Subpixel      uint8

	Modes []ModeInfo

	Props      []uint32
	PropValues []uint64

	Encoders []uint32
}

type Encoder struct {
	ID   uint32
	Type uint32

	CrtcID uint32

	PossibleCrtcs  uint32
	PossibleClones uint32
}

type Crtc struct {
	ID       uint32
	BufferID uint32 // FB id to connect to 0 = disconnect

	// Position on the framebuffer
	X uint32
	Y uint32

	Width     uint32
	Height    uint32
	ModeValid int
	Mode      ModeInfo

	GammaSize int // Number of gamma stops
}

type FB struct {
	Height uint32
	Width  uint32
	BPP    uint32
	Flags  uint32
	Handle uint32
	Pitch  uint32
	Size   uint64
}
