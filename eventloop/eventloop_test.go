package eventloop

import (
    "fmt"
    "time"
    "os"
    "testing"
)

func TestEventLoop(t *testing.T) {
    // Create an event emitter
    emitter := NewEventEmitter()

    // Register event handlers
    emitter.On("start", func(e Event) {
        fmt.Println("Starting...")
    })

    emitter.On("tick", func(e Event) {
        fmt.Println("Tick:", e.Data)
    })

    emitter.On("stop", func(e Event) {
        fmt.Println("Stopping...")
        os.Exit(0)
    })

    // Start the event loop
    go func() {
        // Emit start event
        emitter.Emit("start", nil)

        // Simulate some ticks
        for i := 1; i <= 5; i++ {
            emitter.Emit("tick", i)
            time.Sleep(time.Second)
        }

        // Emit stop event
        emitter.Emit("stop", nil)
    }()

    // Keep the main goroutine running
    select {}
}
