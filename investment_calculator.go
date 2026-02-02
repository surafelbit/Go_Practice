package main
import "fmt"
// import "math"
func main(){
// var investment float64	= 1000
// var revenue float64
// var expenses float64
// var tax_rate float64


// fmt.Print("Enter Your Revenue: ")
// fmt.Scan(&revenue)

var num1, num2, num3 float64= valueGetter()
//featureValues(num1,num2,num3)
// fmt.Print("Enter The Expenses: ")

// fmt.Scan(&expenses)
// fmt.Print("Enter The Tax Rate: ")
// fmt.Scan(&tax_rate)
// earningBeforeTax := revenue-expenses
// earningAfterTax := revenue - earningBeforeTax*(tax_rate/100)
//  fmt.Print("Your earning before tax ")
//  fmt.Println(earningBeforeTax);
//  fmt.Print("Your earning after tax ")
//   fmt.Println(earningAfterTax);
// fmt.Println("The ratio")
// fmt.Println(earningAfterTax/earningBeforeTax)
// outPutText(earningAfterTax/earningBeforeTax)
// fmt.Printf("YOUR RATIO WITH %f", earningAfterTax)

// fmt.Println("this is the thing dawg")
 columbia, wow:=featureValues(num1,num2,num3)
fmt.Printf("This is your earning after tax:%.1f\n", columbia)
fmt.Printf("This is your ration:%.1f ",wow)

}

func outPutText(text float64){
  fmt.Print(text)
}
func featureValues (revenue , expenses, tax_rate  float64)(float64,float64){
  earningBeforeTax := revenue-expenses

  earningAfterTax := revenue - earningBeforeTax*(tax_rate/100)
 ratio:= earningAfterTax/earningBeforeTax
return earningAfterTax,ratio
}
func valueGetter()(float64,float64, float64){
  var num1, num2, num3 float64
  fmt.Print("Enter Revenue: ")
  fmt.Scan(&num1)
  fmt.Print("Enter Expenses: ")
  fmt.Scan(&num2)
  fmt.Print("Enter Tax Value: ")
  fmt.Scan(&num3)
return num1, num2, num3
}
