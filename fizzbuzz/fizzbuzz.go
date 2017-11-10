package fizzbuzz

import (
  "strconv"
  "strings"
)

func determineFizzBuzz(countNumber int) string {

  numberToString := strconv.Itoa(countNumber)

  isFizz := (countNumber % 3 == 0) || (strings.Contains(numberToString, "3"))
  isBuzz := (countNumber % 5 == 0) || (strings.Contains(numberToString, "5"))
  isFizzBuzz := isFizz && isBuzz

  if isFizzBuzz {
    return "FizzBuzz"
  } else if isFizz {
    return "Fizz"
  } else if isBuzz {
    return "Buzz"
  } else {
    return numberToString
  }
}

func FizzBuzz(countNumber int) string {
  printFormat := determineFizzBuzz(countNumber)
  return printFormat
}
