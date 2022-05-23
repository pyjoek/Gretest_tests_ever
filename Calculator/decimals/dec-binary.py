binary = []
n = int(input("Enter the decimal number: "))

while n >= 1:
    p = n % 2
    if p >= 0.5 and p <= 0.9:
        p = 1
    elif p >= 1.0 and p <= 1.4:
        p = 1
    binary.append(int(p))
    n = n/2

while True:
    if len(binary) == 8:
        break
    else:
        binary.append(0)
        
print(binary[::-1])