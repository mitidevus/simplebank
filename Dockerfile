# Build stage
FROM golang:1.21.3-alpine3.18 AS builder
# Declare current working directory inside the Image
WORKDIR /app
# Copy all necessary file to /app folder
# First "." is copy everything from current folder (of our application)
# Second "." is the current working directory inside the Image, where the files and folders are being copied to
COPY . .
# Build our app to a single binary executable file
# -o stands for output, main is the name of output binary file
# Pass in the main entrypoint file of our application, which is main.go (at the root of project)
RUN go build -o main main.go

# Run stage (use alpine linux as the base image)
FROM alpine:3.18
WORKDIR /app     
# Executable binary file from Build stage to run this stage image
# Copy file from Builder stage
# /app/main is the binary file we want to copy
# "." means copy to current working folder, represent the WORKDIR that we set above (/app)
# NOTE: WORKDIR "/app" of Build State and "/app" of Run stage is different
COPY --from=builder /app/main .
# Copy confile file (.env)
COPY app.env .


# Inform Docker that the container listens on the specified network port at run time
EXPOSE 8080

# Define the default command to run when the container starts
# Execute binary file "/app/main" of the final stage
CMD ["/app/main"]