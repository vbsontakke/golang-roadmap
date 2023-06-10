package main

import "fmt"

//

func checksensor(lowerRange int, upperrange int,sensorID string, sensorVal int,sensorcost int){

	if sensorVal < lowerRange || sensorVal > upperrange {
		fmt.Println("Sensor ")
	}
}
type sensorData struct{
	vendorID string
	machineID string
	sensorId string
	sensorVal int
}
type sensorCost struct{
	sensorID string
	lowerLimit int
	upperLimit int
	cost int
}

func main(){
	// 22 37 range val 15
	sensorCost
	for sensorData  range []database{
		for sensorCostData range sensorCostArray{
			strings.HasPrefix(sensorCost.sensorID, sensorData.sensorId)
			if sensorCostData.sensorID == sensorData.sensorId{
				checksensor(sensorCostData.lowerLimit, sensorCostData.upperLimit,sensorData.sensorId,sensorData.sensorVal,sensorCostData.cost)
			}
		}	 
	}
}