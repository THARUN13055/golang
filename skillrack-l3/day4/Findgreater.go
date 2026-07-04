package main
import "fmt"

func main(){
	var T int
	fmt.Scan(&T)
	
	for i:=0;i<T;i++{
	    var A,B,C int
	    fmt.Scan(&A,&B,&C)
	    
	    if A+B > 2*C {
	        fmt.Println("YES")
	    }else {
	        fmt.Println("NO")
	    }
	}
	
}


// Input Format
// The first line of input will contain a single integer 
// T
// T, denoting the number of test cases.
// Each test case consists of 
// 3
// 3 integers 
// A
// ,
// B
// ,
// A,B, and 
// C
// C.
// Output Format
// For each test case, output YES if average of 
// A
// A and 
// B
// B is strictly greater than 
// C
// C, NO otherwise.

// You may print each character of the string in uppercase or lowercase (for example, the strings YeS, yEs, yes and YES will all be treated as identical).

// Constraints
// 1
// ≤
// T
// ≤
// 1000
// 1≤T≤1000
// 1
// ≤
// A
// ,
// B
// ,
// C
// ≤
// 10
// 1≤A,B,C≤10
// Sample 1:
// Input
// Output
// 5
// 5 9 6
// 5 8 6
// 5 7 6
// 4 9 8
// 3 7 2
// YES
// YES
// NO
// NO
// YES