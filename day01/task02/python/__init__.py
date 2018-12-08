liste1 = [0]
summe = 0
s = "numbers2.txt"
fobj = open("numbers2.txt")
for line in fobj:
    print(line)
    i = int(line)
    summe = i + summe
    print(summe)
    liste1.extend(summe)
    for i in range:
        if summe == liste1:
            print(liste1)
            print("ja")
fobj.close()
