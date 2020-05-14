package main

var bbCENTER uint64 = 103481868288
var bbOUTERCENTER uint64 = 66125924401152

var bbFILEA uint64 = 72340172838076673
var bbFILEB uint64 = 144680345676153346
var bbFILEC uint64 = 289360691352306692
var bbFILED uint64 = 578721382704613384
var bbFILEE uint64 = 1157442765409226768
var bbFILEF uint64 = 2314885530818453536
var bbFILEG uint64 = 4629771061636907072
var bbFILEH uint64 = 9259542123273814144

var bbRANK1 uint64 = 255
var bbRANK2 uint64 = 65280
var bbRANK3 uint64 = 16711680
var bbRANK4 uint64 = 4278190080
var bbRANK5 uint64 = 1095216660480
var bbRANK6 uint64 = 280375465082880
var bbRANK7 uint64 = 71776119061217280
var bbRANK8 uint64 = 18374686479671623680

func getFile(p uint64) uint64 {

	if p&bbFILEA != 0 {
		return bbFILEA
	} else if p&bbFILEB != 0 {
		return bbFILEB
	} else if p&bbFILEC != 0 {
		return bbFILEC
	} else if p&bbFILED != 0 {
		return bbFILED
	} else if p&bbFILEE != 0 {
		return bbFILEE
	} else if p&bbFILEF != 0 {
		return bbFILEF
	} else if p&bbFILEG != 0 {
		return bbFILEG
	}

	return bbFILEH
}

func getRank(p uint64) uint64 {

	if p&bbRANK1 != 0 {
		return bbRANK1
	} else if p&bbRANK2 != 0 {
		return bbRANK2
	} else if p&bbRANK3 != 0 {
		return bbRANK3
	} else if p&bbRANK4 != 0 {
		return bbRANK4
	} else if p&bbRANK5 != 0 {
		return bbRANK5
	} else if p&bbRANK6 != 0 {
		return bbRANK6
	} else if p&bbRANK7 != 0 {
		return bbRANK7
	}

	return bbRANK8
}

var squareIndexes = make(map[uint64]int)

func initSquareIndexes() {
	squareIndexes[1] = 0
	squareIndexes[2] = 1
	squareIndexes[4] = 2
	squareIndexes[8] = 3
	squareIndexes[16] = 4
	squareIndexes[32] = 5
	squareIndexes[64] = 6
	squareIndexes[128] = 7
	squareIndexes[256] = 8
	squareIndexes[512] = 9
	squareIndexes[1024] = 10
	squareIndexes[2048] = 11
	squareIndexes[4096] = 12
	squareIndexes[8192] = 13
	squareIndexes[16384] = 14
	squareIndexes[32768] = 15
	squareIndexes[65536] = 16
	squareIndexes[131072] = 17
	squareIndexes[262144] = 18
	squareIndexes[524288] = 19
	squareIndexes[1048576] = 20
	squareIndexes[2097152] = 21
	squareIndexes[4194304] = 22
	squareIndexes[8388608] = 23
	squareIndexes[16777216] = 24
	squareIndexes[33554432] = 25
	squareIndexes[67108864] = 26
	squareIndexes[134217728] = 27
	squareIndexes[268435456] = 28
	squareIndexes[536870912] = 29
	squareIndexes[1073741824] = 30
	squareIndexes[2147483648] = 31
	squareIndexes[4294967296] = 32
	squareIndexes[8589934592] = 33
	squareIndexes[17179869184] = 34
	squareIndexes[34359738368] = 35
	squareIndexes[68719476736] = 36
	squareIndexes[137438953472] = 37
	squareIndexes[274877906944] = 38
	squareIndexes[549755813888] = 39
	squareIndexes[1099511627776] = 40
	squareIndexes[2199023255552] = 41
	squareIndexes[4398046511104] = 42
	squareIndexes[8796093022208] = 43
	squareIndexes[17592186044416] = 44
	squareIndexes[35184372088832] = 45
	squareIndexes[70368744177664] = 46
	squareIndexes[140737488355328] = 47
	squareIndexes[281474976710656] = 48
	squareIndexes[562949953421312] = 49
	squareIndexes[1125899906842624] = 50
	squareIndexes[2251799813685248] = 51
	squareIndexes[4503599627370496] = 52
	squareIndexes[9007199254740992] = 53
	squareIndexes[18014398509481984] = 54
	squareIndexes[36028797018963968] = 55
	squareIndexes[72057594037927936] = 56
	squareIndexes[144115188075855872] = 57
	squareIndexes[288230376151711744] = 58
	squareIndexes[576460752303423488] = 59
	squareIndexes[1152921504606846976] = 60
	squareIndexes[2305843009213693952] = 61
	squareIndexes[4611686018427387904] = 62
	squareIndexes[9223372036854775808] = 63
}
