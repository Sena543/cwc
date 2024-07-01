### Word Count Utility

This command-line utility provides functionality to analyze text files based on user-specified options: counting bytes, words, lines, and characters. It supports both file input and input via standard input (stdin).

### Usage

To use the utility, execute the compiled binary with the following command-line arguments:

```bash
./cwc [-c] [-w] [-l] [-m] [file]
```

-   **-c**: Count bytes.
-   **-w**: Count words.
-   **-l**: Count lines.
-   **-m**: Count characters (alternative to -c).

If no options (-c, -w, -l, -m) are provided, the utility will default to outputting the count of lines, words, and characters, along with the filename.

### Example Usage

1. Count bytes in a file:

    ```bash
    ./cwc -c filename.txt
    ```

2. Count words in a file:

    ```bash
    ./cwc -w filename.txt
    ```

3. Count lines in a file:

    ```bash
    cwc -l filename.txt
    ```

4. Count characters in a file (alternative option):

    ```bash
    ./cwc -m filename.txt
    ```

5. Count lines, words, and characters in a file:
    ```bash
    ./cwc filename.txt
    ```

### Features

-   **ReadFile(filePath string) []byte**: Reads the content of the file specified by `filePath`.
-   **GetFileBytes(fileData []byte) int64**: Returns the number of bytes in `fileData`.
-   **GetLineCount(fileData []byte) int64**: Returns the number of lines in `fileData`.
-   **GetCharacterCount(fileData []byte) int64**: Returns the number of characters in `fileData`.
-   **GetWordCount(fileData []byte) int64**: Returns the number of words in `fileData`.
-   **CountEntry()**: Main function that handles command-line arguments, reads file content, and prints counts based on user options.

### Notes

-   The utility handles both direct file input (`cwc [options] filename`) and piped input (`cat filename | cwc [options]`).
-   If no filename is provided and stdin is empty or not piped, the utility attempts to use the last argument as a filename.
-   Error handling for file reading and command-line argument parsing is minimal and may need extension depending on deployment needs.

