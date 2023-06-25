package main

import "fmt"

/*
  Calculates the number of rabbits after a given number of years.

  Args:
    initial_population: The initial population of rabbits.
    number_of_years: The number of years.

  Returns:
    The number of rabbits after the given number of years.
*/

func main() {
   var years int
   fmt.Scanln(&years)
   r := 7
   // Calculate the rabbit population
   for i:=0;i<years;i++{
     r = r*2
   }
   fmt.Println(s)
}
