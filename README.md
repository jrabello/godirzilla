# Godirzilla

![Godirzilla Logo](<LOGO_URL>)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
  
## Introduction

Godirzilla is a powerful command-line tool designed to make your life easier when managing multiple repositories. It enables you to run a single command across multiple directories simultaneously, saving your precious time and keeping you efficient and productive.

## Features

- Concurrent execution: Run commands in multiple directories concurrently.
- Custom groups: Define your own groups of directories.
- Easy to use: Intuitive command line interface.
- Real-time reporting: Get instant feedback on the success or failure of your commands.

## Installation

1. Clone this repository: `git clone https://github.com/YourUsername/Godirzilla.git`.
2. Navigate to the cloned repository: `cd Godirzilla`.
3. Install the application: `go install`.

## Usage

```bash
gdz run -g my-group "git status"
