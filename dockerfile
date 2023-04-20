# Use an official Go runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/bernie.top/dymest/loan-backend/

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the command to run the app when the container starts
CMD ["./app"]
