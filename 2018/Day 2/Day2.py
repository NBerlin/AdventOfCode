puzzle = open("input")
numberOfTwo = 0
numberOfThree = 0
for line in puzzle:
    characters = list(line)
    print(characters)
    uniqueCharacters = set(characters)
    two = True
    three = True
    for uniqueCharacter in uniqueCharacters:
        value = characters.count(uniqueCharacter)
        
        if value==2 and two:
            numberOfTwo+=1
            two=False
        if value == 3 and three:
            numberOfThree+=1
            three=False
print(numberOfThree*numberOfTwo)

