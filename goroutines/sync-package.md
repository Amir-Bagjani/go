## 6 Key Concepts You Should Know for Concurrency in the Sync Package
**Read the magic of mutex, sync map, sync once,… - vital tools for mastering concurrent programming in Go.**

![alt text](1035fb96-28e3-4d85-a131-54779db31e4b_1600x1067.webp)

As I delved into the sync package, I found its six main concepts for concurrent programming in Go. How about we delve deeper into these together? I'd truly value sharing them and reflecting on my experiences with you.

Let’s discuss them in order of how frequently they’re used, so you can get a feel for which concepts are most relevant in real-world scenarios.


1. ### sync.Mutex and sync.RWMutex
You know, mutex (mutual exclusion) is like an old buddy for us gophers. When dealing with goroutines, it’s super important to make sure they don’t access resources at the same time, and mutex helps us with that.

#### sync.Mutex
Check out this simple example where I didn’t use a mutex to safeguard our variable a:

```go
var a = 0

func Add() {
  a++
}

func main() {
  for i := 0; i < 500; i++ {
    go Add()
  }

  time.Sleep(5 * time.Second)
  fmt.Println(a)
}
```

The outcome of this code is unpredictable. You might get 500 if you’re lucky, but often, the result will be less than 500. Now, let’s enhance our Add function using a mutex:

```go
var mtx = sync.Mutex{}

func Add() {
  mtx.Lock()
  defer mtx.Unlock()
  a++
}
```

Now, the code delivers the expected outcome. But what about using sync.RWMutex?

**“Hold on, I noticed sync.Mutex has a method called TryLock. What’s its purpose?”**

So, the TryLock method is like trying to grab the lock without waiting in line. If it’s free, it takes the lock and gives you a thumbs up with true. But if another goroutine is hogging the lock, it just returns false right away, without hanging around.

#### Why sync.RWMutex?
Imagine you're checking out the a variable, but other goroutines are also tweaking it. You might end up with outdated info. So, what's the fix for this?

Let’s take a step back and use our old method, adding sync.Mutex to our Get() function:
```go
func Add() {
  mtx.Lock()
  defer mtx.Unlock()

  a++
}

func Get() int {
  mtx.Lock()
  defer mtx.Unlock()

  return a
}
```

But the issue here is that if your service or program calls Get() millions of times and only calls Add() a few times, we’re essentially wasting resources locking everything up when we’re not even modifying it most of the time.

That’s where sync.RWMutex swoops in to save our day, This clever little tool was designed to help us handle situations where we’re reading and writing simultaneously.

```go
var mtx = sync.RWMutex{}

func Add() {
  mtx.Lock()
  defer mtx.Unlock()
  
  a++
}

func Get() int {
  mtx.RLock()
  defer mtx.RUnlock()
  
  return a 
}
```

So, what’s so great about RWMutex? Well, it allows for millions of concurrent reads while making sure that only one write can happen at a time.

Let me clarify how it works:

1. When writing, reading and other writing operations are locked.

2. When reading, writing is locked.

3. Multiple reads don’t lock each other.

#### sync.Locker
Oh, by the way, both Mutex and RWMutex implement the sync.Locker interface{}, Here’s what the signature looks like:

```go
// A Locker represents an object that can be locked and unlocked.
type Locker interface {
  Lock()
  Unlock()
}
```

If you ever want to create a function that takes in a Locker, you can use that function with your custom locker or sync mutex:

```go
func Add(mtx sync.Locker) {
  mtx.Lock()
  defer mtx.Unlock()
  
  a++
}
```

2. ### sync.WaitGroup
You might have noticed that I used `time.Sleep(5 * time.Second)` to wait for all goroutines to finish, but honestly, that’s a pretty ugly solution.

That’s where sync.WaitGroup comes in:

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 500; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      Add()
    }()
  }
  
  wg.Wait()
  fmt.Println(a)
}
```

The sync.WaitGroup has 3 main methods: **Add**, **Done**, and **Wait**.

There’s Add(delta int): This method increases the WaitGroup counter by the value of delta. You’d usually call it **before spawning a goroutine**, indicating there’s an extra task that needs completion.

The other 2 methods are pretty straightforward:

Done is called when a goroutine wraps up its task.

Wait blocks the caller until the WaitGroup counter hits zero, meaning all spawned goroutines have finished their tasks.

What do you think would happen if we put the wait group inside the `go func() {}`?

```go
go func() {
  wg.Add(1)
  defer wg.Done()
  Add()
}()
```

My compiler shouts, “should call wg.Add(1) before starting the goroutine to avoid a race” and my runtime panics, “panic: sync: WaitGroup is reused before previous Wait has returned”.

3. ### sync.Once
Imagine you have a CreateInstance() function in a package, but you need to ensure it’s initialized before using it. So you call it multiple times in different places, and your implementation looks like this:

```go
var i = 0
var _isInitialized = false

func CreateInstance() {
  if _isInitialized {
    return
  }
  
  i = GetISomewhere()
  _isInitialized = true
}
```

But what if multiple goroutines call this method? The `i = GetISomeWhere` line would run multiple times, even though you only want it to execute once for stability.

You could use a mutex lock, which we’ve discussed earlier, but the sync package offers a more convenient method: sync.Once

```go
var i = 0
var once = &sync.Once{}

