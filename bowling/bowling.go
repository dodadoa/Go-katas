package bowling

const (
  STRIKE = 10
  SPARE = 10
  JUMP_FRAME_STRIKE = 1
  JUMP_FRAME_SPARE = 2
  JUMP_FRAME_NORMAL = 2
)

func frame(line []int) []int {

  compensatedLine := append(line, 0, 0)

  framing := []int{}

  index := 0
  for index < len(line) {
    if compensatedLine[index] == STRIKE {
      framing = append(
        framing,
        STRIKE + compensatedLine[index + 1] + compensatedLine[index + 2])
      index = index + JUMP_FRAME_STRIKE
    } else if compensatedLine[index] + compensatedLine[index + 1] == SPARE {
      framing = append(
        framing,
        SPARE + compensatedLine[index + 2])
      index = index + JUMP_FRAME_SPARE
    } else {
      framing = append(
        framing,
        compensatedLine[index] + compensatedLine[index+1])
      index = index + JUMP_FRAME_NORMAL
    }
  }

  return framing
}

func score(line []int) int {
  framedLine := frame(line)
  total := 0
  for index := 0; index < 10; index++ {
    total += framedLine[index]
  }
  return total
}
