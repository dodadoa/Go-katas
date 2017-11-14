package marsRover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_Rover_goNextMove(t *testing.T) {
  rover := Rover{
    direction: "N",
    position_x: 0,
    position_y: 0,
  }
  shouldBeNextMoveRover := Rover{
    direction: "N",
    position_x: 0,
    position_y: 1,
  }
  rover.goNextMove()
	assert.Equal(t, shouldBeNextMoveRover, rover)
}


func Test_Rover_turnDirectionNextMove(t *testing.T) {
  rover := Rover{
    direction: "N",
    position_x: 0,
    position_y: 0,
  }
  turnLeftRover := Rover{
    direction: "W",
    position_x: 0,
    position_y: 0,
  }
  turnRightRover := Rover{
    direction: "E",
    position_x: 0,
    position_y: 0,
  }
  rover.turnDirectionNextMove("L")
  assert.Equal(t, turnLeftRover, rover)
  rover.turnDirectionNextMove("R")
  rover.turnDirectionNextMove("R")
  assert.Equal(t, turnRightRover, rover)
}


func Test_Rover_command_move(t *testing.T){
	rover1 := Rover{
		direction: "N",
		position_x: 0,
		position_y: 0,
	}
	finalRover1 := Rover{
		direction: "N",
		position_x: 0,
		position_y: 3,
	}
	rover1.move("MMM")
	assert.Equal(t, finalRover1, rover1)

	rover2 := Rover{
		direction: "N",
		position_x: 1,
		position_y: 2,
	}
	finalRover2 := Rover{
		direction: "N",
		position_x: 1,
		position_y: 3,
	}
	rover2.move("LMLMLMLMM")
	assert.Equal(t, finalRover2, rover2)

	rover3 := Rover{
		direction: "E",
		position_x: 3,
		position_y: 3,
	}
	finalRover3 := Rover{
		direction: "E",
		position_x: 5,
		position_y: 1,
	}
	rover3.move("MMRMMRMRRM")
	assert.Equal(t, finalRover3, rover3)
}
