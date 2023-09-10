fn is_even(number: u8) -> bool {
    number % 2 == 0
}

fn custom_filter<F>(list: &Vec<u8>, func: F) -> Vec<u8>
  where F: Fn(u8) -> bool {
    let mut results: Vec<u8> = vec![];

    for value in list {
      if func(*value) {
        results.push(*value);
      }
    }

    results
}

fn main() {
  let numbers: Vec<u8> = vec![1, 2, 3, 5, 6, 4, 7, 8, 9, 22, 12, 10];

  let even_numbers_from_built_in = numbers.iter().copied().filter(|x| is_even(*x)).collect::<Vec<u8>>();
  let even_numbers_from_custom = custom_filter(&numbers, is_even);

  println!("Even numbers from built-in filter: {:?}", even_numbers_from_built_in);
  println!("Even numbers from custom filter: {:?}", even_numbers_from_custom);

  let answer: [u8; 7] = [2, 6, 4, 8, 22, 12, 10];

  assert_eq!(is_even(2), true);
  assert_eq!(is_even(3), false);
  assert_eq!(is_even(0), true);
  assert_eq!(even_numbers_from_custom, answer);
  assert_eq!(even_numbers_from_built_in, answer);

  println!("Tests passed!");
}