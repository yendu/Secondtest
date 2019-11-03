package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"strings"
	// "os"

)



func main()  {
	// guessingGame()
	// getMultiplesOfTheNumber()
	// printDiamondShape(10)
	// 

	// lala:=3.3
	// p:=&lala
	// *p=3.4
	// lala=45
	// fmt.Println(lala)
	// s:="abc"
	// b:= 4
	// // s2:=string(b)
	// const s=3.13;

	// d:=[...]int{3,4,6,6}
	//  var d [4]int =[4]int{3,4,5}
	
	// fmt.Println(s)
	// var s string ="hello yoni"
	// var st []rune=[]rune{1,3,4,5,6,7}
	// var array []int;
	// var array []int=[]int{1,3,4,5};
	
	print(2,3,45);


    // print(array[0:1])

	// array[0]=3;

	// fmt.Println(cap(array))

	
	// b:=[]int{0,1,2,3,3,5}
	// fmt.Println(se==nil)
	// Reverse(a[:])
	// fmt.Println(a)
	// // fmt.Println(s,strconv.Itoa(b))
	// fmt.Println(f)
	// fmt.Println(err)

	ages :=make(map[string]int)

	agess :=map[string]int{

	}
	
}

func print(s ...int){
	for _,v := range s{
		fmt.Println(v);
	}
}

func Reverse(s []int){
	for i,j :=0,len(s)-1;i<j; i,j=i+1,j-1 {
		s[i],s[j]=s[j],s[i]
	}
}

func getMultiplesOfTheNumber()  {

	for i := 1; i <= 12; i++ {
		fmt.Println("")
		for j := 1; j <= 12; j++ {
	
			if i*j>9{
				fmt.Printf(strconv.Itoa(i*j)+" ")
			
			}else{
				fmt.Printf(strconv.Itoa(i*j)+"  ")
			
			}
			
			
			
		}
	}
}

func printDiamondShape(n int){
	for i:=1; i<n; i++{
		// fmt.Printf()
		fmt.Println()
		// fmt.Println(i)
		fmt.Printf(strings.Repeat(" ",n-i)+strings.Repeat("* ",i)+strings.Repeat(" ",n-i))

	}
	for i:=n; i>0; i--{
		fmt.Println()
		fmt.Printf(strings.Repeat(" ",n-i)+strings.Repeat("* ",i)+strings.Repeat(" ",n-i))
	}
   
   

}
func guessingGame()  {
	var GuessedNumber int
	
	var TempoGuessedNubmer int
	rand.Seed(time.Now().UnixNano())
	SecretNumber:=rand.Intn(20)
	
	for i:=1; i<=5; i++{
		fmt.Println("Guess number")
		
		TempoGuessedNubmer=GuessedNumber
		fmt.Scanln(&GuessedNumber)

		switch{
	
		case TempoGuessedNubmer==GuessedNumber:
			i--	
			fmt.Println("try another number")
		case i==5:
			fmt.Println("You Lost")
			fmt.Println("The Secret Number was"+strconv.Itoa(SecretNumber))	
			guessingGame()
			return			
		case GuessedNumber<SecretNumber:
			fmt.Println("too small")
		case GuessedNumber>SecretNumber:

			fmt.Println("too large")
	
		case GuessedNumber==SecretNumber:
			fmt.Println("you got it")
			fmt.Println("The Secret Number was"+strconv.Itoa(SecretNumber))	
		    
			guessingGame()
			return
		
		
		

		}
		
	}
	
}
// fuc GuessingGame2(){

// }