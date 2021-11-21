import (
        "fmt"
        "time"
)

func loop() {
  fmt.Println("Drawing.");
  time.Sleep(0.5);
  fmt.Println("Drawing..");
  time.Sleep(0.5);
  fmt.Println("Drawing...");
  time.Sleep(0.5);
  loop();
}

loop();
