package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  //variable that keeps track of whether or not our heist is successful. Declare the boolean variable isHeistOn with a value of true.
  isHeistOn := true
  eludedGuards := rand.Intn(100)

// Conditional Statements
  if eludedGuards >= 50 {fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else if !isHeistOn  {fmt.Println("Plan a better disguise next time?")}

//create a variable openedVault and assign it a random number between 0 and 99 using rand.Intn(100).
  openedVault := rand.Intn(100) 

  if isHeistOn && openedVault >= 70 {fmt.Println("Access granted, Grab and GO!")
  } else if !isHeistOn {fmt.Println("Access Denied, The vault cant be open")}

//simulating different senerios
  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely{
      case 0:
        isHeistOn = false 
        fmt.Println("oops,You have been caught")
      case 1:
        isHeistOn = false
        fmt.Println("wow, tripped the sensors, Looks like some prison time")
      case 2:
        isHeistOn = false
        fmt.Println("Turns out vault doors don't open from the inside...")
      case 3:
        isHeistOn = false
        fmt.Println("Looks like this fingerprint scanner won’t accept any fingerprint") 
      default:
        fmt.Println("Start the getaway car!") }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("Looks like we made away with" , amtStolen)
  }

fmt.Println("isHeistOn is currently:", isHeistOn)
}
