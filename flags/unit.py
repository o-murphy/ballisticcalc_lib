from enum import IntFlag


class Angular(IntFlag):
    Radian = 0
    Degree = 1
    MOA = 2
    Mil = 3
    MRad = 4
    Thousand = 5
    InchesPer100Yd = 6
    CmPer100M = 7


class Distance(IntFlag):
    Inch = 10
    Foot = 11
    Yard = 12
    Mile = 13
    NauticalMile = 14
    Millimeter = 15
    Centimeter = 16
    Meter = 17
    Kilometer = 18
    Line = 19


class Energy(IntFlag):
    FootPound = 30
    Joule = 31


class Pressure(IntFlag):
    MmHg = 40
    InHg = 41
    Bar = 42
    HP = 43
    PSI = 44


class Temperature(IntFlag):
    Fahrenheit = 50
    Celsius = 51
    Kelvin = 52
    Rankin = 53


class Velocity(IntFlag):
    MPS = 60
    KMH = 61
    FPS = 62
    MPH = 63
    KT = 64


class Weight(IntFlag):
    Grain = 70
    Ounce = 71
    Gram = 72
    Pound = 73
    Kilogram = 74
    Newton = 75
