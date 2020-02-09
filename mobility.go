package main

func blackPawnMobility(p int) int {
	switch p {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 0
	case 3:
		return 0
	case 4:
		return 0
	case 5:
		return 0
	case 6:
		return 0
	case 7:
		return 0
	case 8:
		return -10
	case 9:
		return 3
	case 10:
		return -14
	case 11:
		return -36
	case 12:
		return -37
	case 13:
		return -7
	case 14:
		return 8
	case 15:
		return -10
	case 16:
		return 10
	case 17:
		return 3
	case 18:
		return -2
	case 19:
		return 10
	case 20:
		return 10
	case 21:
		return 5
	case 22:
		return 9
	case 23:
		return 10
	case 24:
		return 10
	case 25:
		return 0
	case 26:
		return 1
	case 27:
		return 20
	case 28:
		return 20
	case 29:
		return 10
	case 30:
		return 3
	case 31:
		return 10
	case 32:
		return 20
	case 33:
		return 15
	case 34:
		return 0
	case 35:
		return 14
	case 36:
		return 15
	case 37:
		return -2
	case 38:
		return 16
	case 39:
		return 20
	case 40:
		return 20
	case 41:
		return 44
	case 42:
		return 31
	case 43:
		return 40
	case 44:
		return 44
	case 45:
		return 21
	case 46:
		return 29
	case 47:
		return 20
	case 48:
		return 90
	case 49:
		return 85
	case 50:
		return 82
	case 51:
		return 102
	case 52:
		return 73
	case 53:
		return 86
	case 54:
		return 83
	case 55:
		return 78
	case 56:
		return 900
	case 57:
		return 900
	case 58:
		return 900
	case 59:
		return 900
	case 60:
		return 900
	case 61:
		return 900
	case 62:
		return 900
	case 63:
		return 900
	}

	return 0
}

func blackKnightMobility(p int) int {
	switch p {
	case 0:
		return -69
	case 1:
		return -22
	case 2:
		return -35
	case 3:
		return -19
	case 4:
		return -24
	case 5:
		return -26
	case 6:
		return -23
	case 7:
		return -74
	case 8:
		return -20
	case 9:
		return -23
	case 10:
		return 0
	case 11:
		return 2
	case 12:
		return 0
	case 13:
		return 2
	case 14:
		return -15
	case 15:
		return -23
	case 16:
		return -14
	case 17:
		return 11
	case 18:
		return 25
	case 19:
		return 18
	case 20:
		return 22
	case 21:
		return 25
	case 22:
		return 10
	case 23:
		return -18
	case 24:
		return 0
	case 25:
		return 2
	case 26:
		return 35
	case 27:
		return 22
	case 28:
		return 21
	case 29:
		return 31
	case 30:
		return 5
	case 31:
		return -1
	case 32:
		return 17
	case 33:
		return 25
	case 34:
		return 41
	case 35:
		return 33
	case 36:
		return 37
	case 37:
		return 45
	case 38:
		return 24
	case 39:
		return 24
	case 40:
		return -2
	case 41:
		return 62
	case 42:
		return 100
	case 43:
		return 73
	case 44:
		return 74
	case 45:
		return 100
	case 46:
		return 67
	case 47:
		return 10
	case 48:
		return -14
	case 49:
		return -4
	case 50:
		return 62
	case 51:
		return 4
	case 52:
		return -36
	case 53:
		return 100
	case 54:
		return -6
	case 55:
		return -3
	case 56:
		return -70
	case 57:
		return -58
	case 58:
		return -55
	case 59:
		return -10
	case 60:
		return -75
	case 61:
		return -75
	case 62:
		return -53
	case 63:
		return -66
	}

	return 0
}

