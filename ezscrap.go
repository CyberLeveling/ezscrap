package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
	"net/http"
	"regexp"
	 "io/ioutil"
	"crypto/tls"

)



func main() {
	go spinner(100 * time.Millisecond)
        // Open the input file
        inputFile, err := os.Open("input.txt")
        if err != nil {
                fmt.Printf("Error opening input file: %v\n", err)
                return
        }
        defer inputFile.Close()

        // Open the output file
        outputFile, err := os.Create("output.txt")
        if err != nil {
                fmt.Printf("Error creating output file: %v\n", err)
                return
        }
        defer outputFile.Close()

        // Create a scanner to read the input file line by line
        scanner := bufio.NewScanner(inputFile)

        // Read each line of the input file
        for scanner.Scan() {
                line := scanner.Text()

                // Write the line to the output file
                _, err = outputFile.WriteString(line + "\n")
                if err != nil {
                        fmt.Printf("Error writing to output file: %v\n", err)
                        return
                }

                // Write the line again to the output file
                _, err = outputFile.WriteString(line + "\n")
                if err != nil {
                        fmt.Printf("Error writing to output file: %v\n", err)
                        return
                }
        }

        // Check for errors while reading the input file
        if err := scanner.Err(); err != nil {
                fmt.Printf("Error reading input file: %v\n", err)
        }
	timer()
}




func timer() {
        // Declare a variable to store the elapsed time
        var elapsedTime time.Duration

        // Start the timer
        start := time.Now()

        // Call the function you want to measure the elapsed time for
        addingdata()

        // Stop the timer and calculate the elapsed time
        elapsedTime = time.Since(start)

        // Print the elapsed time
        fmt.Printf("\n\033[32m\033 Function 1 took:\033[0m \033[33m\033[4m%s\033[0m \033[32m\033 to complete.\033[0m", elapsedTime)

}




