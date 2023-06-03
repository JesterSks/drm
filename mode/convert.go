package mode

import "github.com/JesterSks/drm/internal/mode"

func convertSysInfoToInfo(i mode.SysModeInfo) ModeInfo {
	return ModeInfo{
		Clock: i.Clock,

		Hdisplay:   i.Hdisplay,
		HsyncStart: i.HsyncStart,
		HsyncEnd:   i.HsyncEnd,
		Htotal:     i.Htotal,
		Hskew:      i.Hskew,

		Vdisplay:   i.Vdisplay,
		VsyncStart: i.VsyncStart,
		VsyncEnd:   i.VsyncEnd,
		Vtotal:     i.Vtotal,
		Vscan:      i.Vscan,
		Vrefresh:   i.Vrefresh,

		Flags: i.Flags,
		Type:  i.Type,
		Name:  i.Name,
	}
}

func convertInfoToSysInfo(i ModeInfo) mode.SysModeInfo {
	return mode.SysModeInfo{
		Clock: i.Clock,

		Hdisplay:   i.Hdisplay,
		HsyncStart: i.HsyncStart,
		HsyncEnd:   i.HsyncEnd,
		Htotal:     i.Htotal,
		Hskew:      i.Hskew,

		Vdisplay:   i.Vdisplay,
		VsyncStart: i.VsyncStart,
		VsyncEnd:   i.VsyncEnd,
		Vtotal:     i.Vtotal,
		Vscan:      i.Vscan,
		Vrefresh:   i.Vrefresh,

		Flags: i.Flags,
		Type:  i.Type,
		Name:  i.Name,
	}
}
