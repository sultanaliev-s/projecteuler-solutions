main = print $ findNthPrime 10001

findNthPrime :: Integer -> Integer
findNthPrime n = findNthPrime' 2 (n - 1)

findNthPrime' :: Integer -> Integer -> Integer
findNthPrime' prime 0 = prime
findNthPrime' prime n = findNthPrime' (nextPrime prime) (n - 1)

nextPrime :: Integer -> Integer
nextPrime x
    | isPrime incr = incr
    | otherwise = nextPrime incr
    where incr = succ x

isPrime :: Integer -> Bool
isPrime x = not $ True `elem` results
    where results = map (isDivisible x) [2..isqrt x]

isDivisible :: Integer -> Integer -> Bool
isDivisible x divisor = x `mod` divisor == 0

isqrt :: Integer -> Integer
isqrt = floor . sqrt . fromIntegral