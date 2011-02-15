package physfs

import(
	"os"
	"unsafe"
)

// #include <physfs.h>
import "C"

func init() {
	if !IsInit() {
		if C.PHYSFS_init(C.CString(os.Args[0])) == 0 {
			panic(GetLastError())
		}
	}
}

type File C.PHYSFS_File

type ArchiveInfo struct {
	Extension string
	Description string
	Author string
	URL string
}

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func IsInit() (bool) {
	if int(C.PHYSFS_isInit()) != 0 {
		return true
	}

	return false
}

func Deinit() (os.Error) {
	if int(C.PHYSFS_deinit()) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func GetLastError() (string) {
	return C.GoString(C.PHYSFS_getLastError())
}

func GetVersion() (ver *Version) {
	ver = new(Version)

	ver.Major = C.PHYSFS_VER_MAJOR
	ver.Minor = C.PHYSFS_VER_MINOR
	ver.Patch = C.PHYSFS_VER_PATCH

	return ver
}

func GetLinkedVersion() (ver *Version) {
	var v C.PHYSFS_Version
	C.PHYSFS_getLinkedVersion(&v)
	ver = (*Version)(unsafe.Pointer(&v))

	return ver
}

//func SupportedArchiveTypes() (ai []ArchiveInfo) {
//	cai := C.PHYSFS_supportedArchiveTypes()
//
//	i := uintptr(0)
//	for {
//		archive := *(**C.PHYSFS_ArchiveInfo)(unsafe.Pointer(uintptr(unsafe.Pointer(cai)) + i))
//		if archive == nil {
//			break
//		}
//
//		ai = append(ai, *(*ArchiveInfo)(unsafe.Pointer(archive)))
//
//		i += uintptr(unsafe.Sizeof(cai))
//	}
//
//	return ai
//}

func GetBaseDir() (string) {
	return C.GoString(C.PHYSFS_getBaseDir())
}

func GetUserDir() (string) {
	return C.GoString(C.PHYSFS_getUserDir())
}

func GetWriteDir() (string) {
	return C.GoString(C.PHYSFS_getWriteDir())
}

func SetWriteDir(dir string) (os.Error) {
	if int(C.PHYSFS_setWriteDir(C.CString(dir))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func GetDirSeparator() (string) {
	return C.GoString(C.PHYSFS_getDirSeparator())
}

func SetSaneConfig(org, app, ext string, cd, arc bool) (os.Error) {
	cdArg := 0
	if cd {
		cdArg = 1
	}

	arcArg := 0
	if arc {
		arcArg = 1
	}

	if int(C.PHYSFS_setSaneConfig(C.CString(org), C.CString(app), C.CString(ext), C.int(cdArg), C.int(arcArg))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func GetCdRomDirs() (sp []string, err os.Error) {
	csp := C.PHYSFS_getCdRomDirs()

	if csp == nil {
		return nil, os.NewError(GetLastError())
	}

	i := uintptr(0)
	for {
		p := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(csp)) + i))
		if p == nil {
			break
		}

		sp = append(sp, C.GoString(p))

		i += uintptr(unsafe.Sizeof(csp))
	}

	C.PHYSFS_freeList(unsafe.Pointer(csp))
	return sp, nil
}

func GetSearchPath() (sp []string, err os.Error) {
	csp := C.PHYSFS_getSearchPath()

	if csp == nil {
		return nil, os.NewError(GetLastError())
	}

	i := uintptr(0)
	for {
		p := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(csp)) + i))
		if p == nil {
			break
		}

		sp = append(sp, C.GoString(p))

		i += uintptr(unsafe.Sizeof(csp))
	}

	C.PHYSFS_freeList(unsafe.Pointer(csp))
	return sp, nil
}

func PermitSymbolicLinks(set bool) {
	s := C.int(0)
	if set {
		s = 1
	}

	C.PHYSFS_permitSymbolicLinks(s)
}

func SymbolicLinksPermitted() (bool) {
	if int(C.PHYSFS_symbolicLinksPermitted()) != 0 {
		return true
	}

	return false
}

func IsSymbolicLink(n string) (bool) {
	if int(C.PHYSFS_isSymbolicLink(C.CString(n))) != 0 {
		return true
	}

	return false
}

func GetRealDir(n string) (string, os.Error) {
	dir := C.PHYSFS_getRealDir(C.CString(n))

	if dir != nil {
		return C.GoString(dir), nil
	}

	return C.GoString(dir), os.NewError(GetLastError())
}

