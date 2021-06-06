package weight

type Kilo float64

type Pond float64

func KiloToPond(k Kilo) Pond {
	return Pond(k * 2.20)
}