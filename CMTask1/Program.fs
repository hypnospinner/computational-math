open System

// additional functions

let error = 9.6M / 30000000M

let sigma k = if k % 2 = 1 then ((-) 0M) else ((+) 0M)

let abs a = if a > 0M then a else -a

// sin function
// based on MacLaurin Series

let rec sinStep (x:decimal) (k:int) (prev:decimal) = 
    let this = prev * x * x / (decimal ((2 * k + 1) * 2 * k))

    if this < error 
        then 0M 
        else (sigma k this) + sinStep x (k+1) this 

let sin x = x + sinStep x 1 x

// cos function
// based on MacLaurin Series

let rec cosStep (x:decimal) (k:int) (prev:decimal) = 
    let this = prev * x * x / (decimal (2 * k * (2 * k - 1)))

    if this < error
        then 0M
        else (sigma k this) + cosStep x (k + 1) this

let cos x = 1M + cosStep x 1 1M

// Square Root function
// Based on Babylonian or Heron's Method

let rec sqrtStep (a:decimal) (prev:decimal) = 

    let this = (prev + a / prev) / 2M

    if abs (this - prev) < error 
        then this
        else sqrtStep a this 

let sqrt x = sqrtStep x 1M

let z1 x = (sqrt (1M + x * x)) * (sin (3M * x + 0.1M) + cos (2M * x + 0.2M))

let z2 x = (MathF.Sqrt (1.0f + x * x)) * (MathF.Sin (3.0f * x + 0.1f) + MathF.Cos (2.0f * x + 0.2f))

[<EntryPoint>]
let main _ =

    [0.2M..0.01M..0.3M] |> List.map (fun x -> 
        printf "z_1(%M) - z_2(%f) = %M - %f = %M\n" 
            x 
            (float x) 
            (z1 x) 
            (z2 (float32 x)) 
            (abs (z1 x - decimal (z2 (float32 x))))) |> ignore
    0 // return an integer exit code
