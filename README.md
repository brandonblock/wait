# wait

Implementing this package in your microservice's main.go (via NewWaiter) with routes for StoreEvent and SetLimit will allow you to limit actions to once per specified duration.

Waiter creates a channel for requests and loads them when StoreEvent is called. The consumer function only handles the request body upon receiving a tick from the <-limiter channel.
