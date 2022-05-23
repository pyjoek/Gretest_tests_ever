use std::io;

fn get_input() -> i32{
    let mut x = String::new();
    io::stdin().read_line(&mut x).expect("failed to read line");
    let number : i32 = x.trim().parse().unwrap();
    return number;
}
fn add(){
    println!("The first number: ");
    let number1 = get_input();
    println!("The second number: ");
    let number2 = get_input();
    let sum = number1 + number2;
    println!("{}",sum);
}
fn sub() {
    println!("The first number: ");
    let number1 = get_input();
    println!("The second number: ");
    let number2 = get_input();
    let sub = number1 - number2;
    println!("{}",sub);
}
fn multy(){
    println!("The first number: ");
    let number1 = get_input();
    println!("The second number: ");
    let number2 = get_input();
    let multy = number1 * number2;
    println!("{}",multy);
}
fn divide() {
    println!("The first number: ");
    let number1 = get_input();
    println!("The second number: ");
    let number2 = get_input();
    let divide = number1 / number2;
    println!("{}",divide);
}
fn loops(){
    println!("Do you wanna do it again!! [yes | y]: ");
    let mut option = String::new();
    io::stdin().read_line(&mut option).expect("failed to read line");
    if option == "yes\n" || option == "y\n"{
        main();
    }else{
        println!("Bye !!!");
    }
}
fn main(){
    println!("Enter the operation you want to use [add | sub | multiply | divide]: ");
    let mut choice = String::new();
    io::stdin().read_line(&mut choice).expect("failed to read line");
    if choice == "add\n"{
        add();
        loops();
    }else if choice == "sub\n"{
        sub();
        loops();
    }else if choice == "multiply\n"{
        multy();
        loops();
    }else if choice == "divide\n"{
        divide();
        loops();
    }else{
        println!("Sorry i don't understand your choice!!!");
    }
}