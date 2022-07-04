// First Declare package
package main

//import package
import "fmt"

//Create function
func main () {
  
  //Variabes and strings to store data for comic book
  var publisher, writer, artist, title string
  var year, pageNumber uint
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  //Description of comics
  fmt.Println(title, "written by", writer, "drawn by", artist)
  fmt.Println("Published by", publisher, "with a total number of", pageNumber, "published" , year, "rated" , grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0


  fmt.Println(title, "written by", writer, "drawn by", artist)
  fmt.Println("Published by", publisher, "with a total number of", pageNumber, "published" , year, "rated" , grade)
  

 


}
