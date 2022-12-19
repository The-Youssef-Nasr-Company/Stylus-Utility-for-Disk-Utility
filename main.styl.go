import (
        "fmt"
        "time"
)

func loop() {
  fmt.Println("Drawing");
  time.Sleep(0.9);
  fmt.Println("Drawing .");
  time.Sleep(0.9);
  fmt.Println("Drawing . .");
  time.Sleep(0.9);
  fmt.Println("Drawing . . .");
  time.Sleep(0.9);
  loop();
}

loop();
