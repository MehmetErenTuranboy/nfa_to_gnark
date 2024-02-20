# NFA to Gnark

`nfa_to_gnark` is a tool designed to transform Non-Finite Automata (NFA) into circuits compatible with the Gnark framework. This project focuses on the logical aspects of generating these circuits and provides outputs for both verification and circuit testing within the Gnark framework.

## Features

- **NFA to Circuit Conversion**: Efficiently converts NFA structures into Gnark-compatible circuits.
- **Verification Output**: Generates verification artifacts to ensure the correctness of the converted circuits.
- **Circuit Testing**: Facilitates testing of the generated circuits within the Gnark framework.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (version 1.15 or later)
- Gnark library installed and configured

### Installation

Clone the repository to your local machine using:

```bash
git clone git@github.com:MehmetErenTuranboy/nfa_to_gnark.git
```

### Generate ZK-regex Circuits
Following command will generate circuits with verification setup to output folder
```go
go run main.go
```

### Start verification
Please type following to see ouputed circuitcuit with verification setup.
```bash
cd output
```

If you want to start verfication please type following command
```go
go run circuit.go
```