func addingdata() {

	// Open the input file
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a map to store the lines we've seen so far
	seenLines := make(map[string]bool)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a new file to save the modified lines to
	outFile, err := os.Create("modified_output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	// Create a writer to write to the output file
	writer := bufio.NewWriter(outFile)

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if we've seen this line before
		if seenLines[line] {
			// If we have, add "http://" to the beginning of the line
			line = "http://" + line
		}

		// Write the line to the output file
		fmt.Fprintln(writer, line)

		// Mark this line as seen
		seenLines[line] = true
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Flush the writer to make sure all the data is written to the file

	writer.Flush()
	timer1()
}



func timer1() {
        // Declare a variable to store the elapsed time
        var elapsedTime time.Duration

        // Start the timer
        start := time.Now()

        // Call the function you want to measure the elapsed time for
        moredata()

        // Stop the timer and calculate the elapsed time
        elapsedTime = time.Since(start)

        // Print the elapsed time
        fmt.Printf("\n\033[32m\033 Function 2 took:\033[0m \033[33m\033[4m%s\033[0m \033[32m\033 to complete.\033[0m", elapsedTime)

}




func moredata() {
	
	// Open the input file

	file, err := os.Open("modified_output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new file to save the modified lines to
	outFile, err := os.Create("final_output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	// Create a writer to write to the output file
	writer := bufio.NewWriter(outFile)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains "http://"
		if !strings.Contains(line, "http://") {
			// If it doesn't, add "https://" to the beginning of the line
			line = "https://" + line
		}

		// Write the line to the output file
		fmt.Fprintln(writer, line)
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Flush the writer to make sure all the data is written to the file
	writer.Flush()
	timer2()
}


func timer2() {
        // Declare a variable to store the elapsed time
        var elapsedTime time.Duration

        // Start the timer
        start := time.Now()

        // Call the function you want to measure the elapsed time for
        fuzz()

        // Stop the timer and calculate the elapsed time
        elapsedTime = time.Since(start)

        // Print the elapsed time
        fmt.Printf("\n\033[32m\033 Function 3 took:\033[0m \033[33m\033[4m%s\033[0m \033[32m\033 to complete.\033[0m", elapsedTime)

}




func fuzz() {
	
	// create http client with custom CheckRedirect function
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// open urls.txt file
	file, err := os.Open("final_output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create slice to store emails
	var urls []string
	var emails []string
	var numRequests int // Add a counter for the number of requests

	// Create a channel that sends a value every 1 second

	// loop through file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// gather URLs from source code
		re := regexp.MustCompile(`(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
		matches := re.FindAllString(scanner.Text(), -1)

		// loop through URLs
		for _, url := range matches {
			// make get request
			resp, err := client.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			// read response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			// gather emails from source code
			re := regexp.MustCompile(`(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
			matches := re.FindAllString(string(body), -1)

			// gather emails from source code
            emailRegex := regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)
            emailMatches := emailRegex.FindAllString(string(body), -1)

			// append emails to emails slice
			urls = append(urls, matches...)
			emails = append(emails, emailMatches...)

			// Increment the counter for the number of requests
			numRequests++

			
		}
	}

	// write emails to emails.txt
    f, err := os.Create("emails.txt")
    if err != nil {
            panic(err)
    }
    defer f.Close()

    for _, email := range emails {
        	fmt.Fprintln(f, email) 
    }
	// write urls to urls1.txt
	f, err = os.Create("urls1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, url := range urls {
		fmt.Fprintln(f, url)
	}


	timer3()
}

func timer3() {
        // Declare a variable to store the elapsed time
        var elapsedTime time.Duration

        // Start the timer
        start := time.Now()

        // Call the function you want to measure the elapsed time for
        sort()

        // Stop the timer and calculate the elapsed time
        elapsedTime = time.Since(start)

        // Print the elapsed time
        fmt.Printf("\n\033[32m\033 Function 4 took:\033[0m \033[33m\033[4m%s\033[0m \033[32m\033 to complete.\033[0m", elapsedTime)

}



func sort() {
        // Open the input file
        inputFile, err := os.Open("urls1.txt")
        if err != nil {
                fmt.Println(err)
                return
        }
        defer inputFile.Close()

        // Create a map to store the lines we've already seen
        seen := make(map[string]bool)

        // Open the output file
        outputFile, err := os.Create("sort1.txt")
        if err != nil {
                fmt.Println(err)
                return
        }
        defer outputFile.Close()

        // Create a scanner to read the input file line by line
        scanner := bufio.NewScanner(inputFile)
        for scanner.Scan() {
                line := scanner.Text()

                // If we haven't seen this line before, write it to the output file
                if !seen[line] {
                        _, err := outputFile.WriteString(line + "\n")
                        if err != nil {
                                fmt.Println(err)
                                return
                        }
                        seen[line] = true
                }
        }

        if err := scanner.Err(); err != nil {
                fmt.Println(err)
                return
        }

		timer4()
}


func timer4() {
        // Declare a variable to store the elapsed time
        var elapsedTime time.Duration

        // Start the timer
        start := time.Now()

        // Call the function you want to measure the elapsed time for
        checkaliveHosts()

        // Stop the timer and calculate the elapsed time
        elapsedTime = time.Since(start)

        // Print the elapsed time
        fmt.Printf("\n\033[32m\033 Function 5 took:\033[0m \033[33m\033[4m%s\033[0m \033[32m\033 to complete.\033[0m", elapsedTime)

}




func checkaliveHosts() {
	
	// Open file for reading
	file, err := os.Open("sort1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create new file for writing
	newFile, err := os.Create("alive_hosts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()

		response, err := http.Get(url)
		if err == nil {
			statusCode := response.StatusCode
			if statusCode == 200 || statusCode == 301 || statusCode == 302 || statusCode == 403 {
				newFile.WriteString(url + "\n")
				response.Body.Close()
		fuzzagain()
			}
		}
	}
}



func fuzzagain() {

    // Open the input file
    // create http client with custom CheckRedirect function
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            if req.TLS == nil || len(req.TLS.VerifiedChains) == 0 {
                // certificate is not valid, skip redirect
                f, err := os.OpenFile("nocertificates.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
                if err != nil {
                    return err
                }
                defer f.Close()
                if _, err := f.WriteString(req.URL.String() + "\n"); err != nil {
                    return err
                }
                return http.ErrUseLastResponse
            }
            // follow redirect and return nil
            return nil
        },
    }
    // open urls.txt file
    file, err := os.Open("alive_hosts.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // create slice to store emails
    var urls []string
    var emails []string
    var numRequests int // Add a counter for the number of requests

    // Create a channel that sends a value every 1 second


    // loop through file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // gather URLs from source code
        re := regexp.MustCompile(`(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
        matches := re.FindAllString(scanner.Text(), -1)

        // loop through URLs
        for _, url := range matches {
            // make get request
            resp, err := client.Get(url)
            if err != nil {
                panic(err)
            }
            defer resp.Body.Close()

            // read response body
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                panic(err)
            }

            // gather emails from source code
            re := regexp.MustCompile(`(http|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
            matches := re.FindAllString(string(body), -1)

			// gather emails from source code
            emailRegex := regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)
            emailMatches := emailRegex.FindAllString(string(body), -1)

			// append emails to emails slice
			urls = append(urls, matches...)
			emails = append(emails, emailMatches...)

			// Increment the counter for the number of requests
			numRequests++
			
			// Print the current number of requests every 1 second

			
		}
	}

    // write emails to emails.txt
    f, err := os.OpenFile("emails2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    for _, email := range emails {
        if _, err := f.WriteString(email + "\n"); err != nil {
            panic(err)
        }
    }

    // write urls to urls.txt
    f, err = os.OpenFile("urls2lvl.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    for _, url := range urls {
        if _, err := f.WriteString(url + "\n"); err != nil {
            panic(err)
        }
    }

    // print elapsed time

	sortz()
}


func sortz() {
	// Open the input file
	inputFile, err := os.Open("urls2lvl.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a map to store the lines we've already seen
	seen := make(map[string]bool)

	// Open the output file
	outputFile, err := os.Create("sort2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// If we haven't seen this line before, write it to the output file
		if !seen[line] {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
			seen[line] = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}


	remove3()
}


func remove3() {

	// Open the file for reading
	file, err := os.Open("sort1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create the output files
	pdfFile, err := os.Create("pdf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdfFile.Close()

	// Create the output files
	jsFile, err := os.Create("js.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsFile.Close()


	jpgFile, err := os.Create("jpg.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jpgFile.Close()

	svgFile, err := os.Create("svg.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer svgFile.Close()

	jsonFile, err := os.Create("json.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()


	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains ".pdf"
		if strings.Contains(line, ".pdf") {
			// Write the line to the pdf.txt file
			_, err := pdfFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Check if the line contains ".pdf"
		if strings.Contains(line, ".json") {
			// Write the line to the pdf.txt file
			_, err := jsonFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}



		// Check if the line contains ".jpg"
		if strings.Contains(line, ".js") {
			// Write the line to the jpg.txt file
			_, err := jsFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}


		// Check if the line contains ".jpg"
		if strings.Contains(line, ".jpg") {
			// Write the line to the jpg.txt file
			_, err := jpgFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Check if the line contains ".svg"
		if strings.Contains(line, ".svg") {
			// Write the line to the svg.txt file
			_, err := svgFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}



	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	remove2()
}



func remove2() {

	// Open the file for reading
	file, err := os.Open("sort2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create the output files
	pdfFile, err := os.Create("pdf2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdfFile.Close()

	// Create the output files
	jsFile, err := os.Create("js1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsFile.Close()


	jpgFile, err := os.Create("jpg2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jpgFile.Close()

	svgFile, err := os.Create("svg2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer svgFile.Close()

	jsonFile, err := os.Create("json2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()


	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains ".pdf"
		if strings.Contains(line, ".pdf") {
			// Write the line to the pdf.txt file
			_, err := pdfFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Check if the line contains ".pdf"
		if strings.Contains(line, ".json") {
			// Write the line to the pdf.txt file
			_, err := jsonFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}



		// Check if the line contains ".jpg"
		if strings.Contains(line, ".js") {
			// Write the line to the jpg.txt file
			_, err := jsFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}


		// Check if the line contains ".jpg"
		if strings.Contains(line, ".jpg") {
			// Write the line to the jpg.txt file
			_, err := jpgFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		// Check if the line contains ".svg"
		if strings.Contains(line, ".svg") {
			// Write the line to the svg.txt file
			_, err := svgFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}



	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	grepfirstword()
}



func grepfirstword() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)

	// Create the output file
	outputFile, err := os.Create("output4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Loop over the lines in the input file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words
		words := strings.Split(line, " ")

		// Write the first word to the output file if it exists
		if len(words) > 0 {
			firstWord := strings.Split(words[0], ".")[0]
			_, err := outputFile.WriteString(firstWord + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// Check for any errors that occurred while reading or writing the files
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	if err := outputFile.Sync(); err != nil {
		fmt.Println(err)
	
	}
	readAndSearch1()
}




func readAndSearch1() error {
    // Open the first file
    firstFileHandle, err := os.Open("output4.txt")
    if err != nil {
        return fmt.Errorf("error opening first file: %v", err)
    }
    defer firstFileHandle.Close()

    // Read the first line of the first file
    firstScanner := bufio.NewScanner(firstFileHandle)
    firstScanner.Scan()
    firstLine := firstScanner.Text()

    // Open the second file
    secondFileHandle, err := os.Open("sort1.txt")
    if err != nil {
        return fmt.Errorf("error opening second file: %v", err)
    }
    defer secondFileHandle.Close()

    // Open the output file
    outputFileHandle, err := os.Create("active1.txt")
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFileHandle.Close()

    // Read the second file line by line and search for a match with the first line
    secondScanner := bufio.NewScanner(secondFileHandle)
    for secondScanner.Scan() {
        secondLine := secondScanner.Text()
        if strings.Contains(secondLine, firstLine) {
            // If a match is found, write the second line to the output file
            _, err := outputFileHandle.WriteString(secondLine + "\n")
            if err != nil {
                return fmt.Errorf("error writing to output file: %v", err)
            }
        }
    }

    if err := secondScanner.Err(); err != nil {
        return fmt.Errorf("error reading second file: %v", err)
    }
	readAndSearch2()
    return nil

}


func readAndSearch2() error {
    // Open the first file
    firstFileHandle, err := os.Open("output4.txt")
    if err != nil {
        return fmt.Errorf("error opening first file: %v", err)
    }
    defer firstFileHandle.Close()

    // Read the first line of the first file
    firstScanner := bufio.NewScanner(firstFileHandle)
    firstScanner.Scan()
    firstLine := firstScanner.Text()

    // Open the second file
    secondFileHandle, err := os.Open("sort2.txt")
    if err != nil {
        return fmt.Errorf("error opening second file: %v", err)
    }
    defer secondFileHandle.Close()

    // Open the output file
    outputFileHandle, err := os.Create("active2.txt")
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFileHandle.Close()

    // Read the second file line by line and search for a match with the first line
    secondScanner := bufio.NewScanner(secondFileHandle)
    for secondScanner.Scan() {
        secondLine := secondScanner.Text()
        if strings.Contains(secondLine, firstLine) {
            // If a match is found, write the second line to the output file
            _, err := outputFileHandle.WriteString(secondLine + "\n")
            if err != nil {
                return fmt.Errorf("error writing to output file: %v", err)
            }
        }
    }

    if err := secondScanner.Err(); err != nil {
        return fmt.Errorf("error reading second file: %v", err)
    }
	urlzboss()
    return nil

}



func urlzboss() {
	// Open the file "javascript.txt" in read-only mode
	file1, err := os.Open("active1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()

	// Open the file "javascript1.txt" in read-only mode
	file2, err := os.Open("active2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	// Create a new file "final_javascript" in write-only mode
	final, err := os.Create("final_active_urlstarget.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer final.Close()

	// Use a bufio.Scanner to read the contents of file1 and file2
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	// Use a bufio.Writer to write the contents of file1 and file2 to final
	writer := bufio.NewWriter(final)
	for scanner1.Scan() {
		text := scanner1.Text()
		_, err := writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for scanner2.Scan() {
		text := scanner2.Text()
		_, err := writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println(err)
		
			return
		
		}
	
	readAndSearch5()
	}

}

func readAndSearch5() {
	// Open the file "javascript.txt" in read-only mode
	file1, err := os.Open("js.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()

	// Open the file "javascript1.txt" in read-only mode
	file2, err := os.Open("js1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	// Create a new file "final_javascript" in write-only mode
	final, err := os.Create("final-js.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer final.Close()

	// Use a bufio.Scanner to read the contents of file1 and file2
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	// Use a bufio.Writer to write the contents of file1 and file2 to final
	writer := bufio.NewWriter(final)
	for scanner1.Scan() {
		text := scanner1.Text()
		_, err := writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	for scanner2.Scan() {
		text := scanner2.Text()
		_, err := writer.WriteString(text + "\n")
		if err != nil {
			fmt.Println(err)
			
			return
		}
	readAndSearch9()
	
	}

}



func readAndSearch9() error {
    // Open the first file
    firstFileHandle, err := os.Open("output4.txt")
    if err != nil {
        return fmt.Errorf("error opening first file: %v", err)
    }
    defer firstFileHandle.Close()

    // Read the first line of the first file
    firstScanner := bufio.NewScanner(firstFileHandle)
    firstScanner.Scan()
    firstLine := firstScanner.Text()

    // Open the second file
    secondFileHandle, err := os.Open("emails2.txt")
    if err != nil {
        return fmt.Errorf("error opening second file: %v", err)
    }
    defer secondFileHandle.Close()

    // Open the output file
    outputFileHandle, err := os.Create("final-emails2.txt")
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFileHandle.Close()

    // Read the second file line by line and search for a match with the first line
    secondScanner := bufio.NewScanner(secondFileHandle)
    for secondScanner.Scan() {
        secondLine := secondScanner.Text()
        if strings.Contains(secondLine, firstLine) {
            // If a match is found, write the second line to the output file
            _, err := outputFileHandle.WriteString(secondLine + "\n")
            if err != nil {
                return fmt.Errorf("error writing to output file: %v", err)
            }
        }
    }

    if err := secondScanner.Err(); err != nil {
        return fmt.Errorf("error reading second file: %v", err)
    }
	sortza()
    return nil
	
}



func sortza() {
	// Open the input file
	inputFile, err := os.Open("final-js.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a map to store the lines we've already seen
	seen := make(map[string]bool)

	// Open the output file
	outputFile, err := os.Create("JSfilesorted.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// If we haven't seen this line before, write it to the output file
		if !seen[line] {
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}
			seen[line] = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
	readAndSearch77()
}



func readAndSearch77() error {
    // Open the first file
    firstFileHandle, err := os.Open("output4.txt")
    if err != nil {
        return fmt.Errorf("error opening first file: %v", err)
    }
    defer firstFileHandle.Close()

    // Read the first line of the first file
    firstScanner := bufio.NewScanner(firstFileHandle)
    firstScanner.Scan()
    firstLine := firstScanner.Text()

    // Open the second file
    secondFileHandle, err := os.Open("JSfilesorted.txt")
    if err != nil {
        return fmt.Errorf("error opening second file: %v", err)
    }
    defer secondFileHandle.Close()

    // Open the output file
    outputFileHandle, err := os.Create("finaljstarget.txt")
    if err != nil {
        return fmt.Errorf("error creating output file: %v", err)
    }
    defer outputFileHandle.Close()

    // Read the second file line by line and search for a match with the first line
    secondScanner := bufio.NewScanner(secondFileHandle)
    for secondScanner.Scan() {
        secondLine := secondScanner.Text()
        if strings.Contains(secondLine, firstLine) {
            // If a match is found, write the second line to the output file
            _, err := outputFileHandle.WriteString(secondLine + "\n")
            if err != nil {
                return fmt.Errorf("error writing to output file: %v", err)
            }
        }
    }

    if err := secondScanner.Err(); err != nil {
        return fmt.Errorf("error reading second file: %v", err)
    }
    return nil
	
}





func spinner(delay time.Duration) {
	for {
		for _, r := range `◐◓◑◒` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
