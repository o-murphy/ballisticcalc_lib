package main

import (
	"C"
	"fmt"
	extball "github.com/gehtsoft-usa/go_ballisticcalc"
	"github.com/gehtsoft-usa/go_ballisticcalc/bmath/unit"
	"reflect"
)

////export CreateAngular
//func CreateAngular(value C.double, units C.int) *unit.Angular {
//	angular, _ := unit.CreateAngular(float64(value), byte(units))
//	return &angular
//}

//export CreateBallisticCoefficient
func CreateBallisticCoefficient(value C.double, DragTable C.int) (C.double, C.int) {
	bc, _ := extball.CreateBallisticCoefficient(float64(value), byte(DragTable))
	return C.double(bc.Value()), C.int(bc.Table())
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
) {

	bc, _ := extball.CreateBallisticCoefficient(float64(bcValue), byte(dragTable))
	projectile := extball.CreateProjectileWithDimensions(
		bc,
		unit.MustCreateDistance(float64(bulletDiameter), byte(diameterUnits)),
		unit.MustCreateDistance(float64(bulletLength), byte(lengthUnits)),
		unit.MustCreateWeight(float64(bulletWeight), byte(weightUnits)),
	)
	ammo := extball.CreateAmmunition(
		projectile,
		unit.MustCreateVelocity(float64(muzzleVelocity), byte(velocityUnits)),
	)
	atmosphere, error := extball.CreateAtmosphere(
		unit.MustCreateDistance(float64(altitude), byte(distanceUnits)),
		unit.MustCreatePressure(float64(pressure), byte(pressureUnits)),
		unit.MustCreateTemperature(float64(temperature), byte(temperatureUnits)),
		float64(humidity),
	)

	if error != nil {
		atmosphere = extball.CreateICAOAtmosphere(unit.MustCreateDistance(float64(altitude), byte(distanceUnits)))
	}

	zero := extball.CreateZeroInfoWithAnotherAmmoAndAtmosphere(
		unit.MustCreateDistance(float64(zeroingDistance), byte(distanceUnits)),
		ammo, atmosphere,
	)
	twist := extball.CreateTwist(
		byte(twistDirection),
		unit.MustCreateDistance(float64(twistRate), byte(twistUnits)),
	)
	weapon := extball.CreateWeaponWithTwist(
		unit.MustCreateDistance(2, byte(sightHeightUnits)),
		zero,
		twist,
	)

	calc := extball.CreateTrajectoryCalculator()

	if maxCalculationSteSize > 0 {
		calc.SetMaximumCalculatorStepSize(
			unit.MustCreateDistance(float64(maxCalculationSteSize), byte(distanceUnits)),
		)
	}

	sightAngle := calc.SightAngle(ammo, weapon, atmosphere)

	shotInfo := extball.CreateShotParameters(
		sightAngle,
		unit.MustCreateDistance(float64(maxShotDistance), byte(distanceUnits)),
		unit.MustCreateDistance(float64(calculationStep), byte(distanceUnits)),
	)
	wind := extball.CreateOnlyWindInfo(
		unit.MustCreateVelocity(float64(windVelocity), byte(velocityUnits)),
		unit.MustCreateAngular(float64(windDirection), byte(angularUnits)),
	)

	data := calc.Trajectory(ammo, weapon, atmosphere, shotInfo, wind)
	//return data
	for _, value := range data {
		fmt.Println(
			value.TravelledDistance().In(byte(distanceUnits)),
			value.Velocity().In(byte(velocityUnits)),
			value.Time().Seconds(),
			value.Drop().In(byte(dropUnits)),
			value.DropAdjustment().In(byte(pathUnits)),
			value.Windage().In(byte(dropUnits)),
			value.WindageAdjustment().In(byte(pathUnits)),
			value.Energy().In(byte(energyUnits)),
			value.OptimalGameWeight().In(byte(ogwUnits)),
			value.MachVelocity(),
		)
	}
}

func main() {
	angular, _ := unit.CreateAngular(10, unit.AngularMil)
	fmt.Println(angular, reflect.TypeOf(angular), &angular)
}
