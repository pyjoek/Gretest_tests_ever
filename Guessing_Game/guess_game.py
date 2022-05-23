import random

def game():
    a = int(input("Enter lowest interval: "))
    b = int(input("Enter lowest interval: "))
    comp = random.randint(a, b)
    guess = int(input("Enter your guess: "))
    if guess == comp:
        print("YOu won. The value was ",comp)
        game()
    else:
        print("You loose. The value was ",comp)

game()