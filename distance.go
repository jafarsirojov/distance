package distance

import (
	"math"
)

func DistanceTwoCoordinatesInMeters(longitude1, latitude1, longitude2, latitude2 float64) (meter uint64) {

	eQuatorialEarthRadius := 6378.1370
	d2r := (math.Pi / 180)

	dLong := (longitude2 - longitude1) * d2r
	dLat := (latitude2 - latitude1) * d2r

	a := math.Pow(math.Sin(dLat/2.0), 2.0) + math.Cos(latitude1*d2r)*math.Cos(latitude2*d2r)*math.Pow(math.Sin(dLong/2.0), 2.0)

	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	d := eQuatorialEarthRadius * c

	meter = uint64(d * 1000)
	return meter
}

func DistanceTwoCoordinatesInKilometers(longitude1, latitude1, longitude2, latitude2 float64) (kilometer uint64) {

	eQuatorialEarthRadius := 6378.1370
	d2r := (math.Pi / 180)

	dLong := (longitude2 - longitude1) * d2r
	dLat := (latitude2 - latitude1) * d2r

	a := math.Pow(math.Sin(dLat/2.0), 2.0) + math.Cos(latitude1*d2r)*math.Cos(latitude2*d2r)*math.Pow(math.Sin(dLong/2.0), 2.0)

	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	d := eQuatorialEarthRadius * c

	kilometer = uint64(d)
	return kilometer
}
