main = let 
  result = [a * b * c | 
            a <- [1..1000], 
            b <- [a..1000],
            let cSq = a^2 + b^2
              c = intSqrt cSq,
            c^2 == cSq && a + b + c == 1000]
  in print . head $ result

intSqrt :: (Integral a) => a -> a
intSqrt = floor . sqrt . fromIntegral
