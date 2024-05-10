# Singleton Pattern

The Singleton Pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. This pattern is particularly useful when there's a need to coordinate actions that can't be shared among multiple instances.

The Singleton Pattern with sync.Once is a robust and elegant solution for managing single instances across multiple goroutines. By combining the benefits of the traditional Singleton Pattern with thread-safety, this approach enables secure and efficient access to shared resources and configurations.

## References

[How singleton pattern works with Golang](https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f)
