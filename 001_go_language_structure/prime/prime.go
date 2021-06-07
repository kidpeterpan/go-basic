package prime

import (
	"fmt"
	"math"
)

var notPrime [1000000]bool

// init func จะทำการ initialize ให้เราอัตโนมัติ (แค่ครั้งเดียว) เมื่อมีการเรียกใช้ package
// ต่อ ให้เรา access local variable ของ function นั้นหลายครั้งก็ตาม
func init() {
	fmt.Println("initialize in prime package")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrime)))))
	for i:=2;i<sqrtLen;i++ {
		if notPrime[i] {
			continue
		}
		notPrime[i] = false
		for j := i*2; j < len(notPrime); j+=i {
			notPrime[j] = true
		}
	}
}

func IsPrime(x int) bool {
	return !notPrime[x]
}
