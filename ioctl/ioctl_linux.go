package ioctl

import (
	"syscall"
	"unsafe"
)

// To decode a hex IOCTL code:
//
// Most architectures use this generic format, but check
// include/ARCH/ioctl.h for specifics, e.g. powerpc
// uses 3 bits to encode read/write and 13 bits for size.
//
//  bits    meaning
//  31-30	00 - no parameters: uses _IO macro
// 	10 - read: _IOR
// 	01 - write: _IOW
// 	11 - read/write: _IOWR
//
//  29-16	size of arguments
//
//  15-8	ascii character supposedly
// 	unique to each driver
//
//  7-0	function #
//
// So for example 0x82187201 is a read with arg length of 0x218,
// character 'r' function 1. Grepping the source reveals this is:
//
// #define VFAT_IOCTL_READDIR_BOTH         _IOR('r', 1, struct dirent [2])
// source: https://www.kernel.org/doc/Documentation/ioctl/ioctl-decoding.txt

type iocDirection uint8

const (
	drmBase = 'd'

	fnBits        = 8
	typeBits      = 8
	sizeBits      = 14
	directionBits = 2

	fnShift        = 0
	typeShift      = fnShift + fnBits
	sizeShift      = typeShift + typeBits
	directionShift = sizeShift + sizeBits

	iocNone  iocDirection = 0x0
	iocWrite iocDirection = 0x1
	iocRead  iocDirection = 0x2
)

func DRMiowr[T any](fn uint8) uint32 {
	return ioIOWR[T](drmBase, fn)
}

func ioIOWR[T any](iocType, fn uint8) uint32 {
	var t T
	return iocCode(iocRead|iocWrite, iocType, fn, uint16(unsafe.Sizeof(t)))
}

func iocCode(direction iocDirection, iocType, fn uint8, size uint16) uint32 {
	return uint32(direction)<<directionShift |
		uint32(size)<<sizeShift |
		uint32(iocType)<<typeShift |
		uint32(fn)<<fnShift
}

func IOCtl(fd, cmd, ptr uintptr) error {
	_, _, errcode := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, ptr)
	if errcode != 0 {
		return errcode
	}
	return nil
}
