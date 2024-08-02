~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
".gitignore" 0L, 0B                                                                                     0,0-1         All
# Binaries
        "github.com/tmc/langchaingo/llms"
        "github.com/tmc/langchaingo/llms/openai"
        return &KnowledgeGraph{
        id := kg.Graph.NewNode().ID()
        return &FileInfo{
                Path:    path,
                Name:    filepath.Base(path),
                Type:    filepath.Ext(path),
        if err != nil {
                return "", err
        }
        defer file.Close()

        for _, file := range kg.Files {
                                }
        prompt := fmt.Sprintf("Summarize the following file content in one sentence:\n\n%s", file.Content)
        completion, err := llm.GenerateContent(ctx, []llms.MessageContent{
                llms.TextParts("system", "You are a helpful assistant that summarizes file contents."),
                llms.TextParts("human", prompt),
        })
                llms.TextParts("human", prompt),
                return err
        return selectedFiles, err
}
anishmaddipoti@anishmaddipoti-mlt folder-kg % clear
- Analyzes multiple file types (.ipynb, .pdf, etc.)
- Uses OpenAI API for content summarization and similarity calculation
## Installation
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

go build
```
## Contributing
- Analyzes multiple file types (.ipynb, .pdf, etc.)
- Uses OpenAI API for content summarization and similarity calculation
- Interactive file selection mode
## Installation

```bash
        data, err := dot.Marshal(dotGraph, "knowledge_graph", "", "  ")
anishmaddipoti@anishmaddipoti-mlt folder-kg % git add .
anishmaddipoti@anishmaddipoti-mlt folder-kg % git commit -m "Initial commit: Folder Knowledge Graph tool"
[master (root-commit) 6036e18] Initial commit: Folder Knowledge Graph tool
 3 files changed, 252 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 go.mod
 create mode 100644 main.go
anishmaddipoti@anishmaddipoti-mlt folder-kg % git remote add origin https://github.com/athreesh/folder-kg
anishmaddipoti@anishmaddipoti-mlt folder-kg % git branch -M ~
~
~
~
~
~
~
~

~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
-- INSERT --                              0,1           All
~
~
~
~
~
~
~
~
~
-- INSERT --                                                                                            0,1           All
main
anishmaddipoti@anishmaddipoti-mlt folder-kg % git push -u origin main
remote: Repository not found.
fatal: repository 'https://github.com/athreesh/folder-kg/' not found
anishmaddipoti@anishmaddipoti-mlt folder-kg % git remote add origin https://github.com/athreesh/folder-kg
error: remote origin already exists.
anishmaddipoti@anishmaddipoti-mlt folder-kg % git push -u origin main
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 12 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 2.50 KiB | 2.50 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To https://github.com/athreesh/folder-kg
 * [new branch]      main -> main
branch 'main' set up to track 'origin/main'.
anishmaddipoti@anishmaddipoti-mlt folder-kg % git status
On branch main
        Content  string
                Files:  make(map[string]*FileInfo),
        from := simple.Node(kg.NodeID[file1])
        to := simple.Node(kg.NodeID[file2])
        kg.Graph.SetWeightedEdge(simple.WeightedEdge{F: from, T: to, W: weight})
}

        if err != nil {
                return nil, err
        }

        }, nil
                        if path1 != path2 {
                                similarity, err := calculateSimilarity(ctx, llm, file1, file2, systemPrompt)
                                if err != nil {
                                        return err
                                }
}

        })
package main
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~

# Folder Knowledge Graph

# Folder Knowledge Graph
~
~
~
~
~
~
~
~
~
~
-- INSERT --                                                                                            0,1           All
You are an AI assistant specializing in analyzing files in a folder on a lapto and creating knowledge graphs. Your task isYour branch is up to date with 'origin/main'.

nothing to commit, working tree clean
anishmaddipoti@anishmaddipoti-mlt folder-kg % touch README.md
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim README.md
anishmaddipoti@anishmaddipoti-mlt folder-kg % git add .
anishmaddipoti@anishmaddipoti-mlt folder-kg % git commit -m "add README"
[main 2e2ddf6] add README
 1 file changed, 141 insertions(+)
 create mode 100644 README.md
anishmaddipoti@anishmaddipoti-mlt folder-kg % git push
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 12 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (3/3), 1.20 KiB | 1.20 MiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To https://github.com/athreesh/folder-kg
   6036e18..2e2ddf6  main -> main
anishmaddipoti@anishmaddipoti-mlt folder-kg % touch main.go
anishmaddipoti@anishmaddipoti-mlt folder-kg % ls
README.md	go.mod		main.go
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim main.go
anishmaddipoti@anishmaddipoti-mlt folder-kg % touch system_prompt.txt
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim system_prompt.txt
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim go.mod
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim main.go
anishmaddipoti@anishmaddipoti-mlt folder-kg % git status
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   main.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	system_prompt.txt

no changes added to commit (use "git add" and/or "git commit -a")
anishmaddipoti@anishmaddipoti-mlt folder-kg % vim README.md
anishmaddipoti@anishmaddipoti-mlt folder-kg % clear
go mod init github.
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~
~module github.com/athreesh/folder-kg

go 1.22

require (
    github.com/tmc/langchaingo v0.1.8
    gonum.org/v1/gonum v0.14.0
)
~
~                                                  1,20          All
