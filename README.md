# min
I'm building a simple API testing framework in Go. Each test will be a json file with a request and (optionally) an expected result. The resulting response will be displayed, and if there is an expected result then a comparison will be made and displayed. The tests will also be timed.

## Install
```bash
git clone https://github.com/gabecatalfo/min.git
cd min
./scripts/install.sh
```

## Uninstall
```bash
./scripts/uninstall.sh
```

## Commands
```bash
min               # run all tests/requests recursively in order
min -p            # run tests/requests in parallel
min -w            # watch for files changes (run tests/requests on change)
min -v            # verbose (displays full output on failure AND success)
min <file.json>   # run the test/request of a single file
min <directory>   # run all tests/requests recursively from said directory
```
