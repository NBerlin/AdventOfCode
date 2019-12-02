file = open("day2")
og = file.readline().strip().split(",")
input = [int(i) for i in og]

'''
#part1
input = [int(i) for i in og]
input[1] = 12
input[2] = 2

i = 0
while input[i] != 99:
    op = input[i]
    firstValue = input[i+1]
    secondValue = input[i+2]
    target = input[i+3]

    if op == 1:
        input[target] = input[firstValue] + input[secondValue]
    elif op == 2:
        input[target] = input[firstValue] * input[secondValue]

    i+=4
'''
#part2
for x in range(99):
    for y in range(99):
        input = [int(i) for i in og]
        input[1] = x
        input[2] = y
        
        i = 0
        while input[i] != 99:
            op = input[i]
            firstValue = input[i+1]
            secondValue = input[i+2]
            target = input[i+3]

            if op == 1:
                input[target] = input[firstValue] + input[secondValue]
            elif op == 2:
                input[target] = input[firstValue] * input[secondValue]

            i+=4
        if input[0] == 19690720:
            print(100*x+y)