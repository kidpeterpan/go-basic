**GOPATH**

- ถ้าเรา `go env` ดูจะเห็นว่า `$GOPATH` ถูก set ไว้ที่ไหน
- เข้าไปจะเจอ directories หลัก 3 อัน 
    + bin เอาไว้เก็บ executable bin
    + pkg เก็บ install object เช่น (พวกที่เรา `go get` มา มันจะทำการ install package ที่เราต้องการใช้ให้ด้วย
      โดยส่วนที่เป็น source code จะไปอยู่ที่ src)
    + src เก็บ source code ต่างๆ ไว้ที่นี่ ทั้ง `built-in` และ `go tools` หรือ `go get` ต่างๆ  
    เวลาเรา import pkg ก็จะมาหาที่นี่  
      
*ตอนนี้ใช้ `go mod` แล้วจะไม่สนใจมันแล้วก็ได้