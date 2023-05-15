package mode

import "github.com/JesterSks/drm/internal/ioctl"

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
	ioctlModeMapDumb = ioctl.DRMiowr[SysMapDumb](0xB3)

	// DRM_IOWR(0xB4, struct drm_mode_destroy_dumb)
	ioctlModeDestroyDumb = ioctl.DRMiowr[SysDestroyDumb](0xB4)
)
