package main
import (
	"fmt"
	// "strconv"
)


func main(){
	

	x:=[]string{"a","b","c"};
	y:=[]string{"1","2","3","4","5"}
	
	fmt.Println(Join(x,y))
	
	fmt.Println(fibonacci(20))
}

func Join(arr1 ,arr2 []string)[]string{
	// var arr3 []string 
	// var arr4 []string
	var arr5 []string

	// if(len(arr1)>len(arr2)){
	// 	arr4,arr3=arr2,arr1
	
	// }else{
	// 	arr4,arr3=arr1,arr2
	// }
	
	for i,_ := range getSmallerArray(arr1,arr2){
		
		arr5=append(arr5,arr1[i],arr2[i])
		
	}
	arr5=append(arr5,getLongerArray(arr1,arr2)[len(getSmallerArray(arr1,arr2)):]...)
	return arr5


	 
}


func getLongerArray(arr1,arr2 []string)[]string{
	if(len(arr1)>len(arr2)){
		return arr1;
	}else{
		return arr2
	}
}
func getSmallerArray(arr1,arr2 []string)[]string{
	if(len(arr1)>len(arr2)){
		return arr2
	}else{
		return arr1
	}
}

// func JoinSecond(arr1 ,arr2 []string)[]string{
// 	var arr3 []string 
// 	var arr4 []string
	
// 	var arr5 []string

// 	if(len(arr1)>len(arr2)){
// 		arr4,arr3=arr2,arr1
// 	}else{
// 		arr4,arr3=arr1,arr2
// 	}
	
// 	for i,_ := range arr4{
// 		a:=arr4[:i]
// 		b:=arr5[i:]
// 		if(i%2==0){
// 			arr4=append(arr4[:i],arr3[i])
// 		}
		
// 		// arr5=append(arr5,arr1[i],arr2[i])
		
// 	}
// 	// arr5=append(arr5,arr3[len(arr4):]...)
// 	arr4=append(arr3,arr3[len(arr4):]...)
// 	return arr4


	 
// }

func fibonacci(num int)[]int{
	var x []int=[]int {}
	x=append(x,0)
	x=append(x,1)


	for i:=2; i < num; i++ {
		// x[i]=x[i-1]+x[i-2]
		x=append(x,x[i-1]+x[i-2])
	} 

   
	

  
	return x;
}