func blackBishopMobility(p int) int {
	switch p {
	case 0:
		return -10
	case 1:
		return -10
	case 2:
		return -15
	case 3:
		return -14
	case 4:
		return -12
	case 5:
		return -15
	case 6:
		return 2
	case 7:
		return -7
	case 8:
		return 16
	case 9:
		return 20
	case 10:
		return 6
	case 11:
		return 7
	case 12:
		return 6
	case 13:
		return 11
	case 14:
		return 20
	case 15:
		return 19
	case 16:
		return 15
	case 17:
		return 20
	case 18:
		return 25
	case 19:
		return 8
	case 20:
		return 15
	case 21:
		return 24
	case 22:
		return 25
	case 23:
		return 14
	case 24:
		return 7
	case 25:
		return 0
	case 26:
		return 16
	case 27:
		return 17
	case 28:
		return 23
	case 29:
		return 17
	case 30:
		return 10
	case 31:
		return 13
	case 32:
		return 10
	case 33:
		return 15
	case 34:
		return 25
	case 35:
		return 26
	case 36:
		return 34
	case 37:
		return 20
	case 38:
		return 17
	case 39:
		return 25
	case 40:
		return -14
	case 41:
		return 28
	case 42:
		return -10
	case 43:
		return 52
	case 44:
		return 41
	case 45:
		return -32
	case 46:
		return 39
	case 47:
		return -9
	case 48:
		return -22
	case 49:
		return 2
	case 50:
		return 31
	case 51:
		return -39
	case 52:
		return -42
	case 53:
		return 35
	case 54:
		return 20
	case 55:
		return -11
	case 56:
		return -50
	case 57:
		return -37
	case 58:
		return -107
	case 59:
		return -23
	case 60:
		return -76
	case 61:
		return -82
	case 62:
		return -78
	case 63:
		return -59
	}

	return 0
}

func blackRookMobility(p int) int {
	switch p {
	case 0:
		return -32
	case 1:
		return -31
	case 2:
		return -18
	case 3:
		return -2
	case 4:
		return 5
	case 5:
		return -18
	case 6:
		return -24
	case 7:
		return -30
	case 8:
		return -53
	case 9:
		return -44
	case 10:
		return -43
	case 11:
		return -29
	case 12:
		return -26
	case 13:
		return -31
	case 14:
		return -38
	case 15:
		return -53
	case 16:
		return -46
	case 17:
		return -26
	case 18:
		return -35
	case 19:
		return -25
	case 20:
		return -25
	case 21:
		return -42
	case 22:
		return -28
	case 23:
		return -42
	case 24:
		return -30
	case 25:
		return -46
	case 26:
		return -29
	case 27:
		return -13
	case 28:
		return -21
	case 29:
		return -16
	case 30:
		return -35
	case 31:
		return -28
	case 32:
		return -6
	case 33:
		return -9
	case 34:
		return -4
	case 35:
		return 18
	case 36:
		return 13
	case 37:
		return 16
	case 38:
		return 5
	case 39:
		return 0
	case 40:
		return 15
	case 41:
		return 25
	case 42:
		return 27
	case 43:
		return 45
	case 44:
		return 33
	case 45:
		return 28
	case 46:
		return 35
	case 47:
		return 19
	case 48:
		return 60
	case 49:
		return 34
	case 50:
		return 62
	case 51:
		return 55
	case 52:
		return 67
	case 53:
		return 56
	case 54:
		return 29
	case 55:
		return 55
	case 56:
		return 50
	case 57:
		return 56
	case 58:
		return 33
	case 59:
		return 37
	case 60:
		return 4
	case 61:
		return 33
	case 62:
		return 29
	case 63:
		return 35
	}

	return 0
}

func blackQueenMobility(p int) int {
	switch p {
	case 0:
		return -42
	case 1:
		return -34
	case 2:
		return -36
	case 3:
		return -31
	case 4:
		return -13
	case 5:
		return -31
	case 6:
		return -30
	case 7:
		return -39
	case 8:
		return -38
	case 9:
		return -21
	case 10:
		return -15
	case 11:
		return -15
	case 12:
		return -19
	case 13:
		return 0
	case 14:
		return -18
	case 15:
		return -36
	case 16:
		return -27
	case 17:
		return -16
	case 18:
		return -11
	case 19:
		return -16
	case 20:
		return -11
	case 21:
		return -13
	case 22:
		return -6
	case 23:
		return -30
	case 24:
		return -22
	case 25:
		return -20
	case 26:
		return -10
	case 27:
		return -1
	case 28:
		return -5
	case 29:
		return -2
	case 30:
		return -15
	case 31:
		return -14
	case 32:
		return -6
	case 33:
		return -13
	case 34:
		return 20
	case 35:
		return 25
	case 36:
		return 17
	case 37:
		return 22
	case 38:
		return -16
	case 39:
		return 1
	case 40:
		return 2
	case 41:
		return 43
	case 42:
		return 63
	case 43:
		return 72
	case 44:
		return 60
	case 45:
		return 32
	case 46:
		return 43
	case 47:
		return -2
	case 48:
		return 24
	case 49:
		return 57
	case 50:
		return 76
	case 51:
		return 20
	case 52:
		return -10
	case 53:
		return 60
	case 54:
		return 32
	case 55:
		return 14
	case 56:
		return 26
	case 57:
		return 88
	case 58:
		return 24
	case 59:
		return 69
	case 60:
		return -104
	case 61:
		return -8
	case 62:
		return 1
	case 63:
		return 6
	}

	return 0
}

