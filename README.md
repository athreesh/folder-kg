# Desktop Knowledge Graph Tool

A Go-based tool that analyzes files in a directory and creates a knowledge graph based on content similarities.

## Features

- Analyzes multiple file types (.ipynb, .pdf, etc.)
- Uses Anthropic API for content summarization and similarity calculation
- Interactive file selection mode
- Outputs graph in DOT format for visualization

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
./folder-kg -dir /path/to/directory -anthropic-key $ANTHROPIC_API_KEY -interactive -output graph.dot
```

Options:
- `-dir`: Directory to process (required)
- `-openai-key`: Anthropic API key (required)
- `-interactive`: Enable interactive file selection
- `-output`: Output DOT file path (default: knowledge_graph.dot)

Visualize the graph:
```bash
dot -Tpng graph.dot -o graph.png
```

## Quick Run

```bash
chmod +x run.sh
./run.sh
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Link: [https://github.com/athreesh/folder-kg](# Desktop Knowledge Graph Tool

A Go-based tool that analyzes files in a directory and creates a knowledge graph based on content similarities.

## Features

- Analyzes multiple file types (.ipynb, .pdf, etc.)
- Uses OpenAI API for content summarization and similarity calculation
- Interactive file selection mode
- Outputs graph in DOT format for visualization

## Prerequisites

- Go 1.16+
- OpenAI API key
- Graphviz (for visualization)

## Installation

```bash
git clone https://github.com/yourusername/desktop-knowledge-graph.git
cd desktop-knowledge-graph
go mod tidy
go build
```

## Usage

Set your OpenAI API key:
```bash
export OPENAI_API_KEY="your-api-key-here"
```

Run the tool:
```bash
./desktop-knowledge-graph -dir /path/to/directory -openai-key $OPENAI_API_KEY -interactive -output graph.dot
```

Options:
- `-dir`: Directory to process (required)
- `-openai-key`: OpenAI API key (required)
- `-interactive`: Enable interactive file selection
- `-output`: Output DOT file path (default: knowledge_graph.dot)

Visualize the graph:
```bash
dot -Tpng graph.dot -o graph.png
```

## Quick Run

```bash
chmod +x run.sh
./run.sh
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Link: [https://github.com/athreesh/folder-kg](https://github.com/athreesh/folder-kg)
