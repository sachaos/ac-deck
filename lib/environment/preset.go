package environment

var Aliases = map[string]string {
	"python": "python3",
	"c++": "c++_gcc",
}

// NOTE: https://atcoder.jp/contests/language-test-202001
var Environments = map[string]*Environment{
	"c++_gcc": {
		Key:          "c++_gcc",
		Language:     "C++ (GCC 9.2.1)",
		SrcName:      "main.cpp",
		Template:     "internal/c++/main.cpp",
		LanguageCode: "4003",

		BuildCmd: "g++ -std=gnu++17 -O2 -o a.out main.cpp",
		Cmd:      "./a.out",
		CleanCmd: "rm ./a.out",

		DockerImage:      "ghcr.io/sachaos/atcoder-gcc:v1.0.0",
		BuildCmdOnDocker: "g++ -std=gnu++17 -Wall -Wextra -I /opt/boost/boost_1_72_0 -L /opt/boost/boost_1_72_0 -I /opt/ac-library -O2 -o a.out main.cpp",
		CmdOnDocker:      "./a.out",
	},
	"go": {
		Key:          "go",
		Language:     "Go (1.14.1)",
		SrcName:      "main.go",
		Template:     "internal/go/main.go",
		LanguageCode: "4026",

		BuildCmd: "go build -o ./a.out main.go",
		Cmd:      "./a.out",
		CleanCmd: "rm ./a.out",

		DockerImage:      "docker.io/library/golang:1.14.1",
		BuildCmdOnDocker: "go build -o a.out main.go",
		CmdOnDocker:      "./a.out",
	},
	"python3": {
		Key:          "python3",
		Language:     "Python3 (3.8.2)",
		SrcName:      "main.py",
		Template:     "internal/python3/main.py",
		LanguageCode: "4006",

		Cmd: "python3 -B main.py",

		CmdOnDocker: "python -B main.py",
		DockerImage: "ghcr.io/sachaos/atcoder-python3:v1.0.0",
	},
	"rust": {
		Key:          "rust",
		Language:     "Rust (1.42.0)",
		SrcName:      "main.rs",
		Template:     "internal/rust/main.rs",
		LanguageCode: "4050",

		BuildCmd: "cargo build --release --offline --quiet",
		Cmd:      "./target/release/rust",
		CleanCmd: "rm ./target/release/rust",

		WorkingDir: "/src",
		SrcDir: "/src/src",

		DockerImage:      "ghcr.io/sachaos/atcoder-rust:v1.0.0",
		BuildCmdOnDocker: "cargo build --release --offline --quiet",
		CmdOnDocker:      "./target/release/rust",
	},
}

