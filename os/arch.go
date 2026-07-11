package os

import "runtime"

// Arch constants are the recognized [runtime.GOARCH] values. Go exposes GOARCH
// only as a string, so these name the tokens to avoid scattering string
// literals across build-time comparisons.
const (
	Arch386      = "386"
	ArchAMD64    = "amd64"
	ArchARM      = "arm"
	ArchARM64    = "arm64"
	ArchLoong64  = "loong64"
	ArchMIPS     = "mips"
	ArchMIPS64   = "mips64"
	ArchMIPS64LE = "mips64le"
	ArchMIPSLE   = "mipsle"
	ArchPPC64    = "ppc64"
	ArchPPC64LE  = "ppc64le"
	ArchRISCV64  = "riscv64"
	ArchS390X    = "s390x"
	ArchWASM     = "wasm"
)

// IsWasm reports whether the program was compiled to WebAssembly. Unlike the OS
// predicates it checks the architecture ([runtime.GOARCH]), not the OS, since
// WebAssembly runs under either the `js` or `wasip1` GOOS.
func IsWasm() bool {
	return runtime.GOARCH == ArchWASM
}
