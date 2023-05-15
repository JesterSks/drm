package drm

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/JesterSks/drm/internal"
)

// Version of DRM driver
type Version struct {
	internal.SysVersion

	Major int32
	Minor int32
	Patch int32
	Name  string // Name of the driver (eg.: i915)
	Date  string
	Desc  string
}

const (
	driPath = "/dev/dri"
)

func Available() (Version, error) {
	f, err := OpenCard(0)
	if err != nil {
		// handle backward linux compat?
		// check /proc/dri/0 ?
		return Version{}, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	return GetVersion(f)
}

func OpenCard(n int) (*os.File, error) {
	return open(fmt.Sprintf("%s/card%d", driPath, n))
}

func OpenControlDev(n int) (*os.File, error) {
	return open(fmt.Sprintf("%s/controlD%d", driPath, n))
}

func OpenRenderDev(n int) (*os.File, error) {
	return open(fmt.Sprintf("%s/renderD%d", driPath, n))
}

func open(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR, 0)
}

func GetVersion(f *os.File) (Version, error) {
	var (
		name, date, desc []byte
	)

	version := internal.SysVersion{}

	if err := internal.GetVersion(f, &version); err != nil {
		return Version{}, err
	}

	if version.NameLen > 0 {
		name = make([]byte, version.NameLen+1)
		version.Name = uintptr(unsafe.Pointer(&name[0]))
	}

	if version.DateLen > 0 {
		date = make([]byte, version.DateLen+1)
		version.Date = uintptr(unsafe.Pointer(&date[0]))
	}
	if version.DescLen > 0 {
		desc = make([]byte, version.DescLen+1)
		version.Desc = uintptr(unsafe.Pointer(&desc[0]))
	}

	if err := internal.GetVersion(f, &version); err != nil {
		return Version{}, err
	}

	// remove C null byte at end
	name = name[:version.NameLen]
	date = date[:version.DateLen]
	desc = desc[:version.DescLen]

	nozero := func(r rune) bool {
		if r == 0 {
			return true
		} else {
			return false
		}
	}

	v := Version{
		SysVersion: version,
		Major:      version.Major,
		Minor:      version.Minor,
		Patch:      version.Patch,

		Name: string(bytes.TrimFunc(name, nozero)),
		Date: string(bytes.TrimFunc(date, nozero)),
		Desc: string(bytes.TrimFunc(desc, nozero)),
	}

	return v, nil
}

func ListDevices() []Version {
	var devices []Version
	files, err := os.ReadDir(driPath)
	if err != nil {
		return devices
	}

	for _, finfo := range files {
		name := finfo.Name()
		if len(name) <= 4 ||
			!strings.HasPrefix(name, "card") {
			continue
		}

		idstr := name[4:]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			continue
		}

		devfile, err := OpenCard(id)
		if err != nil {
			continue
		}
		dev, err := GetVersion(devfile)
		if err != nil {
			continue
		}
		devices = append(devices, dev)
	}

	return devices
}
