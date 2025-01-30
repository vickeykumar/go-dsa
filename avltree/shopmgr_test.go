package avltree

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "testing"
)

func TestProcessQueries(t *testing.T) {
    // Open the input file
    inputFile, err := os.Open("input.txt")
    if err != nil {
        t.Fatalf("Error opening input file: %v", err)
    }
    defer inputFile.Close()

    // Open the output file
    outputFile, err := os.Create("output.txt")
    if err != nil {
        t.Fatalf("Error creating output file: %v", err)
    }
    defer outputFile.Close()

    // Create an AVL tree
    var root *UserNode

    scanner := bufio.NewScanner(inputFile)
    // Read the number of queries
    scanner.Scan()
    numQueries, err := strconv.Atoi(scanner.Text())
    if err != nil {
        t.Fatalf("Error parsing number of queries: %v", err)
    }

    // Process each query
    for i := 0; i < numQueries && scanner.Scan(); i++ {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) < 2 {
            t.Fatalf("Invalid input format in line %d: %s", i+2, line)
        }
        queryType, err := strconv.Atoi(parts[0])
        if err != nil {
            t.Fatalf("Error parsing query type in line %d: %v", i+2, err)
        }
        if queryType == 1 {
            if len(parts) != 3 {
                t.Fatalf("Invalid input format in line %d: %s", i+2, line)
            }
            start, err := strconv.Atoi(parts[1])
            if err != nil {
                t.Fatalf("Error parsing start time in line %d: %v", i+2, err)
            }
            end, err := strconv.Atoi(parts[2])
            if err != nil {
                t.Fatalf("Error parsing end time in line %d: %v", i+2, err)
            }
            // Insert user
            root = InsertUser(root, start, end)
        } else if queryType == 2 {
            if len(parts) != 2 {
                t.Fatalf("Invalid input format in line %d: %s", i+2, line)
            }
            timestamp, err := strconv.Atoi(parts[1])
            if err != nil {
                t.Fatalf("Error parsing timestamp in line %d: %v", i+2, err)
            }
            // Count users at timestamp
            count := CountUsers(root, timestamp)
            // Write the count to the output file
            _, err = fmt.Fprintf(outputFile, "%d\n", count)
            if err != nil {
                t.Fatalf("Error writing output for line %d: %v", i+2, err)
            }
        } else {
            t.Fatalf("Invalid query type in line %d: %s", i+2, line)
        }
    }

    if err := scanner.Err(); err != nil {
        t.Fatalf("Error reading input file: %v", err)
    }
}

func TestCountUsers(t *testing.T) {
    // Create an AVL tree and insert some users
    var root *UserNode
    root = InsertUser(root, 10, 20)
    root = InsertUser(root, 5, 15)
    root = InsertUser(root, 25, 30)

    // Test CountUsers function
    testCases := []struct {
        timestamp int
        expectedCount int
    }{
        {0, 0},   // No users at timestamp 0
        {12, 2},  // Two users at timestamp 12
        {18, 1},  // One user at timestamp 18
        {40, 0},  // No users at timestamp 40
    }

    for _, tc := range testCases {
        count := CountUsers(root, tc.timestamp)
        if count != tc.expectedCount {
            t.Errorf("CountUsers(root, %d) = %d; want %d", tc.timestamp, count, tc.expectedCount)
        } else {
            t.Logf("CountUsers(root, %d) = %d; want %d", tc.timestamp, count, tc.expectedCount)
        }
    }
}
