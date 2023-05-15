package internal

import (
	"github.com/JesterSks/drm/internal/ioctl"
)

var (
	// DRM_IOWR(0x00, struct drm_version)
	ioctlVersion = ioctl.DRMiowr[SysVersion](0x00)

	// DRM_IOWR(0x0c, struct drm_get_cap)
	ioctlGetCap = ioctl.DRMiowr[SysCapability](0x0C)
)
