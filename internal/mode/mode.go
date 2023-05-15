package mode

import (
	"os"
	"unsafe"

	"github.com/JesterSks/drm/internal/ioctl"
)

func GetResources(f *os.File, res *SysResources) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeResources), uintptr(unsafe.Pointer(res)))
}

func GetConnector(f *os.File, conn *SysGetConnector) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeGetConnector), uintptr(unsafe.Pointer(conn)))
}

func GetEncoder(f *os.File, en *SysGetEncoder) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeGetEncoder), uintptr(unsafe.Pointer(en)))
}

func CreateFB(f *os.File, fb *SysCreateDumb) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeCreateDumb), uintptr(unsafe.Pointer(fb)))
}

func AddFB(f *os.File, fb *SysFBCmd) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeAddFB), uintptr(unsafe.Pointer(fb)))
}

func RmFB(f *os.File, fb *SysRmFB) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeRmFB), uintptr(unsafe.Pointer(fb)))
}

func MapDumb(f *os.File, fb *SysMapDumb) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeMapDumb), uintptr(unsafe.Pointer(fb)))
}

func DestroyDumb(f *os.File, fb *SysDestroyDumb) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeDestroyDumb), uintptr(unsafe.Pointer(fb)))
}

func GetCrtc(f *os.File, crtc *SysCrtc) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeGetCrtc), uintptr(unsafe.Pointer(crtc)))
}

func SetCrtc(f *os.File, crtc *SysCrtc) error {
	return ioctl.IOCtl(f.Fd(), uintptr(ioctlModeSetCrtc), uintptr(unsafe.Pointer(crtc)))
}
