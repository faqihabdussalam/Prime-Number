n = int(input())

if n > 1:
    for i in range(2,n):
        if (n % i) == 0:
            print("bukan prima")
            break
    else:
            print("prima")
else:
    print("bukan prima")
