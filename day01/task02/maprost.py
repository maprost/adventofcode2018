summe = 0
arr = []
fobj = open("input_70357.txt")

nums = []
for line in fobj:
    nums.append(int(line))
fobj.close()

while True:
    for num in nums:
        summe = num + summe

        # check if sum is in arr
        for elem in arr:
            if elem == summe:
                print("first reaches " + str(summe) + " twice")
                exit(0)

        arr.append(summe)


