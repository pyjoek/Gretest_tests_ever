use std::io;

fn main(){
    let mut name = String::new();
    let mut age = String::new();
    let mut gender = String::new();
    let mut password = String::new();
    let mut confirm = String::new();
    println!("Enter your name: ");
    io::stdin().read_line(&mut name).expect("failed to read line");
    println!("How old are you: ");
    io::stdin().read_line(&mut age).expect("failed to read line");
    // loop{
    //     if age == "18"{
    //         println!("Too young to sign in right now");
    //     }else{
    //         break;
    //     }
    // }
    println!("What's your gender: ");
    io::stdin().read_line(&mut gender).expect("failed to read line");
    loop{
        println!("Enter your password: ");
        io::stdin().read_line(&mut password).expect("failed to read line");
        println!("confirm your password: ");
        io::stdin().read_line(&mut confirm).expect("failed to read line");
        if password == confirm{
            println!("Account ceated succesfully");
            println!("\n\n");
            println!("#####   ACCOUNT DETAILS ########NAME : {}AGE : {}GENDER : {}",name,age,gender);
            break;
        }else{
            println!("Password mismatch");
        }
    }
}