package marsRover

type Plane struct {
  upper_right_x int
  upper_right_y int
}

type Rover struct {
  direction string
  position_x int
  position_y int
}

func (rover *Rover) goNextMove() {
  switch rover.direction {
  case "N": rover.position_y = rover.position_y + 1
  case "W": rover.position_x = rover.position_x - 1
  case "S": rover.position_y = rover.position_y - 1
  case "E": rover.position_x = rover.position_x + 1
  }
}

func (rover *Rover) turnDirectionNextMove(direction string) {
  switch direction {
  case "L": switch rover.direction {
            case "N": rover.direction = "W"
            case "W": rover.direction = "S"
            case "S": rover.direction = "E"
            case "E": rover.direction = "N"
            }
  case "R": switch rover.direction {
            case "N": rover.direction = "E"
            case "W": rover.direction = "N"
            case "S": rover.direction = "W"
            case "E": rover.direction = "S"
            }
  }
}

func (rover *Rover) move(moveCommands string) {
  for index := 0; index < len(moveCommands); index++ {
    switch string(moveCommands[index]) {
    case "M": rover.goNextMove()
    case "L": rover.turnDirectionNextMove("L")
    case "R": rover.turnDirectionNextMove("R")
    }
  }
}
