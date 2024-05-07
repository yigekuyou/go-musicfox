// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureMediaPlaybackCommandManagerPreviousReceivedEventArgs string = "rc(Windows.Media.Playback.MediaPlaybackCommandManagerPreviousReceivedEventArgs;{525e3081-4632-4f76-99b1-d771623f6287})"

type MediaPlaybackCommandManagerPreviousReceivedEventArgs struct {
	ole.IUnknown
}

func (impl *MediaPlaybackCommandManagerPreviousReceivedEventArgs) GetHandled() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPreviousReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPreviousReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetHandled()
}

func (impl *MediaPlaybackCommandManagerPreviousReceivedEventArgs) SetHandled(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPreviousReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPreviousReceivedEventArgs)(unsafe.Pointer(itf))
	return v.SetHandled(value)
}

func (impl *MediaPlaybackCommandManagerPreviousReceivedEventArgs) GetDeferral() (*foundation.Deferral, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaPlaybackCommandManagerPreviousReceivedEventArgs))
	defer itf.Release()
	v := (*iMediaPlaybackCommandManagerPreviousReceivedEventArgs)(unsafe.Pointer(itf))
	return v.GetDeferral()
}

const GUIDiMediaPlaybackCommandManagerPreviousReceivedEventArgs string = "525e3081-4632-4f76-99b1-d771623f6287"
const SignatureiMediaPlaybackCommandManagerPreviousReceivedEventArgs string = "{525e3081-4632-4f76-99b1-d771623f6287}"

type iMediaPlaybackCommandManagerPreviousReceivedEventArgs struct {
	ole.IInspectable
}

type iMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl struct {
	ole.IInspectableVtbl

	GetHandled  uintptr
	SetHandled  uintptr
	GetDeferral uintptr
}

func (v *iMediaPlaybackCommandManagerPreviousReceivedEventArgs) VTable() *iMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl {
	return (*iMediaPlaybackCommandManagerPreviousReceivedEventArgsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaPlaybackCommandManagerPreviousReceivedEventArgs) GetHandled() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetHandled,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaPlaybackCommandManagerPreviousReceivedEventArgs) SetHandled(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetHandled,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaPlaybackCommandManagerPreviousReceivedEventArgs) GetDeferral() (*foundation.Deferral, error) {
	var out *foundation.Deferral
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDeferral,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.Deferral
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
