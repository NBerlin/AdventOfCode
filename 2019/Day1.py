file = open("day1")
inputs = file.read().splitlines()
inputs = [int(i) for i in inputs]


#part1
totalSum1 = 0
for number in inputs: 
    totalSum1 += (number//3)-2


#part2
totalSum2 = 0
for number in inputs:
    tempNumber = (number//3)-2
    while tempNumber >= 1:
        totalSum2 += tempNumber
        tempNumber= (tempNumber//3)-2

print(totalSum1, totalSum2)