# Folder Knowledge Graph

A Go-based tool that analyzes files in a directory and creates a knowledge graph based on content similarities using the Anthropic API.

## Features

- Analyzes multiple file types
- Uses Anthropic API (Claude) for content summarization and similarity calculation
- Outputs graph in DOT format for visualization
- Customizable system prompt for AI analysis

## Prerequisites

- Go 1.16+
- Anthropic API key
- Graphviz (for visualization)

## Installation

```bash
git clone https://github.com/athreesh/folder-kg.git
cd folder-kg
go mod tidy
go build
```

## Usage

Set your Anthropic API key:
```bash
export ANTHROPIC_API_KEY="your-api-key-here"
```

Run the tool:
```bash
./folder-kg -dir /path/to/directory -anthropic-key $ANTHROPIC_API_KEY -output graph.dot -system-prompt system_prompt.txt
```

Options:
- `-dir`: Directory to process (required)
- `-anthropic-key`: Anthropic API key (required)
- `-output`: Output DOT file path (default: knowledge_graph.dot)
- `-system-prompt`: Path to system prompt file (default: system_prompt.txt)

Visualize the graph:
```bash
dot -Tpng graph.dot -o graph.png
```

## Customizing the System Prompt

Edit the `system_prompt.txt` file to customize the instructions given to the AI for file analysis and similarity calculation.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Link: [https://github.com/athreesh/folder-kg](https://github.com/athreesh/folder-kg)
