# Test Data Bus

Test Data Bus is an application that implements the In-Memory Message Bus.

## Installation and Running

1. Clone the repository:

    ```shell
    git clone https://github.com/MihasBel/test-data-bus.git
    cd test-data-bus
    ```

2. Build and run the application using Docker:

    ```shell
    docker build -t test-data-bus .
    docker run -p 9080:9080 test-data-bus
    ```

   The application will be available on port 9080.

## How It Works

Test Data Bus uses gRPC for efficient communication. The application exposes gRPC methods for other services 
to publish messages and for subscribers to receive messages.
gRPC streams are used to deliver the messages to the subscribers in real-time.


The application uses an in-memory message bus to manage message delivery. Messages of different types
are stored in a common pool. When a subscriber subscribes to a certain type of messages,
it starts receiving new messages of this type.

The message delivery mechanism ensures that messages are not lost and are delivered to subscribers
in the order they were received. Each subscriber has an offset that points to the next message to be read.
When a new message is read, the offset is incremented.

