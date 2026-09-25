fn is_palindrome(s: String) -> bool {
    let mut chars = s
        .bytes()
        .filter(u8::is_ascii_alphanumeric)
        .map(|c| c.to_ascii_lowercase());

    while let (Some(a), Some(b)) = (chars.next(), chars.next_back()) {
        if a != b {
            return false;
        }
    }

    true
}

fn main() {
    let s = "A man, a plan, a canal: Panama".to_string();

    println!("is palindrome: {}", is_palindrome(s))
}
