package main

import "C"
import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	extball "github.com/gehtsoft-usa/go_ballisticcalc"
	"github.com/gehtsoft-usa/go_ballisticcalc/bmath/unit"
	"math"
)

type CalculatedTrajectoryData struct {
	TravelledDistance float64
	Velocity          float64
	Time              float64
	Drop              float64
	DropAdjustment    float64
	Windage           float64
	WindageAdjustment float64
	Energy            float64
	OptimalGameWeight float64
	MachVelocity      float64
}

func toJson(data []CalculatedTrajectoryData) *C.char {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	encodedBytes := make([]byte, base64.StdEncoding.EncodedLen(len(jsonData)))
	base64.StdEncoding.Encode(encodedBytes, jsonData)
	return C.CString(string(encodedBytes))
}

//export CalculateTrajectory
func CalculateTrajectory(
	bcValue C.double,
	dragTable C.int,
	bulletDiameter C.double,
	bulletLength C.double,
	bulletWeight C.double,
	muzzleVelocity C.double,
	zeroingDistance C.double,
	twistDirection C.int,
	twistRate C.double,
	sightHeight C.double,
	maxShotDistance C.double,
	calculationStep C.double,
	maxCalculationSteSize C.double,
	windVelocity C.double,
	windDirection C.double,
	altitude C.double,
	pressure C.double,
	temperature C.double,
	humidity C.double,

	sightHeightUnits C.int,
	twistUnits C.int,
	velocityUnits C.int,
	distanceUnits C.int,
	diameterUnits C.int,
	lengthUnits C.int,
	weightUnits C.int,
	temperatureUnits C.int,
	pressureUnits C.int,
	dropUnits C.int,
	pathUnits C.int,
	angularUnits C.int,
	energyUnits C.int,
	ogwUnits C.int,
) *C.char {

	// converting int enum ty byte
	distanceUnitsB := byte(distanceUnits)
	velocityUnitsB := byte(velocityUnits)
	diameterUnitsB := byte(diameterUnits)
	lengthUnitsB := byte(lengthUnits)
	weightUnitsB := byte(weightUnits)
	pressureUnitsB := byte(pressureUnits)
	temperatureUnitsB := byte(temperatureUnits)
	twistUnitsB := byte(twistUnits)
	sightHeightUnitsB := byte(sightHeightUnits)
	angularUnitsB := byte(angularUnits)
	dropUnitsB := byte(dropUnits)
	pathUnitsB := byte(pathUnits)
	energyUnitsB := byte(energyUnits)
	ogwUnitsB := byte(ogwUnits)

	bc, _ := extball.CreateBallisticCoefficient(float64(bcValue), byte(dragTable))

	projectile := extball.CreateProjectileWithDimensions(
		bc,
		unit.MustCreateDistance(float64(bulletDiameter), diameterUnitsB),
		unit.MustCreateDistance(float64(bulletLength), lengthUnitsB),
		unit.MustCreateWeight(float64(bulletWeight), weightUnitsB),
	)
	ammo := extball.CreateAmmunition(
		projectile,
		unit.MustCreateVelocity(float64(muzzleVelocity), velocityUnitsB),
	)
	atmosphere, err := extball.CreateAtmosphere(
		unit.MustCreateDistance(float64(altitude), distanceUnitsB),
		unit.MustCreatePressure(float64(pressure), pressureUnitsB),
		unit.MustCreateTemperature(float64(temperature), temperatureUnitsB),
		float64(humidity),
	)

	if err != nil {
		atmosphere = extball.CreateICAOAtmosphere(
			unit.MustCreateDistance(float64(altitude), distanceUnitsB),
		)
	}

	zero := extball.CreateZeroInfoWithAnotherAmmoAndAtmosphere(
		unit.MustCreateDistance(float64(zeroingDistance), distanceUnitsB),
		ammo, atmosphere,
	)
	twist := extball.CreateTwist(
		byte(twistDirection),
		unit.MustCreateDistance(float64(twistRate), twistUnitsB),
	)
	weapon := extball.CreateWeaponWithTwist(
		unit.MustCreateDistance(float64(sightHeight), sightHeightUnitsB),
		zero,
		twist,
	)

	calc := extball.CreateTrajectoryCalculator()

	if maxCalculationSteSize > 0 {
		calc.SetMaximumCalculatorStepSize(
			unit.MustCreateDistance(float64(maxCalculationSteSize), distanceUnitsB),
		)
	}

	sightAngle := calc.SightAngle(ammo, weapon, atmosphere)

	shotInfo := extball.CreateShotParameters(
		sightAngle,
		unit.MustCreateDistance(float64(maxShotDistance), distanceUnitsB),
		unit.MustCreateDistance(float64(calculationStep), distanceUnitsB),
	)
	wind := extball.CreateOnlyWindInfo(
		unit.MustCreateVelocity(float64(windVelocity), velocityUnitsB),
		unit.MustCreateAngular(float64(windDirection), angularUnitsB),
	)

	data := calc.Trajectory(ammo, weapon, atmosphere, shotInfo, wind)

	ranges := make([]CalculatedTrajectoryData, len(data))

	for i, value := range data {

		windageAdjustment := value.WindageAdjustment().In(pathUnitsB)
		if math.IsNaN(value.WindageAdjustment().In(pathUnitsB)) {
			windageAdjustment = 0
		}

		ranges[i] = CalculatedTrajectoryData{
			TravelledDistance: value.TravelledDistance().In(distanceUnitsB),
			Velocity:          value.Velocity().In(velocityUnitsB),
			Time:              value.Time().Seconds(),
			Drop:              value.Drop().In(dropUnitsB),
			DropAdjustment:    value.DropAdjustment().In(pathUnitsB),
			Windage:           value.Windage().In(dropUnitsB),
			WindageAdjustment: windageAdjustment,
			Energy:            value.Energy().In(energyUnitsB),
			OptimalGameWeight: value.OptimalGameWeight().In(ogwUnitsB),
			MachVelocity:      value.MachVelocity(),
		}

	}

	return toJson(ranges)

}

func main() {

}
