### 1. Goroutines — what they are, how to launch them

* Goroutines are lightweight threads managed by the Go runtime.
* They are cheaper than OS threads (low memory, fast startup, scalable).
* You launch a goroutine using the `go` keyword before a function call.
* Example:

  ```go
  go myFunction()
  ```
* The function runs concurrently with the rest of the program.
* The `main` function does not wait for goroutines automatically—use synchronization (like channels or `sync.WaitGroup`).

### 2. Channels — typed conduits for communication between goroutines

* Channels allow safe communication between goroutines.
* They are **typed**, meaning they carry values of a specific type.
* Create a channel using:

  ```go
  ch := make(chan int)
  ```
* Send data:

  ```go
  ch <- 10
  ```
* Receive data:

  ```go
  x := <-ch
  ```
* Channels block by default (send waits for receiver, receive waits for sender).

### 3. Buffered channels — channels with capacity

* Buffered channels can store multiple values without immediate receiver.
* Created with capacity:

  ```go
  ch := make(chan int, 3)
  ```
* Sending only blocks when buffer is full.
* Receiving only blocks when buffer is empty.
* Useful for reducing synchronization overhead.

### 4. Range and close — iterating over a channel, knowing when it is done

* A sender can close a channel using:

  ```go
  close(ch)
  ```
* Receivers can detect closure using `range`:

  ```go
  for v := range ch {
      fmt.Println(v)
  }
  ```
* Loop ends automatically when channel is closed and empty.
* Only the sender should close the channel.

### 5. Select — waiting on multiple channel operations

* `select` lets a goroutine wait on multiple channel operations.
* It blocks until one case is ready.
* Example:

  ```go
  select {
  case msg1 := <-ch1:
      fmt.Println(msg1)
  case msg2 := <-ch2:
      fmt.Println(msg2)
  }
  ```
* If multiple cases are ready, one is chosen randomly.

### 6. Default select — non-blocking channel operations

* Adding a `default` case makes `select` non-blocking.
* If no channel is ready, `default` executes immediately.
* Example:

  ```go
  select {
  case msg := <-ch:
      fmt.Println(msg)
  default:
      fmt.Println("no message available")
  }
  ```
* Useful for polling or avoiding blocking behavior.

### 7. sync.Mutex — protecting shared state when channels are not the right tool

* `sync.Mutex` is used to protect shared data from race conditions.
* Lock before accessing shared resource, unlock after:

  ```go
  var mu sync.Mutex

  mu.Lock()
  // critical section
  mu.Unlock()
  ```
* Use when multiple goroutines access the same variable.
* Prefer channels for communication, but mutex for shared memory control.

### 8. Exercise: Equivalent Binary Trees (goroutines + channels + recursion)

* Goal: Check if two binary trees store the same values.
* Approach:

  * Traverse each tree using recursion (in-order traversal).
  * Send values into channels using goroutines.
  * Compare values received from both channels.
* Key concepts used:

  * Goroutines (parallel traversal)
  * Channels (stream values)
  * Synchronization (ensuring both traversals finish)
* Insight: Channels can turn recursive structures into streams for comparison.

### 9. Exercise: Web Crawler (concurrent crawling with goroutines)

* Goal: Crawl web pages concurrently without visiting the same URL twice.
* Approach:

  * Use goroutines to fetch multiple URLs in parallel.
  * Use a shared data structure (map) to track visited URLs.
  * Protect shared state using `sync.Mutex`.
* Key concepts used:

  * Goroutines (parallel fetching)
  * Mutex (avoid duplicate visits)
  * Recursion (crawl links)
* Insight: Combining concurrency + synchronization enables efficient real-world tasks like crawling.
