package main

import (
        "context"
        "flag"
        "fmt"
        "os"
        "path/filepath"

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
                return nil, fmt.Errorf("error reading file %s: %v", path, err)
        }

        return &FileInfo{
                Path:    path,
                Name:    filepath.Base(path),
                Type:    filepath.Ext(path),
                Content: string(content),
        }, nil
}

func analyzeFiles(ctx context.Context, llm *anthropic.LLM, kg *KnowledgeGraph, systemPrompt string) error {
        for _, file := range kg.Files {
                summary, err := summarizeFile(ctx, llm, file, systemPrompt)
                if err != nil {
                        return fmt.Errorf("error summarizing file %s: %v", file.Path, err)
                }
                file.Summary = summary
        }

        for path1, file1 := range kg.Files {
                for path2, file2 := range kg.Files {
                        if path1 != path2 {
                                similarity, err := calculateSimilarity(ctx, llm, file1, file2, systemPrompt)
                                if err != nil {
                                        return fmt.Errorf("error calculating similarity between %s and %s: %v", path1, path2, err)
                                }
                                if similarity > 0.5 { // Arbitrary threshold, adjust as needed
                                        kg.AddEdge(path1, path2, similarity)
                                }
                        }
                }
        }

        return nil
}

func summarizeFile(ctx context.Context, llm *anthropic.LLM, file *FileInfo, systemPrompt string) (string, error) {
        prompt := fmt.Sprintf("%s\n\nSummarize the following file content in one concise sentence, capturing the main idea or purpose:\n\nFile name: %s\nContent:\n%s", 
                systemPrompt, file.Name, file.Content)
        
        completion, err := llm.Call(ctx, prompt)
        if err != nil {
                return "", err
        }
        
        return completion, nil
}

func calculateSimilarity(ctx context.Context, llm *anthropic.LLM, file1, file2 *FileInfo, systemPrompt string) (float64, error) {
        prompt := fmt.Sprintf("%s\n\nCalculate the similarity between these two file summaries, considering main themes, topics, and purposes. Provide a similarity score between 0 (completely different) and 1 (identical), followed by a brief explanation:\n\nFile 1 (%s): %s\nFile 2 (%s): %s", 
                systemPrompt, file1.Name, file1.Summary, file2.Name, file2.Summary)
        
        completion, err := llm.Call(ctx, prompt)
        if err != nil {
                return 0, err
        }
        
        // Extract the similarity score from the response
        var similarity float64
        _, err = fmt.Sscanf(completion, "%f", &similarity)
        if err != nil {
                return 0, fmt.Errorf("error parsing similarity score: %v", err)
        }
        
        return similarity, nil
}

func visualizeGraph(kg *KnowledgeGraph, outputPath string) error {
        data, err := dot.Marshal(kg.Graph, "knowledge_graph", "", "  ")
        if err != nil {
                return fmt.Errorf("error marshaling graph: %v", err)
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

        llm, err := anthropic.New(anthropic.WithToken(*anthropicKey))
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
