# Concurrency and how to test it

## Task
Build a TCP server that shuts down gracefully.

We have:
- A TCP server that listens on a port and accepts connections.
- A handler function that processes incoming connections - reads lines from the connection and writes back the same strings to the client.

We want:
- To listen for a signal (SIGINT) and gracefully shut down the server.

How:
- We add contexts to the server and the handler functions.
- We listen for the sigterm and sigint (for testing purposes) and cancel the context.

Bonus:
In addition to writing the lines back to the client, we write them to a channel.
A separate goroutine reads from the channel and logs the lines.
The goroutine terminates when the server shuts down.