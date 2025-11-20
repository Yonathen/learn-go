## Installing and Configuring Go

Before you start writing Go code, you need to install and configure Go on your system.

### 1. Downloading Go

- Visit the official Go downloads page: [https://go.dev/dl/](https://go.dev/dl/)
- Choose your operating system (Windows, macOS, Linux) and download the installer.

### 2. Installing Go

**Windows**
- Run the installer you downloaded and follow the setup instructions.

**macOS**
- Open the PKG installer and follow the installation instructions.
- Alternatively, you can use Homebrew:
  ```
  brew install go
  ```

**Linux**
- Extract the tarball to `/usr/local`:
  ```
  tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
  ```
- Add Go to your `PATH`:
  ```
  export PATH=$PATH:/usr/local/go/bin
  ```

### 3. Verifying Installation

Open your terminal and run:
```
go version
```
You should see the installed Go version.

### 4. Setting Up Your Go Workspace

- Go recommends using a workspace directory for your code. By default, you can place your code anywhere, but commonly under your home directory:
  ```
  mkdir ~/go
  ```
- Set your workspace and bin paths (optional):
  ```
  export GOPATH=~/go
  export PATH=$PATH:$GOPATH/bin
  ```

### 5. Creating Your First Go File

- Create a directory for your project:
  ```
  mkdir ~/go/hello
  cd ~/go/hello
  ```
- Create a file named `hello.go` and add your code.

### 6. Running Your Code

- Use `go run` to execute:
  ```
  go run hello.go
  ```
- Use `go build` to create an executable:
  ```
  go build hello.go
  ./hello
  ```

---

Now you're ready to start writing and running Go programs!
