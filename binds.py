import base64
from ctypes import *
from pathlib import Path

from tools import *
import json


LIB_PATH = Path(Path(__file__).parent, 'extball.dll')


class BallisticsCalculator:

    class TrajectoryData(Structure):
        _fields_ = [
            ('traveled_distance', c_double),
            ('velocity', c_double),
            ('time', c_double),
            ('drop', c_double),
            ('drop_adjustment', c_double),
            ('windage', c_double),
            ('windage_adjustment', c_double),
            ('energy', c_double),
            ('optimal_game_weight', c_double),
            ('mach_velocity', c_double),
        ]

    # class CalculateTrajectory_return(Structure):
    #     _fields_ = [
    #         ('bc', c_double),
    #         ('dt', c_int)
    #     ]

    def __init__(self):
        self.lib = cdll.LoadLibrary(LIB_PATH.as_posix())

        self.lib.CalculateTrajectory.argtypes = [
            c_double, c_int, c_double, c_double, c_double, c_double,
            c_double, c_int, c_double, c_double, c_double, c_double,
            c_double, c_double, c_double, c_double, c_double, c_double,
            c_int, c_int, c_int, c_int, c_int, c_int, c_int,
            c_int, c_int, c_int, c_int, c_int, c_int, c_int,
        ]

        self.lib.CalculateTrajectory.restype = c_char_p

    def calculate_trajectory(self, profile: BallisticProfile):
        result_bytes = base64.b64decode(self.lib.CalculateTrajectory(*profile))
        result = result_bytes.decode('utf-8')
        data = json.loads(result)
        return [TrajectoryData(**i) for i in data]

    def fill(self):
        return self.lib.Fill()


if __name__ == '__main__':

    example = BallisticProfile(
        bcValue=0.223,
        dragTable=DragTable.G7,
        bulletDiameter=0.308,
        bulletLength=1.282,
        bulletWeight=168,
        muzzleVelocity=800,
        zeroingDistance=100,
        twistDirection=TwistDirection.Right,
        twistRate=11.24,
        sightHeight=90,
        maxShotDistance=2000,
        calculationStep=100,
        maxCalculationSteSize=0,
        windVelocity=5,
        windDirection=-45,
        altitude=11,
        pressure=960,
        temperature=15,
        humidity=50,
        sightHeightUnits=unit.Distance.Millimeter,
        twistUnits=unit.Distance.Inch,
        velocityUnits=unit.Velocity.MPS,
        distanceUnits=unit.Distance.Meter,
        diameterUnits=unit.Distance.Inch,
        lengthUnits=unit.Distance.Inch,
        weightUnits=unit.Weight.Grain,
        temperatureUnits=unit.Temperature.Celsius,
        pressureUnits=unit.Pressure.MmHg,
        dropUnits=unit.Distance.Centimeter,
        pathUnits=unit.Angular.CmPer100M,
        angularUnits=unit.Angular.Degree,
        energyUnits=unit.Energy.Joule,
        ogwUnits=unit.Weight.Kilogram,
    )

    bc = BallisticsCalculator()
    trajectory = bc.calculate_trajectory(example)
    [print(i) for i in trajectory]
    # bc.fill()

    # lib = cdll.LoadLibrary(LIB_PATH.as_posix())
    # lib.Fill()

    import timeit

    def loops():
        global example
        example = BallisticProfile(
            bcValue=0.223,
            dragTable=DragTable.G7,
            bulletDiameter=0.308,
            bulletLength=1.282,
            bulletWeight=168,
            muzzleVelocity=example.muzzleVelocity + 1,
            zeroingDistance=100,
            twistDirection=TwistDirection.Right,
            twistRate=11.24,
            sightHeight=90,
            maxShotDistance=2000,
            calculationStep=100,
            maxCalculationSteSize=0,
            windVelocity=example.windVelocity + 0.1,
            windDirection=-45,
            altitude=11,
            pressure=example.pressure + 1,
            temperature=15,
            humidity=50,
            sightHeightUnits=unit.Distance.Millimeter,
            twistUnits=unit.Distance.Inch,
            velocityUnits=unit.Velocity.MPS,
            distanceUnits=unit.Distance.Meter,
            diameterUnits=unit.Distance.Inch,
            lengthUnits=unit.Distance.Inch,
            weightUnits=unit.Weight.Grain,
            temperatureUnits=unit.Temperature.Celsius,
            pressureUnits=unit.Pressure.MmHg,
            dropUnits=unit.Distance.Centimeter,
            pathUnits=unit.Angular.CmPer100M,
            angularUnits=unit.Angular.Degree,
            energyUnits=unit.Energy.Joule,
            ogwUnits=unit.Weight.Kilogram,
        )
        bc.calculate_trajectory(example)

    t = timeit.timeit(loops, number=100)
    print(t)