func blackKingMobility(p int) int {
	switch p {
	case 0:
		return 18
	case 1:
		return 40
	case 2:
		return -1
	case 3:
		return 6
	case 4:
		return -14
	case 5:
		return -3
	case 6:
		return 30
	case 7:
		return 17
	case 8:
		return 4
	case 9:
		return 13
	case 10:
		return -18
	case 11:
		return -57
	case 12:
		return -50
	case 13:
		return -14
	case 14:
		return 3
	case 15:
		return -4
	case 16:
		return -32
	case 17:
		return -29
	case 18:
		return -32
	case 19:
		return -64
	case 20:
		return -79
	case 21:
		return -43
	case 22:
		return -42
	case 23:
		return -47
	case 24:
		return -50
	case 25:
		return -8
	case 26:
		return -47
	case 27:
		return -51
	case 28:
		return -28
	case 29:
		return -52
	case 30:
		return -43
	case 31:
		return -55
	case 32:
		return -49
	case 33:
		return 0
	case 34:
		return 13
	case 35:
		return -19
	case 36:
		return -4
	case 37:
		return 11
	case 38:
		return 50
	case 39:
		return -55
	case 40:
		return -31
	case 41:
		return 37
	case 42:
		return 28
	case 43:
		return -67
	case 44:
		return 44
	case 45:
		return -57
	case 46:
		return 12
	case 47:
		return -62
	case 48:
		return 3
	case 49:
		return 10
	case 50:
		return 55
	case 51:
		return 56
	case 52:
		return 56
	case 53:
		return 55
	case 54:
		return 10
	case 55:
		return -32
	case 56:
		return -62
	case 57:
		return 83
	case 58:
		return 60
	case 59:
		return -99
	case 60:
		return -99
	case 61:
		return 47
	case 62:
		return 54
	case 63:
		return 4
	}

	return 0
}

func whitePawnMobility(p int) int {
	switch p {
	case 0:
		return 900
	case 1:
		return 900
	case 2:
		return 900
	case 3:
		return 900
	case 4:
		return 900
	case 5:
		return 900
	case 6:
		return 900
	case 7:
		return 900
	case 8:
		return 90
	case 9:
		return 85
	case 10:
		return 82
	case 11:
		return 102
	case 12:
		return 73
	case 13:
		return 86
	case 14:
		return 83
	case 15:
		return 78
	case 16:
		return 20
	case 17:
		return 44
	case 18:
		return 31
	case 19:
		return 40
	case 20:
		return 44
	case 21:
		return 21
	case 22:
		return 29
	case 23:
		return 20
	case 24:
		return 20
	case 25:
		return 15
	case 26:
		return 0
	case 27:
		return 14
	case 28:
		return 15
	case 29:
		return -2
	case 30:
		return 16
	case 31:
		return 20
	case 32:
		return 10
	case 33:
		return 0
	case 34:
		return 1
	case 35:
		return 20
	case 36:
		return 20
	case 37:
		return 10
	case 38:
		return 3
	case 39:
		return 10
	case 40:
		return 10
	case 41:
		return 3
	case 42:
		return -2
	case 43:
		return 10
	case 44:
		return 10
	case 45:
		return 5
	case 46:
		return 9
	case 47:
		return 10
	case 48:
		return -10
	case 49:
		return 3
	case 50:
		return -14
	case 51:
		return -36
	case 52:
		return -37
	case 53:
		return -7
	case 54:
		return 8
	case 55:
		return -10
	case 56:
		return 0
	case 57:
		return 0
	case 58:
		return 0
	case 59:
		return 0
	case 60:
		return 0
	case 61:
		return 0
	case 62:
		return 0
	case 63:
		return 0
	}

	return 0
}

