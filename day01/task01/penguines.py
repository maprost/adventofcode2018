summe = 0
fobj = open("input_435.txt")
for line in fobj:
    i = int(line)
    summe = i + summe

print(summe)
fobj.close()

