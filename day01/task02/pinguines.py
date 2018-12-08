summe = 0
sums = []
nums = [+3, +3, +4, -2, -4]
num = 0
i = num
s = "numbers2.txt"
#fobj = open("numbers2.txt")
#for line in fobj:
#    nums.append(int(line))
#fobj.close()
while True:
    for num in nums:
        print(num)
    #    i = int(line)
        summe = summe + num
        print(summe)
    #    nums.append(num)
        for elem in sums:
            if elem == summe:
                print(summe)
                print("ja")
                exit(0)
        sums.append(summe)
    #fobj.close()
