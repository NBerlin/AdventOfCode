#part1
file = open("input1")
inputs = file.read().splitlines()
inputs = [int(i) for i in inputs]
for x in inputs:
    for y in inputs:
        if x+y == 2020:
            print(x*y)

#part2
for x in inputs:
    for y in inputs:
        for z in inputs:
            if x+y+z == 2020:
                print(x*y*z)