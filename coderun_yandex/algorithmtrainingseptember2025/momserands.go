package algorithmtrainingseptember2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func validateLenghtInput(lenghts [3]float64) bool {
	for _, l := range lenghts {
		if l < 1 || l > 100 {
			return false
		}
	}

	return true
}

func validateSpeedInput(speeds [3]float64) bool {
	for _, l := range speeds {
		if l < 1 || l > 100 {
			return false
		}
	}

	if speeds[0] < speeds[1] {
		return false
	}

	if speeds[1] < speeds[2] {
		return false
	}

	return true
}

// https://coderun.yandex.ru/selections/algorithm-training-september-2025/problems/moms-erands/description
// MomsReands - assignment 2
func MomsReands() {
	firstStrategy := float64(0)
	secondStrategy := float64(0)
	wayFromHomeToShop := float64(0)
	wayFromHomeToPickupPoint := float64(0)

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// input
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	parameters := strings.Fields(line)
	cnt := len(parameters)
	if cnt != 6 {
		panic("input does not match 6")
	}

	lenghts := [3]float64{}
	speeds := [3]float64{}
	for index, p := range parameters {
		parameter, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}

		if index < 3 {
			lenghts[index] = float64(parameter)
		} else {
			speeds[index-3] = float64(parameter)
		}
	}

	if !validateLenghtInput(lenghts) {
		panic("lenght out of range")
	}

	if !validateSpeedInput(speeds) {
		panic("speed out of range")
	}

	wayFromHomeToShop = lenghts[0]
	if wayFromHomeToShop > (lenghts[1] + lenghts[2]) {
		wayFromHomeToShop = lenghts[1] + lenghts[2]
	}

	wayFromHomeToPickupPoint = lenghts[1]
	if wayFromHomeToPickupPoint > (lenghts[0] + lenghts[2]) {
		wayFromHomeToPickupPoint = (lenghts[0] + lenghts[2])
	}

	// strategy 1: Home-Shop-Home-PickUpPoint
	homeShopHome := wayFromHomeToShop/speeds[0] + wayFromHomeToShop/speeds[1]
	homePickUpPointHome := wayFromHomeToPickupPoint/speeds[0] + wayFromHomeToPickupPoint/speeds[1]
	firstStrategy = homeShopHome + homePickUpPointHome

	// strategy 2: Circle
	if wayFromHomeToShop > wayFromHomeToPickupPoint {
		secondStrategy = wayFromHomeToShop/speeds[0] + lenghts[2]/speeds[1] + wayFromHomeToPickupPoint/speeds[2]
	} else {
		secondStrategy = wayFromHomeToPickupPoint/speeds[0] + lenghts[2]/speeds[1] + wayFromHomeToShop/speeds[2]
	}

	writer.WriteString(strconv.FormatFloat(min(firstStrategy, secondStrategy), 'f', -1, 64))
	writer.WriteByte('\n')
}