func EnumerateFiles(dir string) (list []string, err os.Error) {
	clist := C.PHYSFS_enumerateFiles(C.CString(dir))

	if clist == nil {
		return nil, os.NewError(GetLastError())
	}

	i := uintptr(0)
	for {
		item := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(clist)) + i))
		if item == nil {
			break
		}

		list = append(list, C.GoString(item))

		i += uintptr(unsafe.Sizeof(clist))
	}

	C.PHYSFS_freeList(unsafe.Pointer(clist))
	return list, nil
}

func Exists(n string) (bool) {
	if int(C.PHYSFS_exists(C.CString(n))) != 0 {
		return true
	}

	return false
}

func Delete(n string) (os.Error) {
	if int(C.PHYSFS_delete(C.CString(n))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func IsDirectory(dir string) (bool) {
	if int(C.PHYSFS_isDirectory(C.CString(dir))) != 0 {
		return true
	}

	return false
}

func Mkdir(dir string) (os.Error) {
	if int(C.PHYSFS_mkdir(C.CString(dir))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func Mount(dir, mp string, app bool) (os.Error) {
	a := 0
	if app {
		a = 1
	}

	if int(C.PHYSFS_mount(C.CString(dir), C.CString(mp), C.int(a))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func GetMountPoint(dir string) (string, os.Error) {
	mp := C.PHYSFS_getMountPoint(C.CString(dir))

	if mp != nil {
		return C.GoString(mp), nil
	}

	return C.GoString(mp), os.NewError(GetLastError())
}

func AddToSearchPath(dir string, app bool) (os.Error) {
	a := 0
	if app {
		a = 1
	}

	if int(C.PHYSFS_addToSearchPath(C.CString(dir), C.int(a))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func RemoveFromSearchPath(dir string) (os.Error) {
	if int(C.PHYSFS_removeFromSearchPath(C.CString(dir))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func GetLastModTime(n string) (int64, os.Error) {
	n := int64(C.PHYSFS_getLastModTime(C.CString(n)))

	if n != -1 {
		return n, nil
	}

	return n, os.NewError(GetLastError())
}

func Open(name string, flag int) (f *File, err os.Error) {
	switch flag {
		case os.O_RDONLY:
			f = (*File)(C.PHYSFS_openRead(C.CString(name)))
		case os.O_WRONLY:
			f = (*File)(C.PHYSFS_openWrite(C.CString(name)))
		case os.O_WRONLY | os.O_APPEND:
			f = (*File)(C.PHYSFS_openAppend(C.CString(name)))
		default:
			return nil, os.NewError("Unknown flag(s).")
	}

	if f == nil {
		return f, os.NewError(GetLastError())
	}

	return f, nil
}

func (f *File)Close() (os.Error) {
	if int(C.PHYSFS_close((*C.PHYSFS_File)(f))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func (f *File)Read(buf []byte) (n int, err os.Error) {
	n = int(C.PHYSFS_read((*C.PHYSFS_File)(f), unsafe.Pointer(&buf[0]), 1, C.PHYSFS_uint32(len(buf))))

	if n == -1 {
		return n, os.NewError(GetLastError())
	}

	return n, nil
}

func (f *File)Write(buf []byte) (n int, err os.Error) {
	n = int(C.PHYSFS_write((*C.PHYSFS_File)(f), unsafe.Pointer(&buf[0]), 1, C.PHYSFS_uint32(len(buf))))

	if n == -1 {
		return n, os.NewError(GetLastError())
	}

	return n, nil
}

func (f *File)EOF() (bool) {
	if int(C.PHYSFS_eof((*C.PHYSFS_File)(f))) != 0 {
		return true
	}

	return false
}

func (f *File)Tell() (int64, os.Error) {
	r := int64(C.PHYSFS_tell((*C.PHYSFS_File)(f)))
	if r == -1 {
		return r, os.NewError(GetLastError())
	}

	return r, nil
}

func (f *File)Seek(offset int64, none int) (int64, os.Error) {
	r := int64(C.PHYSFS_seek((*C.PHYSFS_File)(f), C.PHYSFS_uint64(offset)))

	if r == 0 {
		return r, os.NewError(GetLastError())
	}

	return r, nil
}

func (f *File)Length() (int64, os.Error) {
	r := int64(C.PHYSFS_fileLength((*C.PHYSFS_File)(f)))

	if r == -1 {
		return r, os.NewError(GetLastError())
	}

	return r, nil
}

func (f *File)SetBuffer(size uint64) (os.Error) {
	if int(C.PHYSFS_setBuffer((*C.PHYSFS_File)(f), C.PHYSFS_uint64(size))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func (f *File)Flush() (os.Error) {
	if int(C.PHYSFS_flush((*C.PHYSFS_File)(f))) != 0 {
		return nil
	}

	return os.NewError(GetLastError())
}

func (f *File)Sync() (os.Error) {
	return f.Flush()
}
