moov-io/x9
===
[![GoDoc](https://godoc.org/github.com/moov-io/x9?status.svg)](https://godoc.org/github.com/moov-io/x9)
[![Build Status](https://travis-ci.com/moov-io/x9.svg?branch=master)](https://travis-ci.com/moov-io/x9)
[![Coverage Status](https://coveralls.io/repos/github/moov-io/x9/badge.svg?branch=master)](https://coveralls.io/github/moov-io/x9?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/x9)](https://goreportcard.com/report/github.com/moov-io/x9)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/x9/master/LICENSE)

X9’s Specifications for ICL (Image Cash ledger) to provide Check 21 services

Package `github.com/moov-io/x9` implements a file reader and writer for parsing [x9](https://en.wikipedia.org/wiki/Check_21_Act) files.

Docs: [docs.moov.io](http://docs.moov.io/en/latest/) | [api docs](https://api.moov.io/)

## Project Status

x9 is under active development, and should not be used in Production.  Please star the project if you are interested in its progress.

* The Library currently supports the reading and writing
	* Cash Letters
	* Bundles
	* Check Detail
	* Return Detail
	* Image Views

* Future Development
    * Increase Test Code Coverage to 95% +
    * Add Routing Number Summary Tests
    * Add Credit Items Functionality
    * Add User Record Functionality
    * Benchmarking and Profiling Tests

## Project Roadmap
* Please open an issue with a valid test file.
* Review the project issues for more detailed information

## Usage and tests
The following is a high level example of reading and writing an X9 file. 

### Read a file

```go
// open a file for reading or pass any io.Reader NewReader()
f, err := os.Open("name-of-your-x9-file.x9")
if err != nil {
	log.Panicf("Can not open local file: %s: \n", err)
}
r := x9.NewReader(f)
x9File, err := r.Read()
if err != nil {
	fmt.Printf("Issue reading file: %+v \n", err)
}
// ensure we have a validated file structure
if x9File.Validate(); err != nil {
	fmt.Printf("Could not validate entire read file: %v", err)
}

```

### Create a file

Explicitly create an X9 file with a Bundle of Forward Presentment Items and a Bundle of Return Items

 ```go
// Create a File
file := NewFile().SetHeader(mockFileHeader())

// Create CheckDetail
cd := mockCheckDetail()

//Add Check Detail
cd.AddCheckDetailAddendumA(mockCheckDetailAddendumA())
cd.AddCheckDetailAddendumB(mockCheckDetailAddendumB())
cd.AddCheckDetailAddendumC(mockCheckDetailAddendumC())

// Add ImageView
cd.AddImageViewDetail(mockImageViewDetail())
cd.AddImageViewData(mockImageViewData())
cd.AddImageViewAnalysis(mockImageViewAnalysis())

// Create Bundle
bundle := NewBundle(mockBundleHeader())

// Add Check Detail to Bundle
bundle.AddCheckDetail(cd)

// Create Return Detail
rd := mockReturnDetail()
rd.AddReturnDetailAddendumA(mockReturnDetailAddendumA())
rd.AddReturnDetailAddendumB(mockReturnDetailAddendumB())
rd.AddReturnDetailAddendumC(mockReturnDetailAddendumC())
rd.AddReturnDetailAddendumD(mockReturnDetailAddendumD())
rd.AddImageViewDetail(mockImageViewDetail())
rd.AddImageViewData(mockImageViewData())
rd.AddImageViewAnalysis(mockImageViewAnalysis())
returnBundle := NewBundle(mockBundleHeader())
returnBundle.AddReturnDetail(rd)

// Create CashLetter
cl := NewCashLetter(mockCashLetterHeader())
cl.AddBundle(bundle)
cl.AddBundle(returnBundle)
cl.Create()
file.AddCashLetter(cl)

// Create file
if err := file.Create(); err != nil {
	t.Errorf("%T: %s", err, err)
}
if err := file.Validate(); err != nil {
	t.Errorf("%T: %s", err, err)
}

````
Which will generate a well formed X9 flat file.

```text
0135T231380104121042882201810012313NCitadel           Wells Fargo        US     
100123138010412104288220181001201810012313IGA1      Contact Name  5558675552    
200123138010412104288220181001201810019999      1   01                          
25      123456789 031300012             555888100001000001              GD1Y030B
261121042882201810011              938383            01   Test Payee     Y10    
2711A             00340                                 CD Addendum B           
2801121042882201810011              Y10A                   0                    
501031300012201810010000000000000000000000000000000000000         0             
52121042882201810011 1              Sec Orig Name   Sec Auth Name   SECURE          0                00000    0000001 
542202222222             10222222222222                                         
70000700000010000000000010000000001                    0                        
200123138010412104288220181001201810019999      1   01                          
31031300012             55588810000100000A03G201810011               2B0        
321121042882201810011              938383            01   Test Payee     Y10    
33Payor Bank Name         1234567891              20181001Payor Account Name    
3411A             00340                                 RD Addendum C           
3501121042882201810011              Y10A                   0                    
501031300012201810010000000000000000000000000000000000000         0             
52121042882201810011 1              Sec Orig Name   Sec Auth Name   SECURE          0                00000    0000001 
542202222222             10222222222222                                         
70000800000010000000000000000000001                    0                        
900000020000001500000000200000000000002                  201810010              
9900000100000023000000150000000000200000                        0               
```

## Getting Help

 channel | info
 ------- | -------
[Project Documentation](http://docs.moov.io/en/latest/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and is not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher, but we ship the vendor directory in our repository.

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.

