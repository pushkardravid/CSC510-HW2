function getAliveNeighbours(row, column, arr)
  rows, columns = size(arr)
  neighbours = Int64[]
  if row > 1
      push!(neighbours, arr[row - 1, column])
  end
  if row<rows
      push!(neighbours, arr[row + 1, column])
  end
  if column>1
      push!(neighbours, arr[row, column - 1])
  end
  if column<columns
      push!(neighbours, arr[row,column + 1])
  end
  if row>1 && column>1
      push!(neighbours, arr[row - 1,column - 1])
  end
  if row<rows && column<columns-1
      push!(neighbours, arr[row + 1,column + 1])
  end
  if row>1 && column<columns-1
    push!(neighbours, arr[row - 1,column + 1])
  end
  if row<rows && column>1
    push!(neighbours, arr[row + 1,column - 1])
  end
  
  return count(x->x==1, neighbours)

end


function getNextGeneration(currentGeneration)
  xsize, ysize = size(currentGeneration)
  newGeneration = zeros(Int8, xsize, ysize)
  for i = 1:xsize
    for j = 1:ysize
      alive = getAliveNeighbours(i,j,currentGeneration)
      if alive > 0
          if alive == 2 || alive == 3
            newGeneration[i][j] = 1
          end
      else
          if alive == 3
					  newGeneration[i][j] = 1
				  end
      end
    end
  end

  return newGeneration
end

function prettyDisplay(life, i, generations)
    life = replace(life, 1 =>"*", 0 =>" ")
    println("Generation $i")
    for row in eachrow(life)
      println(*(row...))
      sleep(0.1)
    end
    if i != generations
      Base.run(`clear`)
    end
end

function main()
    println("Welcome to game of life. Before we start, please provide few parameters")
    println("Size of the square matrix: ")
    size = parse(UInt8, readline())
    println("How many generations do you want: ")
    generations = parse(UInt8, readline())
    initialLife = rand([0, 1], size, size)
    alive = count(x->x==1, initialLife) / 2
    initialLife = replace(initialLife, 1 =>0, count = trunc(Int, alive))

    for i = 1:generations
        prettyDisplay(initialLife, i, generations)
    end
end

main()
