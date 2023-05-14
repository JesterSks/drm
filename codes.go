package drm

import (
	"github.com/JesterSks/drm/internal/ioctl"
)

var (
	// DRM_IOWR(0x00, struct drm_version)
	IOCTLVersion = ioctl.DRMiowr[version](0x00)

	// DRM_IOWR(0x0c, struct drm_get_cap)
	IOCTLGetCap = ioctl.DRMiowr[capability](0x0C)
)
