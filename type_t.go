package main
import (
"fmt"
)

func main(){
  x:=uint16(65000)
  y:=int16(x)
  fmt.Printf("type and value of x is %T and %d\n",x,x)//%T ge shi hua zhi ling dezuo yong shi shu chu bian liang de lei xing

  fmt.Printf("type and value of y is %T and %d\n",y,y)

  var i interface{}=99
  var s interface{}=[]string{"left","right"}
  j:=i.(int)

  fmt.Printf("type and value of j is :%T and %d\n",j,j)

  if s,ok:=s.([]string);ok{
   fmt.Printf("%T _-> %q\n",s,s)
  }
}
