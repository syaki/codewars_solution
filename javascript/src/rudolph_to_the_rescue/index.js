var maze = [
    [
        { x: 0, y: 0 },
        { x: 1, y: 0 },
    ],
    [
        { x: 0, y: 1 },
        { x: 1, y: 1 },
    ],
]
maze.cell = function (x, y) {
    return this[y][x]
}
maze[0][0].south = maze.cell(0, 1)
maze[0][0].east = maze.cell(1, 0)
maze[0][1].west = maze.cell(0, 0)
maze[1][0].north = maze.cell(0, 0)
maze[1][0].east = maze.cell(1, 1)
maze[1][1].west = maze.cell(0, 1)

console.log(findPath(maze[0][0], maze[1][1]))

function findPath(start, end, prev = null) {
    // I wish Santa could do some magic here...
    if (start === end) {
        return end
    }

    for (let d of ['east', 'south', 'west', 'north']) {
        if (start[d] && (!prev || start[d] !== prev)) {
            const temp = findPath(start[d], end, start)
            if (temp) {
                return [start].concat(temp)
            }
        }
    }
    return false
}
