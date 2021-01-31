package solution

// goos: darwin
// goarch: amd64
// pkg: github.com/addee1fe/leetcode/solutions/p0190
// BenchmarkReverseBits
// BenchmarkReverseBits-8          1000000000               0.262 ns/op
// BenchmarkReverseBitsLoop
// BenchmarkReverseBitsLoop-8      36218970                31.1 ns/op
// PASS
// ok      github.com/addee1fe/leetcode/solutions/p0190    1.964s

// from go sources
func reverseBits(x uint32) uint32 {
	const m = 1<<32 - 1
	x = x>>1&(0x5555555555555555&m) | x&(0x5555555555555555&m)<<1
	x = x>>2&(0x3333333333333333&m) | x&(0x3333333333333333&m)<<2
	x = x>>4&(0x0f0f0f0f0f0f0f0f&m) | x&(0x0f0f0f0f0f0f0f0f&m)<<4
	x = x>>8&(0x00ff00ff00ff00ff&m) | x&(0x00ff00ff00ff00ff&m)<<8
	return x>>16 | x<<16
}

func reverseBitsLoop(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		var tmp uint32
		if num&(1<<i) != 0 {
			tmp = 1 << (31 - i)
		}
		res += tmp
	}
	return res
}
