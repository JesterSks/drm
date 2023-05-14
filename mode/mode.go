package mode

import (
	"os"
	"unsafe"

	"github.com/JesterSks/drm/internal/mode"
)

const (
	DisplayInfoLen   = 32
	ConnectorNameLen = 32
	DisplayModeLen   = 32
	PropNameLen      = 32

	Connected         = 1
	Disconnected      = 2
	UnknownConnection = 3
)

type (
	Info struct {
		Clock                                         uint32
		Hdisplay, HsyncStart, HsyncEnd, Htotal, Hskew uint16
		Vdisplay, VsyncStart, VsyncEnd, Vtotal, Vscan uint16

		Vrefresh uint32

		Flags uint32
		Type  uint32
		Name  [DisplayModeLen]uint8
	}

	Resources struct {
		mode.SysResources

		Fbs        []uint32
		Crtcs      []uint32
		Connectors []uint32
		Encoders   []uint32
	}

	Connector struct {
		mode.SysGetConnector

		ID            uint32
		EncoderID     uint32
		Type          uint32
		TypeID        uint32
		Connection    uint8
		Width, Height uint32
		Subpixel      uint8

		Modes []Info

		Props      []uint32
		PropValues []uint64

		Encoders []uint32
	}

	Encoder struct {
		ID   uint32
		Type uint32

		CrtcID uint32

		PossibleCrtcs  uint32
		PossibleClones uint32
	}

	Crtc struct {
		ID       uint32
		BufferID uint32 // FB id to connect to 0 = disconnect

		X, Y          uint32 // Position on the framebuffer
		Width, Height uint32
		ModeValid     int
		Mode          Info

		GammaSize int // Number of gamma stops
	}

	FB struct {
		Height, Width, BPP, Flags uint32
		Handle                    uint32
		Pitch                     uint32
		Size                      uint64
	}
)

func GetResources(f *os.File) (*Resources, error) {
	mres := mode.SysResources{}

	if err := mode.GetResources(f, &mres); err != nil {
		return nil, err
	}

	var (
		fbids, crtcids, connectorids, encoderids []uint32
	)

	if mres.CountFBs > 0 {
		fbids = make([]uint32, mres.CountFBs)
		mres.FBIDPtr = uint64(uintptr(unsafe.Pointer(&fbids[0])))
	}
	if mres.CountCrtcs > 0 {
		crtcids = make([]uint32, mres.CountCrtcs)
		mres.CrtcIDPtr = uint64(uintptr(unsafe.Pointer(&crtcids[0])))
	}
	if mres.CountEncoders > 0 {
		encoderids = make([]uint32, mres.CountEncoders)
		mres.EncoderIDPtr = uint64(uintptr(unsafe.Pointer(&encoderids[0])))
	}
	if mres.CountConnectors > 0 {
		connectorids = make([]uint32, mres.CountConnectors)
		mres.ConnectorIDPtr = uint64(uintptr(unsafe.Pointer(&connectorids[0])))
	}

	if err := mode.GetResources(f, &mres); err != nil {
		return nil, err
	}

	// TODO(i4k): handle hotplugging in-between the ioctls above

	return &Resources{
		SysResources: mres,
		Fbs:          fbids,
		Crtcs:        crtcids,
		Encoders:     encoderids,
		Connectors:   connectorids,
	}, nil
}

func GetConnector(f *os.File, id uint32) (*Connector, error) {
	conn := mode.SysGetConnector{ID: id}

	if err := mode.GetConnector(f, &conn); err != nil {
		return nil, err
	}

	var (
		props, encoders []uint32
		propValues      []uint64
		modes           []Info
	)

	if conn.CountProps > 0 {
		props = make([]uint32, conn.CountProps)
		conn.PropsPtr = uint64(uintptr(unsafe.Pointer(&props[0])))

		propValues = make([]uint64, conn.CountProps)
		conn.PropValuesPtr = uint64(uintptr(unsafe.Pointer(&propValues[0])))
	}

	if conn.CountModes == 0 {
		conn.CountModes = 1
	}

	modes = make([]Info, conn.CountModes)
	conn.ModesPtr = uint64(uintptr(unsafe.Pointer(&modes[0])))

	if conn.CountEncoders > 0 {
		encoders = make([]uint32, conn.CountEncoders)
		conn.EncodersPtr = uint64(uintptr(unsafe.Pointer(&encoders[0])))
	}

	if err := mode.GetConnector(f, &conn); err != nil {
		return nil, err
	}

	ret := &Connector{
		SysGetConnector: conn,
		ID:              conn.ID,
		EncoderID:       conn.EncoderID,
		Connection:      uint8(conn.Connection),
		Width:           conn.MMWidth,
		Height:          conn.MMHeight,

		// convert subpixel from kernel to userspace */
		Subpixel: uint8(conn.Subpixel + 1),
		Type:     conn.ConnectorType,
		TypeID:   conn.ConnectorTypeID,
	}

	ret.Props = make([]uint32, len(props))
	copy(ret.Props, props)
	ret.PropValues = make([]uint64, len(propValues))
	copy(ret.PropValues, propValues)
	ret.Modes = make([]Info, len(modes))
	copy(ret.Modes, modes)
	ret.Encoders = make([]uint32, len(encoders))
	copy(ret.Encoders, encoders)

	return ret, nil
}

