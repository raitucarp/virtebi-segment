# Virtebi Segment
virtebi algorithm segmentation: split multiple joined words with corpus parser in golang

Implementation Virtebi Algorithm by split joined words, into readable words in golang

It was begin, when I searching for solution for parsing domain names, into readable string.
And I found the solution http://stackoverflow.com/questions/195010/how-can-i-split-multiple-joined-words

# How to use
First, you have to get the module
```
go get github.com/raitucarp/virtebi-segment
```

Download the corpus https://github.com/raitucarp/corpus/raw/master/corpus.txt
Place the corpus to anywhere do you want

Import the module
```
package main

import (
    "github.com/raitucarp/virtebi-segment"
)

func main() {
    c := virtebi.NewCorpus()
    c.LoadCorpus("/home/path/to/the/corpus.txt")
    
    segment, prob := c.Segment("thisisatext")
    
    // this is a text
    fmt.Println(segment) 
}
```

# TODO
- Better documentation and Readme.
- Improving code quality.
- Fix some bugs.


# License
The MIT License (MIT)

Copyright (c) 2015 Ribhararnus Pracutiar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.