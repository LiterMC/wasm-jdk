package sun_nio_fs

import (
	"os"

	"github.com/LiterMC/wasm-jdk/cutil"
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getcwd()[B", UnixNativeDispatcher_getcwd)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.dup(I)I", UnixNativeDispatcher_dup)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.open0(JII)I", UnixNativeDispatcher_open0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.openat0(IJII)I", UnixNativeDispatcher_openat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.close0(I)V", UnixNativeDispatcher_close0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.rewind(J)V", UnixNativeDispatcher_rewind)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getlinelen(J)I", UnixNativeDispatcher_getlinelen)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.link0(JJ)V", UnixNativeDispatcher_link0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.unlink0(J)V", UnixNativeDispatcher_unlink0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.unlinkat0(IJI)V", UnixNativeDispatcher_unlinkat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.mknod0(JIJ)V", UnixNativeDispatcher_mknod0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.rename0(JJ)V", UnixNativeDispatcher_rename0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.renameat0(IJIJ)V", UnixNativeDispatcher_renameat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.mkdir0(JI)V", UnixNativeDispatcher_mkdir0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.rmdir0(J)V", UnixNativeDispatcher_rmdir0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.readlink0(J)[B", UnixNativeDispatcher_readlink0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.realpath0(J)[B", UnixNativeDispatcher_realpath0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.symlink0(JJ)V", UnixNativeDispatcher_symlink0)
	native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.stat0(JLsun/nio/fs/UnixFileAttributes;)I", UnixNativeDispatcher_stat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.lstat0(JLsun/nio/fs/UnixFileAttributes;)V", UnixNativeDispatcher_lstat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fstat0(ILsun/nio/fs/UnixFileAttributes;)V", UnixNativeDispatcher_fstat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fstatat0(IJILsun/nio/fs/UnixFileAttributes;)V", UnixNativeDispatcher_fstatat0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.chown0(JII)V", UnixNativeDispatcher_chown0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.lchown0(JII)V", UnixNativeDispatcher_lchown0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fchown0(III)V", UnixNativeDispatcher_fchown0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.chmod0(JI)V", UnixNativeDispatcher_chmod0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fchmod0(II)V", UnixNativeDispatcher_fchmod0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.utimes0(JJJ)V", UnixNativeDispatcher_utimes0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.futimes0(IJJ)V", UnixNativeDispatcher_futimes0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.futimens0(IJJ)V", UnixNativeDispatcher_futimens0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.lutimes0(JJJ)V", UnixNativeDispatcher_lutimes0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.opendir0(J)J", UnixNativeDispatcher_opendir0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fdopendir(I)J", UnixNativeDispatcher_fdopendir)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.closedir(J)V", UnixNativeDispatcher_closedir)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.readdir0(J)[B", UnixNativeDispatcher_readdir0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.read0(IJI)I", UnixNativeDispatcher_read0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.write0(IJI)I", UnixNativeDispatcher_write0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.access0(JI)I", UnixNativeDispatcher_access0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getpwuid(I)[B", UnixNativeDispatcher_getpwuid)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getgrgid(I)[B", UnixNativeDispatcher_getgrgid)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getpwnam0(J)I", UnixNativeDispatcher_getpwnam0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.getgrnam0(J)I", UnixNativeDispatcher_getgrnam0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.statvfs0(JLsun/nio/fs/UnixFileStoreAttributes;)V", UnixNativeDispatcher_statvfs0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.strerror(I)[B", UnixNativeDispatcher_strerror)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fgetxattr0(IJJI)I", UnixNativeDispatcher_fgetxattr0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fsetxattr0(IJJI)V", UnixNativeDispatcher_fsetxattr0)
	// native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.fremovexattr0(IJ)V", UnixNativeDispatcher_fremovexattr0)
	native.RegisterDefaultNative("sun/nio/fs/UnixNativeDispatcher.init()I", UnixNativeDispatcher_init)
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
func UnixNativeDispatcher_stat0(vm ir.VM) error {
	stack := vm.GetStack()
	pathAddress := stack.GetVarInt64(0)
	path := cutil.GoString(pathAddress)
	println("path:", path)
	stack.PushInt32(0)
	return nil
}

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
