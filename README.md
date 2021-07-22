# Golang Basic
## 1. Variable
### Khai bao mot bien
### Khai bao nhieu bien cung kieu du lieu
```
C1:
var a,b int
a = 1
b = 2
C2:
var a1,b1 int = 10, 11
C3: Type Interface
var a2, a3 = 12, 13
```
### Khai bao nhieu bien khac kieu du lieu
```
C1:
var ( name string
      address string
      age int
)
name = "Ken"
address = "Dong Lao"
age = 25
C2:
var (
      name1 string = "name1"
      address1 string = "add1"
      age int = 25
     )
C3:
var name2, address2, age2 = "name2", "address2", 25
```
### Khai bao ngan gon
```
language := "Golang"
```
==> Su dung nhieu

## Kieu du lieu
### int
```
Range: 
fmt.Println(math.MinInt32)
fmt.Println(math.MaxInt32)
e.x
Range int 8: -128 -- 127
Bits:
fmt.Println(bits.OnesCount8(math.MaxUint8))
```
### uint so nguyen duong
var myUint uint = -10 ==> sai
### Byte
Byte la kieu uint8
### Float: kieu so thuc
float32, float64
### complex: kieu so phuc
z = a + bi
```
var z1 complex64 = 10 + 1i
```