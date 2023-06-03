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
		modes           []ModeInfo
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

	modes = make([]ModeInfo, conn.CountModes)
	conn.ModesPtr = uint64(uintptr(unsafe.Pointer(&modes[0])))

	if conn.CountEncoders > 0 {
		encoders = make([]uint32, conn.CountEncoders)
		conn.EncodersPtr = uint64(uintptr(unsafe.Pointer(&encoders[0])))
	}

	if err := mode.GetConnector(f, &conn); err != nil {
		return nil, err
	}

	return &Connector{
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

		Props:      props,
		PropValues: propValues,
		Modes:      modes,
		Encoders:   encoders,
	}, nil
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

	ret := Crtc{
		ID:        crtc.ID,
		X:         crtc.X,
		Y:         crtc.Y,
		ModeValid: int(crtc.ModeValid),
		BufferID:  crtc.FBID,
		GammaSize: int(crtc.GammaSize),
	}

	ret.Mode = convertSysInfoToInfo(crtc.Mode)
	ret.Width = uint32(crtc.Mode.Hdisplay)
	ret.Height = uint32(crtc.Mode.Vdisplay)
	return &ret, nil
}

func SetCrtc(f *os.File, crtcid, bufferid, x, y uint32, connectors *uint32, count int, modeInfo *ModeInfo) error {
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
		crtc.Mode = convertInfoToSysInfo(*modeInfo)
		crtc.ModeValid = 1
	}

	return mode.SetCrtc(f, &crtc)
}