func GetEncoder(f *os.File, id uint32) (*Encoder, error) {
	encoder := mode.SysGetEncoder{ID: id}

	if err := mode.GetEncoder(f, &encoder); err != nil {
		return nil, err
	}

	return &Encoder{
		ID:             encoder.ID,
		CrtcID:         encoder.CrtcID,
		Type:           encoder.Type,
		PossibleCrtcs:  encoder.PossibleCrtcs,
		PossibleClones: encoder.PossibleClones,
	}, nil
}

func CreateFB(f *os.File, width, height uint16, bpp uint32) (*FB, error) {
	fb := mode.SysCreateDumb{
		Width:  uint32(width),
		Height: uint32(height),
		BPP:    bpp,
	}

	if err := mode.CreateFB(f, &fb); err != nil {
		return nil, err
	}

	return &FB{
		Height: fb.Height,
		Width:  fb.Width,
		BPP:    fb.BPP,
		Handle: fb.Handle,
		Pitch:  fb.Pitch,
		Size:   fb.Size,
	}, nil
}

func AddFB(f *os.File, width, height uint16,
	depth, bpp uint8, pitch, boHandle uint32) (uint32, error) {
	fb := mode.SysFBCmd{
		Width:  uint32(width),
		Height: uint32(height),
		Pitch:  pitch,
		BPP:    uint32(bpp),
		Depth:  uint32(depth),
		Handle: boHandle,
	}

	if err := mode.AddFB(f, &fb); err != nil {
		return 0, err
	}

	return fb.FBID, nil
}

func RmFB(f *os.File, handle uint32) error {
	return mode.RmFB(f, &mode.SysRmFB{Handle: handle})
}

func MapDumb(f *os.File, handle uint32) (uint64, error) {
	mreq := mode.SysMapDumb{Handle: handle}

	if err := mode.MapDumb(f, &mreq); err != nil {
		return 0, err
	}

	return mreq.Offset, nil
}

func DestroyDumb(f *os.File, handle uint32) error {
	return mode.DestroyDumb(f, &mode.SysDestroyDumb{Handle: handle})
}

func GetCrtc(f *os.File, id uint32) (*Crtc, error) {
	crtc := mode.SysCrtc{ID: id}

	if err := mode.GetCrtc(f, &crtc); err != nil {
		return nil, err
	}

	ret := &Crtc{
		ID:        crtc.ID,
		X:         crtc.X,
		Y:         crtc.Y,
		ModeValid: int(crtc.ModeValid),
		BufferID:  crtc.FBID,
		GammaSize: int(crtc.GammaSize),
	}

	ret.Mode = crtc.Mode
	ret.Width = uint32(crtc.Mode.Hdisplay)
	ret.Height = uint32(crtc.Mode.Vdisplay)
	return ret, nil
}

func SetCrtc(f *os.File, crtcid, bufferid, x, y uint32, connectors *uint32, count int, modeInfo *Info) error {
	crtc := mode.SysCrtc{
		X:               x,
		Y:               y,
		ID:              crtcid,
		FBID:            bufferid,
		CountConnectors: uint32(count),
	}

	if connectors != nil {
		crtc.SetConnectorsPtr = uint64(uintptr(unsafe.Pointer(connectors)))
	}

	if modeInfo != nil {
		crtc.Mode = *modeInfo
		crtc.ModeValid = 1
	}

	return mode.SetCrtc(f, &crtc)
}
