package mode

import (
	"os"
	"unsafe"

	"github.com/JesterSks/drm/internal/ioctl"
)

var (
	// DRM_IOWR(0xA0, struct drm_mode_card_res)
	ioctlModeResources = ioctl.DRMiowr[SysResources](0xA0)

	// DRM_IOWR(0xA1, struct drm_mode_crtc)
	ioctlModeGetCrtc = ioctl.DRMiowr[SysCrtc](0xA1)

	// DRM_IOWR(0xA2, struct drm_mode_crtc)
	ioctlModeSetCrtc = ioctl.DRMiowr[SysCrtc](0xA2)

	// DRM_IOWR(0xA6, struct drm_mode_get_encoder)
	ioctlModeGetEncoder = ioctl.DRMiowr[SysGetEncoder](0xA6)

	// DRM_IOWR(0xA7, struct drm_mode_get_connector)
	ioctlModeGetConnector = ioctl.DRMiowr[SysGetConnector](0xA7)

	// DRM_IOWR(0xAE, struct drm_mode_fb_cmd)
	ioctlModeAddFB = ioctl.DRMiowr[SysFBCmd](0xAE)

	// DRM_IOWR(0xAF, unsigned int)
	ioctlModeRmFB = ioctl.DRMiowr[uint32](0xAF)

	// DRM_IOWR(0xB2, struct drm_mode_create_dumb)
	ioctlModeCreateDumb = ioctl.DRMiowr[SysCreateDumb](0xB2)

	// DRM_IOWR(0xB3, struct drm_mode_map_dumb)
	ioctlModeMapDumb = ioctl.DRMiowr[MapDumb](0xB3)

	// DRM_IOWR(0xB4, struct drm_mode_destroy_dumb)
	ioctlModeDestroyDumb = ioctl.DRMiowr[DestroyDumb](0xB4)
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
