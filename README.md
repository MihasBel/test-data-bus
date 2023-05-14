# Test Data Bus

Test Data Bus is an application that implements the In-Memory Message Bus.

application implements a message bus using gRPC, with separate services for publishing and subscribing to messages. 
The protobuf definitions for these services can be found at the following locations in the repository:

Bus Service (Receiver): delivery/grpc/gen/v1/bus
Publisher Service (Sender): delivery/grpc/gen/v1/publisher
Bus Service: The Bus Service is responsible for sending messages into the bus. The protobuf definition 
([Bus proto](https://github.com/MihasBel/test-data-bus/tree/main/delivery/grpc/proto/v1/bus)) outlines the gRPC call
for this functionality.

Publisher Service: The Publisher Service handles subscriptions to message types.
Clients can open a gRPC stream to receive messages of a specified type, and can also manage their subscriptions using 
the gRPC calls defined in the protobuf file
([Publisher proto](https://github.com/MihasBel/test-data-bus/tree/main/delivery/grpc/proto/v1/publisher)).

Messages contain data encoded in base64 format and also include a message type.
The types of messages that can be processed by the application are specified in the application configuration.

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

