package main

import (
"errors"
//"fmt"
)

func divide(a int,b int)(num int,err error){
  if b==0 {
      err:=errors.New("bei chu shu bu neng wei 0!")
      return a/b err
  }
  return a/b,nil
}

func main(){
 divide(4,2)
//  fmt.printf(s)

}
