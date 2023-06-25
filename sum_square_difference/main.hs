main = print $ let
    n = 100
    sqrOfSum = (n * (n + 1) `div` 2) ^ 2
    sumOfSqrs = (n * (n + 1) * (2 * n + 1)) `div` 6
    in sqrOfSum - sumOfSqrs