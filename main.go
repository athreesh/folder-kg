package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/encoding/dot"
)

type FileInfo struct {
	Path     string
	Name     string
	Type     string
	Content  string
	Summary  string
}

type KnowledgeGraph struct {
	Files  map[string]*FileInfo
	Graph  *simple.WeightedUndirectedGraph
	NodeID map[string]int64
}

func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		Files:  make(map[string]*FileInfo),
		Graph:  simple.NewWeightedUndirectedGraph(0, 0),
		NodeID: make(map[string]int64),
	}
}

func (kg *KnowledgeGraph) AddFile(file *FileInfo) {
	kg.Files[file.Path] = file
	id := kg.Graph.NewNode().ID()
	kg.NodeID[file.Path] = id
	kg.Graph.AddNode(simple.Node(id))
}

func (kg *KnowledgeGraph) AddEdge(file1, file2 string, weight float64) {
	from := simple.Node(kg.NodeID[file1])
	to := simple.Node(kg.NodeID[file2])
	kg.Graph.SetWeightedEdge(simple.WeightedEdge{F: from, T: to, W: weight})
}

func readSystemPrompt(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func processFile(path string) (*FileInfo, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Path:    path,
		Name:    filepath.Base(path),
		Type:    filepath.Ext(path),
		Content: string(content),
	}, nil
}

func analyzeFiles(ctx context.Context, llm llms.LanguageModel, kg *KnowledgeGraph, systemPrompt string) error {
	for _, file := range kg.Files {
		summary, err := summarizeFile(ctx, llm, file, systemPrompt)
		if err != nil {
			return err
		}
		file.Summary = summary
	}

	for path1, file1 := range kg.Files {
		for path2, file2 := range kg.Files {
			if path1 != path2 {
				similarity, err := calculateSimilarity(ctx, llm, file1, file2, systemPrompt)
				if err != nil {
					return err
				}
				if similarity > 0.5 {
					kg.AddEdge(path1, path2, similarity)
				}
			}
		}
	}

	return nil
}

func summarizeFile(ctx context.Context, llm llms.LanguageModel, file *FileInfo, systemPrompt string) (string, error) {
	prompt := fmt.Sprintf("Summarize the following file content in one sentence:\n\nFile name: %s\nContent:\n%s", file.Name, file.Content)
	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, systemPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	})
	if err != nil {
		return "", err
	}
	return completion.Choices[0].Content, nil
}

func calculateSimilarity(ctx context.Context, llm llms.LanguageModel, file1, file2 *FileInfo, systemPrompt string) (float64, error) {
	prompt := fmt.Sprintf("Calculate the similarity between these two file summaries:\n\nFile 1 (%s): %s\nFile 2 (%s): %s\n\nProvide only the numerical score between 0 and 1.", file1.Name, file1.Summary, file2.Name, file2.Summary)
	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, systemPrompt),
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	})
	if err != nil {
		return 0, err
	}
	var similarity float64
	_, err = fmt.Sscanf(strings.TrimSpace(completion.Choices[0].Content), "%f", &similarity)
	return similarity, err
}

func visualizeGraph(kg *KnowledgeGraph, outputPath string) error {
	dotGraph := dot.NewGraph(kg.Graph)
	data, err := dot.Marshal(dotGraph, "knowledge_graph", "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func main() {
	directory := flag.String("dir", "", "Directory to process")
	anthropicKey := flag.String("anthropic-key", "", "Anthropic API Key")
	outputFile := flag.String("output", "knowledge_graph.dot", "Output DOT file path")
	systemPromptFile := flag.String("system-prompt", "system_prompt.txt", "System prompt file")
	flag.Parse()

	if *directory == "" || *anthropicKey == "" {
		fmt.Println("Please specify a directory using -dir and provide an Anthropic API key using -anthropic-key")
		os.Exit(1)
	}

	systemPrompt, err := readSystemPrompt(*systemPromptFile)
	if err != nil {
		fmt.Printf("Error reading system prompt: %v\n", err)
		os.Exit(1)
	}

	kg := NewKnowledgeGraph()

	err = filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := processFile(path)
			if err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
				return nil
			}
			kg.AddFile(file)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		os.Exit(1)
	}

	llm, err := anthropic.New(anthropic.WithApiKey(*anthropicKey))
	if err != nil {
		fmt.Printf("Error initializing Anthropic: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	err = analyzeFiles(ctx, llm, kg, systemPrompt)
	if err != nil {
		fmt.Printf("Error analyzing files: %v\n", err)
		os.Exit(1)
	}

	err = visualizeGraph(kg, *outputFile)
	if err != nil {
		fmt.Printf("Error visualizing graph: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Knowledge graph generated and saved as '%s'\n", *outputFile)
}
