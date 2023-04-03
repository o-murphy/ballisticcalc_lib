from ctypes import *
from pathlib import Path

from flags import *
from .types import BallisticProfile

dllpath = Path(Path(__file__).parent, 'extball.dll')
lib = cdll.LoadLibrary(dllpath.as_posix())


lib.CalculateTrajectory.argtypes = [
    c_double, c_int, c_double, c_double, c_double, c_double,
    c_double, c_int, c_double, c_double, c_double, c_double,
    c_double, c_double, c_double, c_double, c_double,
    c_int, c_int, c_int, c_int, c_int, c_int, c_int,
    c_int, c_int, c_int, c_int, c_int, c_int, c_int,
]
# lib.CalculateTrajectory.argtypes = [BallisticProfile._field_types]

# class CalculateTrajectory_return(Structure):
#     _fields_ = [('bc', c_double), ('dt', c_int)]
# lib.CalculateTrajectory.restype = CalculateTrajectory_return



class BallisticsCalculator:
    def __init__(self, profile: BallisticProfile):
        self.profile = profile


if __name__ == '__main__':
    print(BallisticProfile._field_types)


    example = BallisticProfile(
        bcValue=0.223,
        dragTable=DragTable.G7,
        bulletDiameter=0.308,
        bulletLength=1.282,
        bulletWeight=168,
        muzzleVelocity=2750,
        zeroingDistance=100,
        twistDirection=TwistDirection.Right,
        twistRate=11.24,
        maxShotDistance=1000,
        calculationStep=100,
        maxCalculationSteSize=0,
        windVelocity=5,
        windDirection=-45,
        altitude=11,
        pressure=960,
        temperature=15,
        humidity=50,
        sightHeightUnits=unit.Distance.Inch,
        twistUnits=unit.Distance.Inch,
        velocityUnits=unit.Velocity.FPS,
        distanceUnits=unit.Distance.Foot,
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

    ret = lib.CalculateTrajectory(*example)
