package reference

type StopInfo struct { 
	Longitude float64
	Latitude  float64
	Direction string
	TimeStamp int // in milliseconds
}

const (
	StopBtime = 251973
	StopAtime = 457254 //in milliseconds
)

var ReferenceMap = []StopInfo{ //forward is going from A -> B, reverse B -> A
	{Longitude: -79.9219256, Latitude: 43.2601414, Direction: "forward", TimeStamp: 0}, //Point A
	{Longitude: -79.9209266, Latitude: 43.2601393, Direction: "forward", TimeStamp: 22019},
	{Longitude: -79.9190291, Latitude: 43.2597076, Direction: "forward", TimeStamp: 71957},
	{Longitude: -79.9190397, Latitude: 43.2585466, Direction: "forward", TimeStamp: 91927},
	{Longitude: -79.9180279, Latitude: 43.2578761, Direction: "forward", TimeStamp: 111765},
	{Longitude: -79.9159219, Latitude: 43.2590493, Direction: "forward", TimeStamp: 141781},
	{Longitude: -79.9158942, Latitude: 43.2607345, Direction: "forward", TimeStamp: 161717},
	{Longitude: -79.916043, Latitude: 43.261486, Direction: "forward", TimeStamp: 171701},
	{Longitude: -79.9165057, Latitude: 43.262646, Direction: "forward", TimeStamp: 201761},
	{Longitude: -79.9163497, Latitude: 43.2634842, Direction: "forward", TimeStamp: 222013}, 
	{Longitude: -79.9166429, Latitude: 43.2632088, Direction: "forward", TimeStamp: 251973 }, //Point B
	{Longitude: -79.9166429, Latitude: 43.2632088, Direction: "reverse", TimeStamp: 251973 }, //Point B
	{Longitude: -79.9168373, Latitude: 43.2623833, Direction: "reverse", TimeStamp: 291853 },
	{Longitude: -79.9158826, Latitude: 43.2614362, Direction: "reverse", TimeStamp: 322038 },
	{Longitude: -79.9159878, Latitude: 43.2602883, Direction: "reverse", TimeStamp: 331970 },
	{Longitude: -79.915909, Latitude: 43.2589812, Direction: "reverse", TimeStamp: 351936},
	{Longitude: -79.9173126, Latitude: 43.2580618, Direction: "reverse", TimeStamp: 372088 },
	{Longitude: -79.9183927, Latitude: 43.2583166, Direction: "reverse", TimeStamp: 381971 },
	{Longitude: -79.9190962, Latitude: 43.259267, Direction: "reverse", TimeStamp: 402035 },
	{Longitude: -79.9193478, Latitude: 43.2600841, Direction: "reverse", TimeStamp: 422045},
	{Longitude: -79.9210478, Latitude: 43.2600492, Direction: "reverse", TimeStamp: 445502},
	{Longitude: -79.9219256 , Latitude: 43.2601414, Direction: "reverse", TimeStamp: 457254 }, //Point A (back to start)
}
