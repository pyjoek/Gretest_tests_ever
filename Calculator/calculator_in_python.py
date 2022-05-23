def add():
    num1 = float(input("Enter the first number: "))
    num2 = float(input("Enter the first number: "))
    sum = num1 + num2
    print(sum)

def sub():
    num1 = float(input("Enter the first number: "))
    num2 = float(input("Enter the first number: "))
    sum = num1 - num2
    print(sum)

def multy():
    num1 = float(input("Enter the first number: "))
    num2 = float(input("Enter the first number: "))
    sum = num1 * num2
    print(sum)

def divide():
    num1 = float(input("Enter the first number: "))
    num2 = float(input("Enter the first number: "))
    sum = num1 / num2
    print(sum)

def loop():
    de = input("Do you wanna run it again: ").lower()
    if de == "yes" or de == "y":
        main()
    else:
        print("GoodBye !!!!!!")

def main():
    print("##### Welcome to Python Calculator    #####")
    option = input("Enter the oeration you wanna use [add | sub | multy | divide]: ")
    if option == "add":
        add()
        loop()
    elif option == "sub":
        sub()
        loop()
    elif option == "multy":
        multy()
        loop()
    elif option == "divide":
        divide()
        loop()
    else:
        print("There's no that option in the list")

main()