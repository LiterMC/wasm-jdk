package sun_nio_fs

import (
	"os"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.init()I", UnixNativeDispatcher_init)
	native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getcwd()[B", UnixNativeDispatcher_getcwd)
}

// static native byte[] getcwd();
func UnixNativeDispatcher_getcwd(vm ir.VM) error {
	stack := vm.GetStack()
	path, err := os.Getwd()
	if err != nil {
		stack.PushRef(nil)
		return nil
	}
	pathRef := vm.NewArray(desc.DescByteArray, (int32)(len(path)))
	copy(pathRef.GetByteArr(), path)
	stack.PushRef(pathRef)
	return nil
}

// static native int dup(int filedes) throws UnixException;
// private static native int open0(long pathAddress, int flags, int mode) throws UnixException;
// private static native int openat0(int dfd, long pathAddress, int flags, int mode) throws UnixException;
// private static native void close0(int fd) throws UnixException;
// static native void rewind(long stream) throws UnixException;
// static native int getlinelen(long stream) throws UnixException;
// private static native void link0(long existingAddress, long newAddress) throws UnixException;
// private static native void unlink0(long pathAddress) throws UnixException;
// private static native void unlinkat0(int dfd, long pathAddress, int flag) throws UnixException;
// private static native void mknod0(long pathAddress, int mode, long dev) throws UnixException;
// private static native void rename0(long fromAddress, long toAddress) throws UnixException;
// private static native void renameat0(int fromfd, long fromAddress, int tofd, long toAddress) throws UnixException;
// private static native void mkdir0(long pathAddress, int mode) throws UnixException;
// private static native void rmdir0(long pathAddress) throws UnixException;
// private static native byte[] readlink0(long pathAddress) throws UnixException;
// private static native byte[] realpath0(long pathAddress) throws UnixException;
// private static native void symlink0(long name1, long name2) throws UnixException;
// private static native int stat0(long pathAddress, UnixFileAttributes attrs);
// private static native void lstat0(long pathAddress, UnixFileAttributes attrs) throws UnixException;
// private static native void fstat0(int fd, UnixFileAttributes attrs) throws UnixException;
// private static native void fstatat0(int dfd, long pathAddress, int flag, UnixFileAttributes attrs) throws UnixException;
// private static native void chown0(long pathAddress, int uid, int gid) throws UnixException;
// private static native void lchown0(long pathAddress, int uid, int gid) throws UnixException;
// static native void fchown0(int fd, int uid, int gid) throws UnixException;
// private static native void chmod0(long pathAddress, int mode) throws UnixException;
// private static native void fchmod0(int fd, int mode) throws UnixException;
// private static native void utimes0(long pathAddress, long times0, long times1) throws UnixException;
// private static native void futimes0(int fd, long times0, long times1) throws UnixException;
// private static native void futimens0(int fd, long times0, long times1) throws UnixException;
// private static native void lutimes0(long pathAddress, long times0, long times1) throws UnixException;
// private static native long opendir0(long pathAddress) throws UnixException;
// static native long fdopendir(int dfd) throws UnixException;
// static native void closedir(long dir) throws UnixException;
// static native byte[] readdir0(long dir) throws UnixException;
// private static native int read0(int fildes, long buf, int nbyte) throws UnixException;
// private static native int write0(int fildes, long buf, int nbyte) throws UnixException;
// private static native int access0(long pathAddress, int amode);
// static native byte[] getpwuid(int uid) throws UnixException;
// static native byte[] getgrgid(int gid) throws UnixException;
// private static native int getpwnam0(long nameAddress) throws UnixException;
// private static native int getgrnam0(long nameAddress) throws UnixException;
// private static native void statvfs0(long pathAddress, UnixFileStoreAttributes attrs) throws UnixException;
// static native byte[] strerror(int errnum);
// private static native int fgetxattr0(int filedes, long nameAddress, long valueAddress, int valueLen) throws UnixException;
// private static native void fsetxattr0(int filedes, long nameAddress, long valueAddress, int valueLen) throws UnixException;
// private static native void fremovexattr0(int filedes, long nameAddress) throws UnixException;
// static native int flistxattr(int filedes, long listAddress, int size) throws UnixException;
// private static native int init();
func UnixNativeDispatcher_init(vm ir.VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}
