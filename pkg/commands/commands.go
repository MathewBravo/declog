package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/MathewBravo/declog/pkg/decisions"
)

type Command byte

const (
	Init Command = iota
	NewLog
	NewDecision
	Unrecognized
)

type Flag byte

const (
	Todo Flag = iota
)

func CommandExecutor(cmd Command) {
	switch cmd {
	case Init:
		initialize()
	case NewLog:
		newLog()
	case NewDecision:
		newDecision()
	case Unrecognized:
		unrecognized()
	}
}

func newDecision() {
	title, desc := getTitleDesc()
	decision := decisions.NewDecision(title, desc)
	decisions.AppendLocalDecision(decision)
}



func unrecognized() {
	fmt.Println("Unknown Command")
}

func newLog() {
	logDir := "declog"
	decLogJson := filepath.Join(logDir, "declog.json")
	if err := os.Mkdir(logDir, os.ModePerm); err != nil {
		fmt.Println("This directory already has an initialized declog repo")
		os.Exit(1)
	}
	file, err := os.Create(decLogJson)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	emptyDecisions := decisions.Decisions{}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(emptyDecisions); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Local Decision Log Created")
	fmt.Println("Begin adding decisions with with [nd || newdec] command")
}

func initialize() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	declogDir := filepath.Join(homeDir, ".declog")
	if err := os.Mkdir(homeDir+"/.declog", os.ModePerm); err != nil {
		fmt.Println("This computer already has an initialized declog repo")
		os.Exit(1)
	}

	masterlogPath := filepath.Join(declogDir, "masterlog.json")
	file, err := os.Create(masterlogPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	emptyDecisions := decisions.Decisions{}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(emptyDecisions); err != nil {
		log.Fatal(err)
	}
}
