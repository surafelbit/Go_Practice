package main
import "fmt"
import "os"
import "errors"
import "strconv"
// import "math"

func main(){
var accountBalance , err = readFile()
if err !=nil{
  fmt.Println("Error")
  fmt.Println(err)
  fmt.Println("_____________")
  return
}
fmt.Println("Welcome to Go bank")
fmt.Println("Enter Your Choice Number: ")
fmt.Println("Press 1 to check balance: ")
fmt.Println("Press 2 to deposit to bank: ")
fmt.Println("Press 3 to withdraw balance: ")
fmt.Println("Press 4 to exit: ")
var Choice int 
fmt.Print("Enter Your Choice: ")
fmt.Scan(&Choice)
if Choice == 1 {
  fmt.Print("Your Balance Is: ", accountBalance)
} else if (Choice==2){
  var deposit float64
  fmt.Print("How much do you want to deposit: ")
  fmt.Scan(&deposit)
  accountBalance = accountBalance + deposit
    writeToFile(accountBalance)

  fmt.Print("Your New Balance is ", accountBalance)
} else if (Choice==3){
  var withdraw float64
  fmt.Print("How much do you want to deposit: ")
  fmt.Scan(&withdraw)
  accountBalance = accountBalance - withdraw
  writeToFile(accountBalance)
  fmt.Print("Your New Balance is ", accountBalance)
} else{
  fmt.Print("Goodbye")
}


var num1, num2, num3 float64= valueGetter()

 columbia, wow:=featureValues(num1,num2,num3)
fmt.Printf("This is your earning after tax:%.1f\n", columbia)//////////

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
//
func writeToFile(account float64){
  accountAmount := fmt.Sprint(account)
  os.WriteFile("balance.txt", []byte(accountAmount), 0644)
}
func readFile() (float64,error){
 data,err:= os.ReadFile("balance.txt")
 if err != nil{
  return 100 , errors.New("Failed to read file")
 }
 balanceText := string(data)
 balance,err :=strconv.ParseFloat(balanceText, 64)
 if err != nil{
  return 100 , errors.New("Failed to convert ")
 }
return balance,nil
}