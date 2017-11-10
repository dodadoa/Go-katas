package bowling

func frame(line []int) []int {

  const STRIKE int = 10
  const SPARE int = 10
  const JUMP_FRAME_STRIKE int = 1
  const JUMP_FRAME_SPARE int = 2
  const JUMP_FRAME_NORMAL int = 2

  compensatedLine := append(line, 0, 0)

  framing := []int{}

  index := 0
  for index < len(line) {
    if compensatedLine[index] == STRIKE {
      framing = append(framing, STRIKE + compensatedLine[index + 1] + compensatedLine[index + 2])
      index = index + JUMP_FRAME_STRIKE
    } else if compensatedLine[index] + compensatedLine[index + 1] == SPARE {
      framing = append(framing, SPARE + compensatedLine[index + 2])
      index = index + JUMP_FRAME_SPARE
    } else {
      framing = append(framing, compensatedLine[index] + compensatedLine[index+1])
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
