from typing import NamedTuple

from . import unit
from .flags import *


class BallisticProfile(NamedTuple):
    bcValue: float
    dragTable: DragTable
    bulletDiameter: float
    bulletLength: float
    bulletWeight: float
    muzzleVelocity: float
    zeroingDistance: float
    twistDirection: TwistDirection
    twistRate: float
    maxShotDistance: float
    calculationStep: float
    maxCalculationSteSize: float
    windVelocity: float
    windDirection: float
    altitude: float
    pressure: float
    temperature: float
    humidity: float

    # units
    sightHeightUnits: unit.Distance
    twistUnits: unit.Distance
    velocityUnits: unit.Velocity
    distanceUnits: unit.Distance
    diameterUnits: unit.Distance
    lengthUnits: unit.Distance
    weightUnits: unit.Weight
    temperatureUnits: unit.Temperature
    pressureUnits: unit.Pressure
    dropUnits: unit.Distance
    pathUnits: unit.Angular
    angularUnits: unit.Angular
    energyUnits: unit.Energy
    ogwUnits: unit.Weight


class TrajectoryData(NamedTuple):
    TravelledDistance: float
    Velocity: float
    Time: float
    Drop: float
    DropAdjustment: float
    Windage: float
    WindageAdjustment: float
    Energy: float
    OptimalGameWeight: float
    MachVelocity: float
