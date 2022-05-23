import MySQLdb

conn = MySQLdb.connect(host="localhost",user="joe",password="password",autocommit=True)
my_cursor = conn.cursor()
my_cursor.execute("create database if not exists school_managements")
my_cursor.execute("use school_managements")
my_cursor.execute("create table if not exists student(fname varchar(10) UNIQUE,mname varchar(10),lname varchar(10) UNIQUE,gender enum('female','male'),age int)")

####   login begins here ##################
fname = input("What's your first name: ")
mname = input("What's your middle name: ")
lname = input("What's your last name: ")
gender = input("What's your gender: ")
while True:
    try:
        age = int(input("How old are you: "))
        if type(age) == int:
            break
        else:
            print(" ")
    except ValueError:
        print("Invalid age input")

#######  database inserting starts here ###########
try:
    val = (fname,mname,lname,gender,age)
    seq = "insert into student values(%s,%s,%s,%s,%s)"
    my_cursor.execute(seq,val)
    print("##########  ACCOUNT CREATED SUCCESFULLY  ############")
except Exception as e:
    print("Account not created succesful \n Maybe the name already exists")