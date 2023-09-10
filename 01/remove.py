numbers = [1, 2, 3, 5, 6, 4, 7, 8, 9, 22, 12, 10]

def is_even(num: int):
  return num % 2 == 0

def customFilter(func, list: list[int]):
  result = []

  for value in list:
    if func(value):
      result.append(value)

  return result


evenNumbersFromBuiltIn = list(filter(is_even, numbers))
evenNumbersFromCustom = customFilter(is_even, numbers)

print("From built-in: ", evenNumbersFromBuiltIn)
print("From custom: ", evenNumbersFromCustom)

def tests():
  assert is_even(2) == True
  assert is_even(3) == False
  assert is_even(0) == True
  assert list(filter(is_even, numbers)) == [2, 6, 4, 8, 22, 12, 10]
  assert customFilter(is_even, numbers) == [2, 6, 4, 8, 22, 12, 10]

if __name__ == "__main__":
  tests()
  print("Tests passed")