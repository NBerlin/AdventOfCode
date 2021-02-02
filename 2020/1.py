file = open("input1")
inputs = file.read().splitlines()
inputs = [int(i) for i in inputs]

def part1():
    for x in inputs:
        for y in inputs:
            if x+y == 2020:
                return x*y

def part2():
    for x in inputs:
        for y in inputs:
            for z in inputs:
                if x+y+z == 2020:
                    return x*y*z

print(part1(),part2())