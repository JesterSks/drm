package mode

import "github.com/JesterSks/drm/internal/mode"

func convertSysInfoToInfo(i mode.SysInfo) Info {
	return Info{
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

func convertInfoToSysInfo(i Info) mode.SysInfo {
	return mode.SysInfo{
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
