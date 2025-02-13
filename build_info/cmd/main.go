package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("H")
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("No build info")
		return
	}
	fmt.Println("Module:", info.Main.Path)
	fmt.Println("Version:", info.Main.Version)
	fmt.Println("Sum:", info.Main.Sum)
	fmt.Println("GoVersion:", info.GoVersion)
	fmt.Println("Dependencies:")
	for _, dep := range info.Deps {
		fmt.Println("  ", dep.Path, dep.Version)
	}
	for _, s := range info.Settings {
		fmt.Println("  ", s.Key, s.Value)
	}
}
