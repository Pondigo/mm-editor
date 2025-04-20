# Mermaid Editor

A Mermaid diagram editor built with Go and WebAssembly.

## Features

- Create and edit Mermaid diagrams visually
- Export diagrams as Mermaid syntax
- Real-time preview of diagram rendering

## Prerequisites

- Go 1.18 or later
- A modern web browser (Chrome, Firefox, Safari, Edge)

## Building the Project

1. Clone the repository:
   ```
   git clone https://github.com/Pondigo/mermaid-editor.git
   cd mermaid-editor
   ```

2. Make sure the build script is executable:
   ```
   chmod +x scripts/build.sh
   ```

3. Build the WebAssembly binary:
   ```
   ./scripts/build.sh
   ```
   
   The script will automatically download `wasm_exec.js` if it can't find it in your Go installation.

## Running the Editor

1. The simplest way to run the editor is with our start script:
   ```
   ./scripts/start.sh
   ```
   
   This will:
   - Build the application if needed
   - Start an HTTP server on port 8080
   - Store the process ID for easy shutdown

2. To stop the server:
   ```
   ./scripts/stop.sh
   ```

3. Alternatively, you can manually run an HTTP server:
   ```
   # Python
   python3 -m http.server -d public 8080
   
   # Node.js (http-server)
   npx http-server public
   
   # Go
   go install golang.org/x/tools/cmd/goserve@latest
   goserve -dir public
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## Development Workflow

For development with hot reload:

1. Install the required dependencies:
   ```bash
   # macOS
   brew install entr
   npm install -g browser-sync
   
   # Linux
   apt-get install entr
   npm install -g browser-sync
   ```

2. Start the development server:
   ```bash
   ./scripts/dev.sh
   ```
   This will:
   - Build the WebAssembly application
   - Start a browser-sync server on port 8080
   - Watch for file changes in Go source files
   - Automatically rebuild when changes are detected
   - Refresh the browser when files are updated

3. Open your browser at http://localhost:8080

4. To stop the development server:
   ```bash
   ./scripts/stop-dev.sh
   ```

## How to Use

1. Once loaded, you'll see the diagram canvas and the Mermaid preview panel.
2. Click "Add Node" to create a new node.
3. Drag nodes to position them.
4. Use "Add Edge" to connect nodes.
5. The Mermaid syntax will update in real-time in the preview panel.
6. Click "Export" to download the Mermaid diagram as a .mmd file.

## Project Structure

- `/cmd/webapp`: WebAssembly entry point
- `/internal/core/diagram`: Core diagram models and logic
- `/internal/wasm`: JavaScript bindings and WebAssembly interfaces
- `/public`: Web assets and compiled WebAssembly
- `/scripts`: Build and utility scripts

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT 