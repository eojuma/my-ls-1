
# my-ls 📂

A high-performance, custom implementation of the Unix `ls` command written in **Go**. This project focuses on low-level file system interactions, bitwise permission manipulation, and custom sorting algorithms without relying on forbidden standard libraries.

## 🚀 Overview

`my-ls` replicates the core functionality of the standard Unix `ls` utility. It allows users to list directory contents with support for essential flags, providing a deep dive into how Go interfaces with the Unix underlying system calls.

## 🛠 Features & Flags

We have implemented the following core flags to match the behavior of the system `ls`:

* **`-l` (Long Listing):** Displays detailed information including file permissions, number of links, owner, group, size, and last modification time.
* **`-a` (All):** Includes directory entries whose names begin with a dot (`.`), which are hidden by default.
* **`-r` (Reverse):** Reverses the order of the sort.
* **`-t` (Time):** Sorts by modification time, newest first.
* **`-R` (Recursive) [Bonus]:** Recursively lists subdirectories encountered.

## 🏗 Project Architecture

To maintain high standards and meet program constraints, the project is structured around:
1.  **Custom Sorting:** Implementation of sorting algorithms to handle ASCII and time-based ordering without the `sort` package.
2.  **Permission Parsing:** Manual conversion of `file mode` bits into the standard `-rwxr-xr-x` string format.
3.  **Metadata Extraction:** Utilizing `syscall` and `os` packages to fetch UID/GID and link counts for the `-l` flag.

## ⚠️ Constraints & Allowed Packages

This implementation strictly adheres to the project requirements. The use of `os/exec` is **prohibited**. We rely solely on the following allowed packages:

| Category | Packages |
| :--- | :--- |
| **Core** | `fmt`, `os`, `errors`, `io/fs` |
| **System** | `syscall`, `os/user` |
| **Data** | `strconv`, `strings`, `time`, `math/rand` |

## 💻 Usage

Compile the project and run it against any directory:

```bash
# Basic usage
go run . [directory_path]

# Using flags
go run . -latR [directory_path]
```

## 👥 Contributors
* **Evans Juma**
* **Silas Lelei**
* **Magret Kerubo**

---
