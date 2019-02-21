# Testing in Go

`Go Test Files`:
Most go files in a project will also have a FILENAME_test.go file that goes along with it. This is that file's tests.


`Golden Files`:
It is useful to include a file that contains the expected output of your tests so that other developers can tell whether or not they're passing. This file is known as a golden file.

`Mocking`:
Sometimes using real objects in tests is impractical. In scenarios like this, we create mock data/objects to run through our tests. This is known as mocking.