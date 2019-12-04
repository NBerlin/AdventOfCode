from collections import Counter
correctNumbers = []
'''
for i in range(138307,654504):
  iList = [int(x) for x in str(i)]
  twoInARow = False
  allow = True
  for k in range(1,6):
    if iList[k] < iList[k-1]:
      allow = False
    if iList[k] == iList[k-1]: 
      twoInARow = True
  
  if twoInARow and allow:
    correctNumbers.append(iList)
'''

for i in range(138307,654504):
  iList = [int(x) for x in str(i)]
  twoInARow = False
  allow = True
  for k in range(0,5):
    if iList[k] > iList[k+1]:
      allow = False
      break
    if iList[k] == iList[k+1]: 
      if ((k-1>-1 and iList[k-1] != iList[k]) or  k-1 == -1) and ((k+2<6 and iList[k+1] != iList[k+2]) or k+2==6):
        twoInARow = True
  
  if twoInARow and allow:
    correctNumbers.append(iList)

print(len(correctNumbers))
