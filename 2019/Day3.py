a,b = open("day3").read().splitlines()
a,b = [x.split(',') for x in [a,b]]

dx = {'L':-1, 'R':1, 'U':0, 'D':0}
dy = {'L':0, 'R':0, 'U':1, 'D':-1}



def setUpValues(listObject):
    x,y = 0,0
    values = {}
    traveled = 0
    for item in listObject:
        direction = item[:1]
        length = int(item[1:])
        for i in range(length):
            traveled +=1
            x += dx[direction]
            y += dy[direction]
            if (x,y) not in values:
                values[(x,y)] = traveled
    return values

pa, pb = setUpValues(a), setUpValues(b)
bothPoints = set(pa.keys() & pb.keys())
print(min([abs(x)+abs(y) for (x,y) in bothPoints]))
print(min([pa[k] + pb[k] for k in bothPoints]))