func whiteKnightMobility(p int) int {
	switch p {
	case 0:
		return -70
	case 1:
		return -58
	case 2:
		return -55
	case 3:
		return -10
	case 4:
		return -75
	case 5:
		return -75
	case 6:
		return -53
	case 7:
		return -66
	case 8:
		return -14
	case 9:
		return -4
	case 10:
		return 62
	case 11:
		return 4
	case 12:
		return -36
	case 13:
		return 100
	case 14:
		return -6
	case 15:
		return -3
	case 16:
		return -2
	case 17:
		return 62
	case 18:
		return 100
	case 19:
		return 73
	case 20:
		return 74
	case 21:
		return 100
	case 22:
		return 67
	case 23:
		return 10
	case 24:
		return 17
	case 25:
		return 25
	case 26:
		return 41
	case 27:
		return 33
	case 28:
		return 37
	case 29:
		return 45
	case 30:
		return 24
	case 31:
		return 24
	case 32:
		return 0
	case 33:
		return 2
	case 34:
		return 35
	case 35:
		return 22
	case 36:
		return 21
	case 37:
		return 31
	case 38:
		return 5
	case 39:
		return -1
	case 40:
		return -14
	case 41:
		return 11
	case 42:
		return 25
	case 43:
		return 18
	case 44:
		return 22
	case 45:
		return 25
	case 46:
		return 10
	case 47:
		return -18
	case 48:
		return -20
	case 49:
		return -23
	case 50:
		return 0
	case 51:
		return 2
	case 52:
		return 0
	case 53:
		return 2
	case 54:
		return -15
	case 55:
		return -23
	case 56:
		return -69
	case 57:
		return -22
	case 58:
		return -35
	case 59:
		return -19
	case 60:
		return -24
	case 61:
		return -26
	case 62:
		return -23
	case 63:
		return -74
	}

	return 0
}

func whiteBishopMobility(p int) int {
	switch p {
	case 0:
		return -50
	case 1:
		return -37
	case 2:
		return -107
	case 3:
		return -23
	case 4:
		return -76
	case 5:
		return -82
	case 6:
		return -78
	case 7:
		return -59
	case 8:
		return -22
	case 9:
		return 2
	case 10:
		return 31
	case 11:
		return -39
	case 12:
		return -42
	case 13:
		return 35
	case 14:
		return 20
	case 15:
		return -11
	case 16:
		return -14
	case 17:
		return 28
	case 18:
		return -10
	case 19:
		return 52
	case 20:
		return 41
	case 21:
		return -32
	case 22:
		return 39
	case 23:
		return -9
	case 24:
		return 10
	case 25:
		return 15
	case 26:
		return 25
	case 27:
		return 26
	case 28:
		return 34
	case 29:
		return 20
	case 30:
		return 17
	case 31:
		return 25
	case 32:
		return 7
	case 33:
		return 0
	case 34:
		return 16
	case 35:
		return 17
	case 36:
		return 23
	case 37:
		return 17
	case 38:
		return 10
	case 39:
		return 13
	case 40:
		return 15
	case 41:
		return 20
	case 42:
		return 25
	case 43:
		return 8
	case 44:
		return 15
	case 45:
		return 24
	case 46:
		return 25
	case 47:
		return 14
	case 48:
		return 16
	case 49:
		return 20
	case 50:
		return 6
	case 51:
		return 7
	case 52:
		return 6
	case 53:
		return 11
	case 54:
		return 20
	case 55:
		return 19
	case 56:
		return -10
	case 57:
		return -10
	case 58:
		return -15
	case 59:
		return -14
	case 60:
		return -12
	case 61:
		return -15
	case 62:
		return 2
	case 63:
		return -7
	}

	return 0
}

func whiteRookMobility(p int) int {
	switch p {
	case 0:
		return 50
	case 1:
		return 56
	case 2:
		return 33
	case 3:
		return 37
	case 4:
		return 4
	case 5:
		return 33
	case 6:
		return 29
	case 7:
		return 35
	case 8:
		return 60
	case 9:
		return 34
	case 10:
		return 62
	case 11:
		return 55
	case 12:
		return 67
	case 13:
		return 56
	case 14:
		return 29
	case 15:
		return 55
	case 16:
		return 15
	case 17:
		return 25
	case 18:
		return 27
	case 19:
		return 45
	case 20:
		return 33
	case 21:
		return 28
	case 22:
		return 35
	case 23:
		return 19
	case 24:
		return -6
	case 25:
		return -9
	case 26:
		return -4
	case 27:
		return 18
	case 28:
		return 13
	case 29:
		return 16
	case 30:
		return 5
	case 31:
		return 0
	case 32:
		return -30
	case 33:
		return -46
	case 34:
		return -29
	case 35:
		return -13
	case 36:
		return -21
	case 37:
		return -16
	case 38:
		return -35
	case 39:
		return -28
	case 40:
		return -46
	case 41:
		return -26
	case 42:
		return -35
	case 43:
		return -25
	case 44:
		return -25
	case 45:
		return -42
	case 46:
		return -28
	case 47:
		return -42
	case 48:
		return -53
	case 49:
		return -44
	case 50:
		return -43
	case 51:
		return -29
	case 52:
		return -26
	case 53:
		return -31
	case 54:
		return -38
	case 55:
		return -53
	case 56:
		return -32
	case 57:
		return -31
	case 58:
		return -18
	case 59:
		return -2
	case 60:
		return 5
	case 61:
		return -18
	case 62:
		return -24
	case 63:
		return -30
	}

	return 0
}

