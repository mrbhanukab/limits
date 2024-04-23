# Limits Simulator

A Limits Simulator is a command-line tool that helps users visualize the limits of mathematical functions as a variable approaches a specific value.

![Watch the video](https://i.ytimg.com/vi/ZRb3XgcyKiI/hqdefault.jpg?sqp=-oaymwE2CNACELwBSFXyq4qpAygIARUAAIhCGAFwAcABBvABAfgB_gmAAtAFigIMCAAQARhyIEkoNzAP&rs=AOn4CLAzqUmIYDjE99lh4IMgA_aSTCmhDw)


Video:- https://www.youtube.com/watch?v=ZRb3XgcyKiI

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Supported Operating Systems](#supported-operating-systems)
- [Limits of "Limits Simulator"](#limits-of-limits-simulator)
- [Setup & Installation](#setup--installation)
- [Usage](#usage)
- [License](#license)
- [Credits](#credits)

## Introduction

The Limits Simulator is a tool designed to help users understand the concept of limits in calculus. It allows users to input a mathematical function, a value for the variable, and a range around the variable to visualize the behavior of the function as the variable approaches the given value.

## Features

- Allows users to input a mathematical function and a value for the variable.
- Visualizes the behavior of the function as the variable approaches the given value.
- Supports a range around the variable for visualization.

## Technologies Used

- **Language/Framework:** Go
- **Library:** [govaluate](https://github.com/Knetic/govaluate) (for evaluating mathematical expressions)

## Supported Operating Systems

![Windows](	https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![Mac](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=apple&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

<details>
<summary>Also Support</summary>
<ul>
<li>aix/ppc64</li>
<li>android/386</li>
<li>android/amd64</li>
<li>android/arm</li>
<li>android/arm64</li>
<li>darwin/amd64</li>
<li>darwin/arm64</li>
<li>dragonfly/amd64</li>
<li>freebsd/386</li>
<li>freebsd/amd64</li>
<li>freebsd/arm</li>
<li>freebsd/arm64</li>
<li>freebsd/riscv64</li>
<li>illumos/amd64</li>
<li>ios/amd64</li>
<li>ios/arm64</li>
<li>js/wasm</li>
<li>linux/386</li>
<li>linux/amd64</li>
<li>linux/arm</li>
<li>linux/arm64</li>
<li>linux/loong64</li>
<li>linux/mips</li>
<li>linux/mips64</li>
<li>linux/mips64le</li>
<li>linux/mipsle</li>
<li>linux/ppc64</li>
<li>linux/ppc64le</li>
<li>linux/riscv64</li>
<li>linux/s390x</li>
<li>netbsd/386</li>
<li>netbsd/amd64</li>
<li>netbsd/arm</li>
<li>netbsd/arm64</li>
<li>openbsd/386</li>
<li>openbsd/amd64</li>
<li>openbsd/arm</li>
<li>openbsd/arm64</li>
<li>openbsd/ppc64</li>
<li>plan9/386</li>
<li>plan9/amd64</li>
<li>plan9/arm</li>
<li>solaris/amd64</li>
<li>wasip1/wasm</li>
<li>windows/386</li>
<li>windows/amd64</li>
<li>windows/arm</li>
<li>windows/arm64</li>
</ul>
</details>

## Limits of "Limits Simulator"
- The function variable should be denoted by 'x'.
- Trigonometric functions (e.g., sin, cos, tan) are not supported.
- The value of X cannot be set to infinity (e.g., x->infinity is not supported).
- Only these arithmetic operators are supported: -, *, /, ** or ^, % (** and ^ refers to "take to the power of", for example 2**2 or 2^2 is equal to 4.)

## Setup & Installation

### Recommended Way for Most Users

- **Step 1:** Download the executable file for your operating system:
  - [Download for Windows (64-bit)](https://github.com/mrbhanukab/limits/raw/main/build/limits-windows-amd64.exe)
  - [Download for Windows (32-bit)](https://github.com/mrbhanukab/limits/raw/main/build/limits-windows-386.exe)
  - [Download for macOS (64-bit)](https://github.com/mrbhanukab/limits/raw/main/build/limits-darwin-amd64)
  - [Download for macOS (M1,M2, ...)](https://github.com/mrbhanukab/limits/raw/main/build/limits-darwin-arm64)
  - [Download for Linux (64-bit)](https://github.com/mrbhanukab/limits/raw/main/build/limits-linux-amd64)
  - [Download for Linux (32-bit)](https://github.com/mrbhanukab/limits/raw/main/build/limits-linux-386)

- **Step 2:** Once downloaded, navigate to the directory where the executable is located.
- **Step 3:** Give execute permission (only on linux and macos) to the executable by running the following command:
```bash
  chmod +x ./limits-linux-amd64
```
- **Step 4**: Open a terminal or command prompt in that directory.
- **Step 5:** Run the executable file by typing its name and pressing Enter.
  - For example, on Windows: `limits-windows-amd64.exe`
  - For example, on macOS: `./limits-darwin-amd64`
  - For example, on Linux: `./limits-linux-amd64`


### Build by Yourself

#### Prerequisites
- Go installed on your machine.

#### Steps
1. Clone the repository: `git clone https://github.com/mrbhanukab/limits.git`.
2. Navigate to the project folder: `cd limits`
3. Build the project: `go build`.
4. Run the executable file for your operating system located in the `build` directory.


## Usage

To use the Limits Simulator, follow these steps:

1. Enter a mathematical function when prompted.
2. Provide a value for the variable (x).
3. Specify the range around the variable for visualization.
4. The simulator will display the behavior of the function as the variable approaches the given value.

```bash
you@computer:/path/to/project/limits$ ./build/limits-linux-amd64

{Logo & Notes will visible here}

Enter the function:
2*x^5+8*x-5
Value of X (x->?):
0
Range around X (e.g., if you enter 2, it will use 2 -> -2):
4
Step size for evaluation:
1
⬇️ -4.00 -> -2085.00
⬇️ -3.00 -> -515.00
⬇️ -2.00 -> -85.00
⬇️ -1.00 -> -15.00
⬇️ 0.00 -> -5.00
🔽

0.00 -> -5.00

🔼
⬆️ 1.00 -> 5.00
⬆️ 2.00 -> 75.00
⬆️ 3.00 -> 505.00
⬆️ 4.00 -> 2075.00



 Limit of 2*x**5+8*x-5 as (x -> 4.00) approaches -5.00

```

## License

This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License). See the [LICENSE](./LICENSE) file for details.

## Credits

Limits Simulator is created by [Bhanuka Bandara](https://github.com/mrbhanukab). Feel free to contribute, report issues, or provide suggestions on the GitHub repository.

[![github](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/mrbhanukab)
[![linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/bhanuka-bandara-8a209420a)
[![Mastodon](https://img.shields.io/badge/Mastodon-563acc?style=for-the-badge&logo=Mastodon&logoColor=white)](https://www.kaggle.com/bhanukabandara)
[![Kaggle](https://img.shields.io/badge/Kaggle-20BEFF?style=for-the-badge&logo=Kaggle&logoColor=white)](https://mastodon.social/@mrbhanuka)
