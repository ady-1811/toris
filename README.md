# TORIS - Terminal Organized and Rational IntelliSense

TORIS is an AI-powered terminal assistant designed to enhance your command-line experience. It leverages the power of Google's Gemini models to translate natural language into executable commands, automatically scan your terminal for errors, and provide actionable fixes directly in your workflow.

## ✨ Features

- **Natural Language to CLI:** Describe what you want to do in plain English, and TORIS will generate and execute the appropriate terminal command.
- **Smart Error Resolution:** Run into an error? TORIS can read your terminal's output history and instantly generate instructions and commands to fix it.
- **Safety First:** Built-in risk scoring system. Critical commands that might affect system stability require user confirmation before execution.
- **Cross-Platform:** Supports Linux, macOS, and Windows environments.

## 🚀 Prerequisites

- [Go](https://go.dev/doc/install) 1.21 or higher
- A [Google Gemini API Key](https://aistudio.google.com/app/apikey)

## 🛠️ Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/toris.git
   cd toris
   ```

2. **Configure Environment Variables:**
   TORIS requires a Gemini API key to communicate with the AI model. 
   Create a `.env` file in the root directory (or one level above your working directory, depending on your build setup) and add your API key:
   ```env
   GEMINI_API_KEY=your_gemini_api_key_here
   ```
   Alternatively, you can export it directly in your shell:
   ```bash
   export GEMINI_API_KEY="your_gemini_api_key_here"
   ```

3. **Install Dependencies & Build:**
   ```bash
   go mod download
   go build -o toris .
   ```
   *(Optional)* Move the binary to a directory in your system's `PATH` (e.g., `/usr/local/bin`):
   ```bash
   sudo mv toris /usr/local/bin/
   ```

## 📖 Usage

### 1. Initialization (Highly Recommended)
To enable the automatic error-scanning feature, you must initialize TORIS. This injects a small, safe recording script snippet into your `~/.bashrc` or `~/.zshrc` file.

```bash
toris init
```
*Note: Restart your terminal or run `source ~/.bashrc` (or `~/.zshrc`) after initialization for the changes to take effect.*

### 2. Running Commands via Natural Language
Use the `run` command to translate natural language into a terminal command. TORIS will print the generated command, its confidence score, risk level, and execute it.

```bash
toris run "List all hidden files in the current directory by size"
toris run "Find and kill the process running on port 8080"
```
*If a command is flagged as high-risk, TORIS will prompt you for confirmation `[y/n]` before executing.*

### 3. Scanning and Fixing Errors
If a command fails and spits out an error, simply type:

```bash
toris scan
```
TORIS will capture the recent terminal output, analyze the error using Gemini, and automatically suggest (and run) the fix.

## 🤝 Contributing Guidelines

Contributions are welcome and greatly appreciated! To contribute to TORIS:

1. **Fork the Project**
2. **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3. **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4. **Push to the Branch** (`git push origin feature/AmazingFeature`)
5. **Open a Pull Request**

### Development Notes
- **Adding new commands:** TORIS uses the Cobra CLI framework. New commands should be added inside the `app/cmd/` directory.
- **Code Quality:** Ensure your code is formatted with `go fmt` and passes standard `go vet` checks before submitting a pull request.

## 📄 License

Distributed under the MIT License. See the copyright headers in the source code files for more information.