package main  
import (  
   "fmt"  
)  
type user struct {  
   name string  
   email string  
   password string  
}  
func main() {  
   u := user{name: "krishna", email: "krishnalumia535@gmail.com", password: "test1234", }  
   fmt.Println(u)  
   fmt.Println("Name is :: ", u.name)
   fmt.Println("Email is :: ", u.email)
   fmt.Println("Password is :: ", u.password)  
}  