func CreateInstance() {
  once.Do(func() {
    i = GetISomewhere()
  })
}
```
With sync.Once, you can make sure a function is executed just a single time, no matter how many times it’s called or how many goroutines call it at the same time.

4. ### sync.Pool
Imagine you’ve got a pool that holds a bunch of objects you’d like to reuse over and over. This can take some pressure off the garbage collector, especially if creating and destroying these resources is expensive.

So, whenever you need an object, you can just take it from the pool. And when you’re finished using it, you can put it back into the pool for reuse later on.

```go
var pool = sync.Pool{
  New: func() interface{} {
    return 0
  },
}

func main() {
  pool.Put(1)
  pool.Put(2)
  pool.Put(3)
  
  a := pool.Get().(int)
  b := pool.Get().(int)
  c := pool.Get().(int)
  
  fmt.Println(a, b, c) // Output: 1, 3, 2 (order may vary)
}
```

Keep in mind that the order in which you put objects into the pool isn’t necessarily the order they’ll come out, even if the sequence doesn’t appear random when you run the above code multiple times.

Let me share some tips for using sync.Pool:

It’s great for objects that live a long time and have multiple instances you need to manage, like database connections (1000 connections?), worker goroutines, or even buffers.

Always reset the state of objects before returning them to the pool. This way, you can avoid any unintentional data leaks or strange behaviors.

Don’t count on the objects that are already in the pool, because they could be deallocated unexpectedly.

5. ### sync.Map
6. When you’re working with maps concurrently, it’s a bit like using RWMutex. You can have multiple reads happening at once, but you can’t have multiple read-writes or write-writes. If there’s a conflict, your service will crash instead of overwriting data or causing unexpected behaviors.

That’s where sync.Map comes in handy, as it helps us avoid this problem. Let’s take a closer look at what sync.Map has to offer:

1. CompareAndDelete(key, old any) - go 1.20: deletes a key’s entry if values match; returns false if no value exists or old value is nil.

2. CompareAndSwap(key, old, new any) - go 1.20: swaps old and new values for a key if they match, just make sure old value is comparable.

3. Swap(key, value any) (previous any, loaded bool) (go 1.20): exchanges the value for a key and returns the old value, if it exists.

4. LoadOrStore(key, value any) (actual any, loaded bool): gets the current key value or saves and returns the provided value if it’s not there

5. Range (f func(key, value any) bool): loops through the map, applying function f to each key-value pair. If f says returns false, it stops.

Store, Delete, Load, LoadAndDelete

**“Why we don’t just use a regular map with a Mutex?”**

I usually go for a map with an RWMutex, but it’s important to recognize the power of sync.Map in certain situations. So, where does it really shine?

If you’ve got many goroutines accessing separate key sets in a map, a regular map with a single mutex can cause contention since it locks the entire map for just a single write operation.

On the other hand, sync.Map uses a more refined locking mechanism, helping to minimize contention in such scenarios.

6. ### sync.Cond

Think of sync.Cond as a condition variable that supports multiple goroutines waiting and interacting with each other. To get a better understanding, let’s see how to use it.

First off, we need to create the `sync.Cond` with a locker (I’ve explained what this interface is earlier):

```go
var mtx sync.Mutex
var cond = sync.NewCond(&mtx)
```

A goroutine calls cond.Wait and waits for a signal from somewhere else to continue its execution:

```go
func dummyGoroutine(id int) {
  cond.L.Lock()
  defer cond.L.Unlock()
  fmt.Printf("Goroutine %d is waiting...\n", id)
  cond.Wait()
  fmt.Printf("Goroutine %d received the signal.\n", id)
}
```

Then, another goroutine (like the main goroutine) calls cond.Signal(), allowing our waiting goroutine to carry on:

```go
func main() {
  go dummyGoroutine(1)
  
  time.Sleep(1 * time.Second)

  fmt.Println("Sending signal...")
  cond.Signal()

  time.Sleep(1 * time.Second)
}
```
Here’s what the result looks like:

```bash
Goroutine 1 is waiting...
Sending signal...
Goroutine 1 received the signal.
```

What if there are multiple goroutines waiting for our signal? That’s when we can use Broadcast:

```go
func main() {
  go dummyGoroutine(1)
  go dummyGoroutine(2)
  
  time.Sleep(1 * time.Second)
  cond.Broadcast() // broadcast to all goroutines
  time.Sleep(1 * time.Second)
}
```

And guess what? Here are the logs:

```bash
Goroutine 1 is waiting...
Goroutine 2 is waiting...
Goroutine 2 received the signal.
Goroutine 1 received the signal.
```

**“Why can goroutine 2 join the waiting state when we locked the mutex at the beginning of dummyGoroutine?”**
Well, both goroutines can enter the waiting state because cond.Wait() actually unlocks the mutex (cond.L) inside it. This allows other goroutines to acquire the lock and move forward.

**"Why do I need to use sync.Cond instead of channels?"**
1. **Channels**: In Go, channels are the primary means to pass data between goroutines, One goroutine sends data, and another receives it.

2. **sync.Cond**: It's about waiting for certain conditions, a goroutine can wait, and when the condition is met, it gets notified. And the standout feature is… If you have multiple goroutines waiting, sync.Cond can alert them all at once.

**“What if I choose to use Signal() instead of Broadcast()?”**

In that case, only one goroutine will receive the signal, while the others will remain blocked until they get a new signal.

To sum up:
Use channels in Go for direct data transfer and synchronization.

But when you're dealing with a more complex condition, or scenarios where multiple goroutines wait and need to act together, turn to sync.Cond.