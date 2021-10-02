import (
        "fmt"
        "time"
)

func draw_looped_load() {
  fmt.Println("Drawing.");
  time.Sleep(0.5);
  fmt.Println("Drawing..");
  time.Sleep(0.5);
  fmt.Println("Drawing...");
  time.Sleep(0.5);
  draw_looped_load();
}

draw_looped_load();
