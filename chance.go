package chance

import (
  "log"
  "time"
  "math"
)

func epoch_seconds(date time.Time) {
  
}

func Score(ups, downs int) int {
  return ups - downs
}

func Hot(ups, downs int, sometime time.Time) float64 {
    somescore := Score(ups, downs)

    s := float64(somescore)

    order := math.Log10(math.Max(math.Abs(s), 1))

    var sign float64

    sign = 0.0

    if s > 0 {
      sign = 1
    } else if s < 0 {
      sign = -1
    }

    magictime, _ := time.Parse(time.RFC1123, "Thu, 08 Dec 2005 07:46:43 GMT")

    seconds := sometime.Sub(magictime)

    return order + sign * float64((seconds.Nanoseconds() / 1000000000)) / 45000.0
}

func calculate(ups, down int) float64 {
  n := float64(ups + down)

  if n == 0.0 {
    return 0.0
  }

  z := 1.0

  phat := float64(ups) / n

  log.Println( math.Sqrt(phat+z*z/(2*n)-z*((phat*(1-phat)+z*z/(4*n))/n))/(1+z*z/n) )
  return math.Sqrt(phat+z*z/(2*n)-z*((phat*(1-phat)+z*z/(4*n))/n))/(1+z*z/n)
}

func Confidence(ups, downs int) float64 {
  if ups + downs == 0 {
    return 0.0
  }

  return calculate(ups, downs)
}
