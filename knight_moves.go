type Point struct {
    X int
    Y int
}

type Position struct {
    X       int
    Y       int
    moves   int
}

func isInside(x, y, size int) bool {
    if x >= 1 && x <= size and y >= 1 and y <= size {
        return true
    }
    return false
}

func noMoves(x, y, size int) {
    dx = [4]int{2, -2, 1, -1}
    dy = [4]int{1, 1, 2, 2}
    var queue []Position
    var visited [][]bool
    for true {
        if len(queue) > 0{
            t = queue[0]
            queue = queue[1:]
            if t.X == size - 1 {
                return t.moves
            }
            for i := 0; i < 4; i++ {
                x := t.X + dx[i]
                y := t.Y + dy[i]
                if isInside(x, y, size) and not visited[x][y]{
                    visited[x][y] = true
                    queue = append(queue, Position{x, y, t.moves + 1})
                }
            }
            } else {
                break
            }
    }
}
