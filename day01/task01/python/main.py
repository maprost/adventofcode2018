summe = 0
fobj = open("numbers.txt")
for line in fobj:
    print(line)
    i = int(line)
    summe = i + summe
    print(summe)
fobj.close()
