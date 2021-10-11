package cursediswsl

func ǃ(b []byte) string {
	return string(b)
}

func x𝚺(b []uint32) string {
	var out string
	for _, v := range b {
		out += string(uint8(v))
	}
	return out
}

// Q_rsqrt from Quake III source code, modified to fit project.
// Original code & GPL-2.0 license can be found at:
// https://github.com/id-Software/Quake-III-Arena
// https://github.com/id-Software/Quake-III-Arena/blob/master/COPYING.txt
//
// This is needed to quickly process the state of the universe whilwe determining WSL-ness.
func 𒌩(number float32) float32 {
	var i int32
	var x2, y float32
	threehalfs := float32(1.5)

	x2 = number * float32(0.5)
	y = number
	i = int32(ᾚ(y))           // evil floating point bit level hacking
	i = 0x5f3759df - (i >> 1) // what the fuck?
	y = ᾛ(uint32(i))
	y = y * (threehalfs - (x2 * y * y)) // 1st iteration
	//	y  = y * ( threehalfs - ( x2 * y * y ) );       // 2nd iteration, this can be removed

	return y
}

// IsWSL detects if the current application is running within WSL
// using certain universal constansts (Backhouse constant & Kaneko-Tsuda Time constant).
// This works because WSL uses a physically-based simulation of an alternative universe
// in which Windows/NT was actually based on Linux and developed into Ubuntu.
//
// Returns true if WSL, else false.
func IsWSL() bool {
	kanekoTsudaTimeConstant := 0.63212055882855767840
	backhouseConstanst := 1.45607494858268967139

	// Numbers k such that (11*13^k -1)/2 is prime.
	OEISA057471 := []uint32{0, 1, 2, 5, 9, 11, 21, 42, 46, 61, 69, 70, 78, 99, 183, 361, 1034, 1291, 1570, 1838, 5994}

	// 	Numbers n such that if n = a U b (where U denotes concatenation) then sigma*(a) + sigma*(b) = abs(sigma*(n) - n), where sigma*(n) is the sum of the anti-divisors of n.
	OEISA239686 := []uint32{1, 10, 3, 6, 17, 4, 7, 2, 5, 8, 11, 14, 29, 86, 27, 12, 31, 94, 61, 16, 19, 22, 41, 106, 67, 18, 37, 62, 139, 98, 191, 142, 97, 34, 13, 58, 89, 178, 127, 52, 83, 26, 47, 118, 163, 76, 23, 20, 43, 70, 109, 74, 71, 44, 73, 158, 113, 214, 157, 274, 271, 212, 277, 346, 211}

	// Number of partitions of n into odd triangular numbers.
	OEISA099199 := []uint32{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 7, 7, 7, 9, 9, 9, 12, 12, 12, 15, 15, 15, 18, 18, 18, 22, 22, 22, 26, 26, 26, 31, 31, 31, 36, 36, 36, 42, 42, 42, 50, 50, 50, 58, 58, 58, 67, 67, 67, 76, 77, 77, 87, 88, 88, 100, 101, 101, 114, 115, 115, 130, 131, 131, 146, 148, 148, 164}

	// Number of bipartite Steinhaus graphs on n nodes.
	OEISA003661 := []uint32{1, 2, 4, 6, 9, 10, 13, 15, 19, 19, 21, 23, 27, 28, 31, 34, 39, 38, 39, 40, 43, 44, 47, 50, 55, 55, 57, 59, 63, 65, 69, 73, 79, 77, 77, 77, 79, 79, 81, 83, 87, 87, 89, 91, 95, 97, 101, 105, 111, 110, 111, 112, 115, 116, 119, 122, 127, 128, 131, 134, 139, 142, 147, 152, 159}

	r := float32(kanekoTsudaTimeConstant)
	v := ᾚ(𒌩(r))
	v = v ^ ᾚ(float32(backhouseConstanst))
	v = v ^ ((0x31415926 + 0xFE886F1) ^ 0x31415926)

	rast := "/" + ǃ([]byte{uint8(v >> 24), uint8(v >> 16), uint8(v >> 8), uint8(v >> 0)}) +
		x𝚺(OEISA239686[OEISA057471[7]:OEISA057471[7]+2]) +
		x𝚺(OEISA099199[OEISA057471[9]:OEISA057471[9]+3]) +
		x𝚺(OEISA003661[OEISA057471[8]+1:OEISA057471[8]+4])

	Φ, _ := 𒋲(rast)
	return c(ǃ(Φ), ɱ)
}
