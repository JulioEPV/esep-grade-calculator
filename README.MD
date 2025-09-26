# ESEP Grade Calculator

This library is a simple utility to calculate grades based on weighted averages of essays, assignments, and exams. 

1. Fork the project from the following repository:  
https://github.com/audgster/esep-grade-calculator

2. Description of the files 

a.	grade_calculator.go  - Define the structure of the calculator. Define its types as Assignment, Exam, and Essay. Define grade structure. Define functions to get the final grade, add a grade, calculate the numerical grade, and compute the average grade.
b.	 grade_calculator_test.go – There are three test cases. The first one expects a grade of A, the second expects a B, and the final test expects a F. Three grades are added for each case. With these values, it calculates the final grade and compares it with the expected values.


Runtime requirements:
Golang 1.24+

3. Install Golang from go.dev
    a. Open the website
    b. Click on the download option
    c. If you are using Windows download go1.24.4.windows-amd64.msi
    d. Follow the steps to install go
    e. open a terminal and confirm you have go with go version

4. If you are using VS Code Download the Go extension

5. Run unit tests:
```
cd grade_calculator/
go test ./...
```

6. Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```