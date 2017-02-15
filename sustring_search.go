package main

import "fmt"


func main() {
	var haystack string = "Felix Hsieh" 
	var needle string = "Hs" 
	var found bool = true
	
	for i:=0; i< len(haystack)-len(needle)+1; i++ {
		for j:=0; j<len(needle); j++{
			 if(haystack[i+j]!=needle[j]) { break      }
			 if((j+1)==len(needle))       { found=true 
			 }
		} 
	}
	fmt.Println(found);

}
