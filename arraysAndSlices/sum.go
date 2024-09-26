package arrays

func Sum(numbers []int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(numbersToSum ...[]int) []int {
  var sumArray []int
  for _, array := range numbersToSum {
    sumArray = append(sumArray, Sum(array))
  }
  return sumArray
}

func SumAllTails(numbersToSum ...[]int) []int {
  var sumArray []int
  for _, array := range numbersToSum {
    tail := array[1:]
    sumArray = append(sumArray, Sum(tail))
  }
  return sumArray
}

