package golangpractice
package main


func CalculateProgress(done int, total int) float64 {
  return float64(done) / float64(total)
}

func main() {
	CalculateProgress()
}