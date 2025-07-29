# Mengcp [re: me-nge-cap]
Random tools coded while learning and exploring MCP (Model Context Protocol) for building AI agents.

## Tools

> [!NOTE]
> All tools are designed for my personal use, so they may not work as expected for you. Feel free to modify the code to suit your needs.

### Code Editing Agent
Contains simple tools for code editing by AI agent using Claude (currently). It's based on the tutorial from [AmpCode](https://ampcode.com/how-to-build-an-agent) with slight modifications to the code structure and file organization for improved tidiness and performance.

- More modular code with organized file structure
- Performance optimizations in the code with map and in-memory cache for the definition setting

Feel free to explore and modify the code for your own purposes!

### Poke Agent
Inspired by [allenthomas](https://allenthomas.vercel.app/posts/mcp) technical walkthrough to explain what and how MCP works in simple terms.

## Development
### Prerequisites
- Go 1.23 or later
- Anthropic API key

### Local Setup

> [!CAUTION]
> Beware that this project is still in its early stages and may not be fully functional or stable yet. Use at your own risk. Especially the function where the agent needs to read and write files in your local machine.


1. Clone the repository
2. Install dependencies using `go mod tidy`
3. Set your Anthropic API key in the environment variable `ANTHROPIC_API_KEY`
   ```bash
   export ANTHROPIC_API_KEY=your_api_key_here
   ```
4. Run the application using `go run main.go`
   > Ensure you have the necessary permissions to run the application, especially if you're using a Mac

## Diagram

Check out the [diagram](/docs/diagram.md) for the visual explanation.

## Demo
See the [demo](/docs/demo.md) for a quick overview of the tools in action.

## Todo
- [ ] add tools for interacting with hosted actual budget
- [ ] add tools for interacting with local logseq graph