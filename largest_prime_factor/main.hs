num :: Integer
num = 600_851_475_143

main :: IO ()
main = print $ findLargestPrimeFactor num 2

findLargestPrimeFactor :: Integer -> Integer -> Integer
findLargestPrimeFactor 1 prime = prime
findLargestPrimeFactor x prime
    | x `isDivisible` prime = findLargestPrimeFactor (x `div` prime) prime
    | otherwise = findLargestPrimeFactor x $ nextPrime prime

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