func whiteQueenMobility(p int) int {
	switch p {
	case 0:
		return 26
	case 1:
		return 88
	case 2:
		return 24
	case 3:
		return 69
	case 4:
		return -104
	case 5:
		return -8
	case 6:
		return 1
	case 7:
		return 6
	case 8:
		return 24
	case 9:
		return 57
	case 10:
		return 76
	case 11:
		return 20
	case 12:
		return -10
	case 13:
		return 60
	case 14:
		return 32
	case 15:
		return 14
	case 16:
		return 2
	case 17:
		return 43
	case 18:
		return 63
	case 19:
		return 72
	case 20:
		return 60
	case 21:
		return 32
	case 22:
		return 43
	case 23:
		return -2
	case 24:
		return -6
	case 25:
		return -13
	case 26:
		return 20
	case 27:
		return 25
	case 28:
		return 17
	case 29:
		return 22
	case 30:
		return -16
	case 31:
		return 1
	case 32:
		return -22
	case 33:
		return -20
	case 34:
		return -10
	case 35:
		return -1
	case 36:
		return -5
	case 37:
		return -2
	case 38:
		return -15
	case 39:
		return -14
	case 40:
		return -27
	case 41:
		return -16
	case 42:
		return -11
	case 43:
		return -16
	case 44:
		return -11
	case 45:
		return -13
	case 46:
		return -6
	case 47:
		return -30
	case 48:
		return -38
	case 49:
		return -21
	case 50:
		return -15
	case 51:
		return -15
	case 52:
		return -19
	case 53:
		return 0
	case 54:
		return -18
	case 55:
		return -36
	case 56:
		return -42
	case 57:
		return -34
	case 58:
		return -36
	case 59:
		return -31
	case 60:
		return -13
	case 61:
		return -31
	case 62:
		return -30
	case 63:
		return -39
	}

	return 0
}

func whiteKingMobility(p int) int {
	switch p {
	case 0:
		return -62
	case 1:
		return 83
	case 2:
		return 60
	case 3:
		return -99
	case 4:
		return -99
	case 5:
		return 47
	case 6:
		return 54
	case 7:
		return 4
	case 8:
		return 3
	case 9:
		return 10
	case 10:
		return 55
	case 11:
		return 56
	case 12:
		return 56
	case 13:
		return 55
	case 14:
		return 10
	case 15:
		return -32
	case 16:
		return -31
	case 17:
		return 37
	case 18:
		return 28
	case 19:
		return -67
	case 20:
		return 44
	case 21:
		return -57
	case 22:
		return 12
	case 23:
		return -62
	case 24:
		return -49
	case 25:
		return 0
	case 26:
		return 13
	case 27:
		return -19
	case 28:
		return -4
	case 29:
		return 11
	case 30:
		return 50
	case 31:
		return -55
	case 32:
		return -50
	case 33:
		return -8
	case 34:
		return -47
	case 35:
		return -51
	case 36:
		return -28
	case 37:
		return -52
	case 38:
		return -43
	case 39:
		return -55
	case 40:
		return -32
	case 41:
		return -29
	case 42:
		return -32
	case 43:
		return -64
	case 44:
		return -79
	case 45:
		return -43
	case 46:
		return -42
	case 47:
		return -47
	case 48:
		return 4
	case 49:
		return 13
	case 50:
		return -18
	case 51:
		return -57
	case 52:
		return -50
	case 53:
		return -14
	case 54:
		return 3
	case 55:
		return -4
	case 56:
		return 18
	case 57:
		return 40
	case 58:
		return -1
	case 59:
		return 6
	case 60:
		return -14
	case 61:
		return -3
	case 62:
		return 30
	case 63:
		return 17
	}

	return 0
}