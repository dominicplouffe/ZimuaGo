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
