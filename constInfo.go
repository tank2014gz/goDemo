
package main

import(
    "fmt"
)

func main(){
  const limit=512
  fmt.Println(limit)
  const top uint16=1421
  fmt.Println(top)

  const Pi float64=3.1415926
   fmt.Println(Pi)

   const x,y int =3,6

   fmt.Println(x+y)

   const(
    a=iota
    b=iota
    c=iota
    )
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   const d=iota
   fmt.Println(d)

}
