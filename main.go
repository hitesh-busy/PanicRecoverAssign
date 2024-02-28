package main

import (
	"fmt"
)

/*
Recover is useful only when called inside deferred functions. Executing a call to recover inside a deferred function stops the panicking sequence by restoring normal execution and retrieves the error message passed to the panic function call. If recover is called outside the deferred function, it will not stop a panicking sequence.

Line 29 calls recover() which returns the value passed to panic function call. Here we are just printing the value returned by recover in line no. 18. recovery() is being deferred in line no. 29 of the divide function.

When divide panics, the deferred function recovery() will be called which uses recover() to stop the panicking sequence.
If we recover from a panic, we lose the stack trace about the panic. Even in the program above after recovery, we lost the stack trace.
Recover works only when it is called from the same goroutine which is panicking. It’s not possible to recover from a panic that has happened in a different goroutine.
*/

func recovery() {
	if r := recover(); r != nil {

		
		//calling the funciton again with normal value
		//as well know when panic occurs rest of the lines are not executed
		fmt.Println("Calling the function again normally")
		//debug.PrintStack() //used to print call stack  in case of recovering

		divide(5, 1)
	}
}

func divide(a int, b int) {
	//using defer to call recovery(). because to only want to handle the panic after it occurs, so deferring its execution.
	defer fmt.Println("Inside Divide Function")
	//jese hi panic wali jagah hit karta h code, recovery ke baad sabkuch uske baad ka skip hojata h except defer statements
	defer recovery()
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Println("Statement After panic statement")

}

func main() {
	for {
		var a, b int
		fmt.Println("Enter no  or -1 to quit")
		fmt.Scanln(&a)
		if a == -1 {
			fmt.Println("Quiting...")
			break
		}
		fmt.Println("Enter no  or -1 to quit")
		fmt.Scanln(&b)
		if b == -1 {
			fmt.Println("Quiting...")
			break
		}
		divide(a, b)
	}
	fmt.Println("normally returned from main")
}
