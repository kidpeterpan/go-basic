package weight

type Kilo float64

type Pond float64

func KiloToPond(k Kilo) Pond {
	return Pond(k * 2.20)
}

// เราสามารถ push package ไปบน github แล้วโหลดมาใช้ด้วย go get ก็ได้ (แชร์ให้คนอื่นด้วย)
