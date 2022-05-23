use std::io;
use rand::Rng;

fn int_input() -> i32 {
    let mut user = String::new();
    io::stdin().read_line(&mut user).expect("failed to read line");
    let number : i32 = user.trim().parse().unwrap();
    return number;
}

fn main(){
    let user_guess = int_input();
    let comp_guess = rand::thread_rng().gen_range(1,5);
    if user_guess == comp_guess{
        println!("you won the guess was {}",comp_guess);
    }else{
        println!("you loose the guess was {}",comp_guess);
    }
}