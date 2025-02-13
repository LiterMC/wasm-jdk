package jdk_internal_util

import (
	"runtime"
	"strconv"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"

	misc "github.com/LiterMC/wasm-jdk/native/jdk/internal_/misc"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/util/SystemProps$Raw.vmProperties()[Ljava/lang/String;", SystemProps_Raw_vmProperties)
	native.RegisterDefaultNative("jdk/internal/util/SystemProps$Raw.platformProperties()[Ljava/lang/String;", SystemProps_Raw_platformProperties)
}

var vmProperties = []string{
	"java.home", "/java",
}

// private static native String[] vmProperties();
func SystemProps_Raw_vmProperties(vm ir.VM) error {
	propertiesRef := vm.NewArray(desc.DescStringArray, (int32)(len(vmProperties)))
	properties := propertiesRef.GetRefArr()
	for i, v := range vmProperties {
		properties[i] = vm.RefToPtr(vm.NewString(v))
	}
	vm.GetStack().PushRef(propertiesRef)
	return nil
}

const (
	SystemProps_Raw__display_country_NDX = iota // 0
	SystemProps_Raw__display_language_NDX
	SystemProps_Raw__display_script_NDX
	SystemProps_Raw__display_variant_NDX
	SystemProps_Raw__file_encoding_NDX
	SystemProps_Raw__file_separator_NDX
	SystemProps_Raw__format_country_NDX
	SystemProps_Raw__format_language_NDX
	SystemProps_Raw__format_script_NDX
	SystemProps_Raw__format_variant_NDX
	SystemProps_Raw__ftp_nonProxyHosts_NDX
	SystemProps_Raw__ftp_proxyHost_NDX
	SystemProps_Raw__ftp_proxyPort_NDX
	SystemProps_Raw__http_nonProxyHosts_NDX
	SystemProps_Raw__http_proxyHost_NDX
	SystemProps_Raw__http_proxyPort_NDX
	SystemProps_Raw__https_proxyHost_NDX
	SystemProps_Raw__https_proxyPort_NDX
	SystemProps_Raw__java_io_tmpdir_NDX
	SystemProps_Raw__line_separator_NDX
	SystemProps_Raw__os_arch_NDX
	SystemProps_Raw__os_name_NDX
	SystemProps_Raw__os_version_NDX
	SystemProps_Raw__path_separator_NDX
	SystemProps_Raw__socksNonProxyHosts_NDX
	SystemProps_Raw__socksProxyHost_NDX
	SystemProps_Raw__socksProxyPort_NDX
	SystemProps_Raw__stderr_encoding_NDX
	SystemProps_Raw__stdout_encoding_NDX
	SystemProps_Raw__sun_arch_abi_NDX
	SystemProps_Raw__sun_arch_data_model_NDX
	SystemProps_Raw__sun_cpu_endian_NDX
	SystemProps_Raw__sun_cpu_isalist_NDX
	SystemProps_Raw__sun_io_unicode_encoding_NDX
	SystemProps_Raw__sun_jnu_encoding_NDX
	SystemProps_Raw__sun_os_patch_level_NDX
	SystemProps_Raw__user_dir_NDX
	SystemProps_Raw__user_home_NDX
	SystemProps_Raw__user_name_NDX
	SystemProps_Raw_FIXED_LENGTH
)

var (
	cpuEndianStr       string
	unicodeEncodingStr string
)

func init() {
	if misc.BigEndian {
		cpuEndianStr = "big"
		unicodeEncodingStr = "UnicodeBig"
	} else {
		cpuEndianStr = "little"
		unicodeEncodingStr = "UnicodeLittle"
	}
}

// private static native String[] platformProperties();
func SystemProps_Raw_platformProperties(vm ir.VM) error {
	emptyStr := vm.RefToPtr(vm.GetStringInternOrNew(""))
	utf8Str := vm.RefToPtr(vm.GetStringInternOrNew("UTF-8"))

	propertiesRef := vm.NewArray(desc.DescStringArray, SystemProps_Raw_FIXED_LENGTH)
	properties := propertiesRef.GetRefArr()
	for i := range SystemProps_Raw_FIXED_LENGTH {
		properties[i] = emptyStr
	}
	properties[SystemProps_Raw__display_country_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("us"))
	properties[SystemProps_Raw__display_language_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("en"))
	properties[SystemProps_Raw__display_script_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("English"))
	properties[SystemProps_Raw__file_encoding_NDX] = utf8Str
	properties[SystemProps_Raw__file_separator_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("/"))
	properties[SystemProps_Raw__format_country_NDX] = properties[SystemProps_Raw__display_country_NDX]
	properties[SystemProps_Raw__format_language_NDX] = properties[SystemProps_Raw__display_language_NDX]
	properties[SystemProps_Raw__format_script_NDX] = properties[SystemProps_Raw__display_script_NDX]
	properties[SystemProps_Raw__java_io_tmpdir_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("/tmp"))
	properties[SystemProps_Raw__line_separator_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("\n"))
	properties[SystemProps_Raw__os_arch_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(runtime.GOARCH))
	properties[SystemProps_Raw__os_name_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(runtime.GOOS))
	properties[SystemProps_Raw__os_version_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("1.0"))
	properties[SystemProps_Raw__path_separator_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(":"))
	properties[SystemProps_Raw__stderr_encoding_NDX] = utf8Str
	properties[SystemProps_Raw__stdout_encoding_NDX] = utf8Str
	properties[SystemProps_Raw__sun_arch_data_model_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(strconv.Itoa((int)(unsafe.Sizeof(uintptr(0))) * 8)))
	properties[SystemProps_Raw__sun_cpu_endian_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(cpuEndianStr))
	properties[SystemProps_Raw__sun_io_unicode_encoding_NDX] = vm.RefToPtr(vm.GetStringInternOrNew(unicodeEncodingStr))
	properties[SystemProps_Raw__sun_jnu_encoding_NDX] = utf8Str
	properties[SystemProps_Raw__user_dir_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("/wome"))
	properties[SystemProps_Raw__user_home_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("/wome"))
	properties[SystemProps_Raw__user_name_NDX] = vm.RefToPtr(vm.GetStringInternOrNew("browser_user"))
	vm.GetStack().PushRef(propertiesRef)
	return nil
}
