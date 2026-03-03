package main
import ("fmt")

func main() {
  nilai := 75
  if nilai >= 70 {
    fmt.Println("nilai kamu mencukupi syarat.")
    if nilai > 80 {
      fmt.Println("kerja bagus, kamu mendapat nilai A.")
    } else {
      fmt.Println("kerja bagus, kamu mendapat nilai B.")
    }
  } else {
    fmt.Println("nilai kamu kurang dari 70.")
  }
}