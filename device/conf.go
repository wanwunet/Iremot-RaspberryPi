package device

// Physical, BCM, WiringPi
var PinMaping = [][]int8{
	{1, -1, -1}, {2, -1, -1},
	{3, 2, 8}, {4, -1, -1},
	{5, 3, 9}, {6, -1, -1},
	{7, 4, 7}, {8, 14, 15},
	{9, -1, -1}, {10, 15, 16},
	{11, 17, 0}, {12, 18, 1},
	{13, 27, 2}, {14, -1, -1},
	{15, 22, 3}, {16, 23, 4},
	{17, -1, -1}, {18, 24, 5},
	{19, 10, 12}, {20, -1, -1},
	{21, 9, 13}, {22, 25, 6},
	{23, 11, 14}, {24, 8, 10},
	{25, 2, -1}, {26, 7, 11},
	{27, 0, 30}, {28, 1, 31},
	{29, 5, 21}, {30, -1, -1},
	{31, 6, 22}, {32, 12, 26},
	{33, 13, 23}, {34, -1, -1},
	{35, 19, 24}, {36, 16, 27},
	{37, 26, 25}, {38, 20, 28},
	{39, -1, -1}, {40, 21, 29},
}

var GpioPhysical = []int8{
	3, 5, 7, 8, 10, 11, 12, 13, 15, 16, 18, 19, 21, 22, 23, 24, 26, 27, 28, 29, 31, 32, 33, 35, 36, 37, 38, 40,
}

var Physical2BCM = make(map[int8]int8)
var Physical2WiringPi = make(map[int8]int8)
var BCM2Physical = make(map[int8]int8)
var BCM2WiringPi = make(map[int8]int8)
var WiringPi2Physical = make(map[int8]int8)
var WiringPi2BCM = make(map[int8]int8)

func init() {
	for i := 0; i < len(PinMaping); i++ {
		Physical2BCM[PinMaping[i][0]] = PinMaping[i][1]
		Physical2WiringPi[PinMaping[i][0]] = PinMaping[i][2]
		BCM2Physical[PinMaping[i][1]] = PinMaping[i][0]
		BCM2WiringPi[PinMaping[i][1]] = PinMaping[i][2]
		WiringPi2Physical[PinMaping[i][2]] = PinMaping[i][0]
		WiringPi2BCM[PinMaping[i][2]] = PinMaping[i][1]
	}
}
