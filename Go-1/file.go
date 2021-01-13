package main  
import (  
   "os"  
   "log"  
   "io/ioutil"  
   "fmt"  
)  
func main() {  
   file, err := os.Create("hi.go")  
   if err != nil {  
      log.Fatal(err)  
   }  
   file.WriteString("Hello World")  
   file.Close()  
   stream, err:= ioutil.ReadFile("hello.txt")  
   if err != nil {  
      log.Fatal(err)  
   }  
   readString := string(stream)  
   fmt.Println(readString)  
}  

