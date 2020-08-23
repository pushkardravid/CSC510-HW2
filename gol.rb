def getAliveNeighborsCount(row, col, grid)
	rows = grid.length
	columns = grid[0].length
	neighbors = Array.new()
	if row>0
		neighbors << grid[row - 1][col]
	end
	if row<rows-1
		neighbors << grid[row + 1][col]
	end
	if col>0
		neighbors << grid[row][col - 1]
	end
	if col<columns-1
		neighbors << grid[row][col + 1]
	end
	if row>0 and col>0
		neighbors << grid[row - 1][col - 1]
	end
	if row<rows-1 and col<columns - 1
		neighbors << grid[row + 1][col + 1]
	end
	if row>0 and col<columns - 1
		neighbors << grid[row - 1][col + 1]
	end
	if row<rows-1 and col>0
		neighbors << grid[row + 1][col - 1]
	end
	return neighbors.count(1)
end

def printGrid(grid, generation)
	for row in grid do
		puts("\t" + row.map{ |elem| elem == 1 ? '*' : '' 	}.join(' '))
	end
	puts("\t Generation: " + generation.to_s)
	sleep 2
	Gem.win_platform? ? (system "cls") : (system "clear")
end

#read input from user
puts("Enter space separated values for number of rows, columns, and generations to simulate. Example: 10 10 20")
rows, columns, generations = gets.split().map{ |num| num.to_i}

#create grid
grid = Array.new(rows) { Array.new(columns) { 0 } }

#initialize random cells as alive
(0..rows-1).each do |row|
	(0..columns-1).each do |col|
		if row*col > (rand 0..(rows*columns))
			grid[row][col] = 1
		end 
	end
end

#simulate generations for grid
(0..generations).each do |generation|
	printGrid(grid, generation)
	next_grid = Array.new(rows) { Array.new(columns) { 0 } }
	(0..rows-1).each do |row|
		(0..columns-1).each do |col|
			alive_neighbors = getAliveNeighborsCount(row, col, grid)
			alive = grid[row][col] == 1? true : false
			if alive
				if alive_neighbors == 2 || alive_neighbors == 3
					next_grid[row][col] = 1
				end
			else
				if alive_neighbors == 3
					next_grid[row][col] = 1
				end
			end
		end
	end
	grid = next_grid
end
