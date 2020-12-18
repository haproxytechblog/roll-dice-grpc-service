# RollDiceService gRPC Example

An example gRPC client and server.

1. Build the client and server applications:

    ```
    sudo docker build -t haproxy/rolldice-client -f Dockerfile.client .
    sudo docker build -t haproxy/rolldice-server -f Dockerfile.server .
    ```

2. Copy the **client** and **server** binaries to your host computer:

    ```
    sudo docker create -it --name client nramirez/rolldice-client bash
    sudo docker cp client:/app/client ./
    sudo docker rm -f client

    sudo docker create -it --name server nramirez/rolldice-server bash
    sudo docker cp server:/app/server ./
    sudo docker rm -f server
    ```

3. Start the server:

   ```
   SERVER_ADDRESS=:3000 ./server
   ```

4. Start the client. It sends a request to the server every other second:

   ```
   SERVER_ADDRESS=localhost:3000 NUMBER_OF_DICE=2 ./client
   ```