package internal

import (
	"os"
	"unsafe"

	"github.com/JesterSks/drm/internal/ioctl"
)

func GetVersion(f *os.File, v *SysVersion) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlVersion), uintptr(unsafe.Pointer(v)))
}

func GetCapability(f *os.File, c *SysCapability) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlGetCap), uintptr(unsafe.Pointer(c)))
}
