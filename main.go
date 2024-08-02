// main.go
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
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

func processFile(path string) (*FileInfo, error) {
	content, err := readFileContent(path)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Path:    path,
		Name:    filepath.Base(path),
		Type:    filepath.Ext(path),
		Content: content,
	}, nil
}

func readFileContent(path string) (string, error) {
	// Implement file reading logic based on file type
	// For now, we'll just read the first few lines of text files
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for i := 0; i < 10 && scanner.Scan(); i++ {
		content += scanner.Text() + "\n"
	}

	return content, nil
}

func analyzeFiles(llm llms.LanguageModel, kg *KnowledgeGraph) error {
	ctx := context.Background()

	for _, file := range kg.Files {
		summary, err := summarizeFile(ctx, llm, file)
		if err != nil {
			return err
		}
		file.Summary = summary
	}

	// Create edges based on similarity
	for path1, file1 := range kg.Files {
		for path2, file2 := range kg.Files {
			if path1 != path2 {
				similarity, err := calculateSimilarity(ctx, llm, file1, file2)
				if err != nil {
					return err
				}
				if similarity > 0.5 { // Arbitrary threshold
					kg.AddEdge(path1, path2, similarity)
				}
			}
		}
	}

	return nil
}

func summarizeFile(ctx context.Context, llm llms.LanguageModel, file *FileInfo) (string, error) {
	prompt := fmt.Sprintf("Summarize the following file content in one sentence:\n\n%s", file.Content)
	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts("system", "You are a helpful assistant that summarizes file contents."),
		llms.TextParts("human", prompt),
	})
	if err != nil {
		return "", err
	}
	return completion.Choices[0].Content, nil
}

func calculateSimilarity(ctx context.Context, llm llms.LanguageModel, file1, file2 *FileInfo) (float64, error) {
	prompt := fmt.Sprintf("On a scale of 0 to 1, how similar are these two file summaries?\n\nFile 1: %s\nFile 2: %s\n\nProvide only the numerical score.", file1.Summary, file2.Summary)
	completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts("system", "You are a helpful assistant that calculates similarity between file summaries."),
		llms.TextParts("human", prompt),
	})
	if err != nil {
		return 0, err
	}
	var similarity float64
	_, err = fmt.Sscanf(completion.Choices[0].Content, "%f", &similarity)
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

func interactiveFileSelection(directory string) ([]string, error) {
	var selectedFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Printf("Include %s? (y/n): ", path)
			var response string
			fmt.Scanln(&response)
			if response == "y" || response == "Y" {
				selectedFiles = append(selectedFiles, path)
			}
		}
		return nil
	})
	return selectedFiles, err
}

func main() {
	directory := flag.String("dir", "", "Directory to process")
	openaiKey := flag.String("openai-key", "", "OpenAI API Key")
	interactive := flag.Bool("interactive", false, "Enable interactive file selection")
	outputFile := flag.String("output", "knowledge_graph.dot", "Output DOT file path")
	flag.Parse()

	if *directory == "" || *openaiKey == "" {
		fmt.Println("Please specify a directory using -dir and provide an OpenAI API key using -openai-key")
		os.Exit(1)
	}

	kg := NewKnowledgeGraph()

	var filesToProcess []string
	var err error

	if *interactive {
		filesToProcess, err = interactiveFileSelection(*directory)
	} else {
		err = filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				filesToProcess = append(filesToProcess, path)
			}
			return nil
		})
	}

	if err != nil {
		fmt.Printf("Error processing directory: %v\n", err)
		os.Exit(1)
	}

	for _, path := range filesToProcess {
		file, err := processFile(path)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", path, err)
			continue
		}
		kg.AddFile(file)
	}

	llm, err := openai.New(openai.WithToken(*openaiKey))
	if err != nil {
		fmt.Printf("Error initializing OpenAI: %v\n", err)
		os.Exit(1)
	}

	err = analyzeFiles(llm, kg)
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
