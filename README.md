## Go-Mazer

[Mazer](https://github.com/zuramai/mazer) is a Admin Dashboard Template that can help you develop faster. Made with Bootstrap 5. No jQuery dependency.

[Go-Mazer](https://github.com/ansufw/go-mazer) is another version of Mazer that is built using Go (Golang) [Templ](https://github.com/a-h/templ) and [Fiber](https://github.com/gofiber/fiber). 

## Requirements

- Go (Golang) 1.21 or higher
- Fiber
- HTML, CSS, and JavaScript
- Bootstrap v5.3.3
- JWT for authentication
- Esbuild for building assets - https://esbuild.github.io

## Features

- Frontend Dashboard


## Usage

### Development

1. Clone the repository:
   ```bash
   git clone https://github.com/ansufw/go-mazer.git
   cd go-mazer
   ```

2. Install dependencies:
   ```bash
   bun install
   go mod tidy
   ```

3. Run the development server:
   - Start the frontend build in watch mode:
     ```bash
     bun run dev
     ```
   - Start the backend server with Air (live reload):
     ```bash
     air
     ```

### Production Build

1. Build frontend assets:
   ```bash
   bun run build
   ```

2. Build the Go binary:
   ```bash
   go build -o bin/server cmd/main.go
   ```

## License

This project is licensed under the MIT License.

### Third-Party Licenses

This project uses the **Mazer Admin Template**, which is licensed under the MIT License.

- Project: Mazer
- Author: zuramai
- Source: https://github.com/zuramai/mazer
- License: MIT License

The original copyright notice and license text
from the Mazer project are included in this repository.