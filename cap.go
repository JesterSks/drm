package drm

import (
	"os"

	"github.com/JesterSks/drm/internal"
)

const (
	CapDumbBuffer uint64 = iota + 1
	CapVBlankHighCRTC
	CapDumbPreferredDepth
	CapDumbPreferShadow
	CapPrime
	CapTimestampMonotonic
	CapAsyncPageFlip
	CapCursorWidth
	CapCursorHeight

	CapAddFB2Modifiers = 0x10
)

func HasDumbBuffer(f *os.File) bool {
	cap, err := GetCap(f, CapDumbBuffer)
	if err != nil {
		return false
	}
	return cap != 0
}

func GetCap(f *os.File, id uint64) (uint64, error) {
	c := internal.SysCapability{ID: id}

	if err := internal.GetCapability(f, &c); err != nil {
		return 0, err
	}

	return c.Val, nil
}
