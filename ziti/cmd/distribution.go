package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"
)

type DistributionManifest struct {
    Name         string                               `yaml:"name"`
    Version      string                               `yaml:"version"`
    Author       string                               `yaml:"author"`
    Components   map[string]map[string]interface{}    `yaml:"components"`
    Distribution map[string]interface{}               `yaml:"distribution"`
    Notes        string                               `yaml:"notes"`
}

func NewDistributionCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "distribution",
        Short: "Show the Petar Petrov OpenZiti distribution manifest",
        Long:  "Displays the full distribution.yaml manifest for the Petar Petrov OpenZiti Distribution.",
        Run: func(cmd *cobra.Command, args []string) {
            data, err := os.ReadFile("distribution.yaml")
            if err != nil {
                fmt.Printf("Error reading distribution.yaml: %v\n", err)
                return
            }

            var manifest DistributionManifest
            if err := yaml.Unmarshal(data, &manifest); err != nil {
                fmt.Printf("Error parsing distribution.yaml: %v\n", err)
                return
            }

            fmt.Printf("\n📦 %s\n", manifest.Name)
            fmt.Printf("Version: %s\n", manifest.Version)
            fmt.Printf("Author: %s\n\n", manifest.Author)

            fmt.Println("Components:")
            for name, comp := range manifest.Components {
                fmt.Printf("  - %s:\n", name)
                for k, v := range comp {
                    fmt.Printf("      %s: %v\n", k, v)
                }
            }

            fmt.Println("\nDistribution Metadata:")
            for k, v := range manifest.Distribution {
                fmt.Printf("  %s: %v\n", k, v)
            }

            fmt.Printf("\nNotes:\n%s\n\n", manifest.Notes)
        },
    }

    return cmd
}

