package bowling


func score(line [4]int) int {
  total := 0
  for _,item := range line {
    total += item
  }
  return total
}
