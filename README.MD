
# WASM JDK

## How to test

1. Clone [openjdk/jdk21u](https://github.com/openjdk/jdk21u)
2. Ensure you have the dependencies setup (such as openjdk 21/20)
3. Build java source code into classes:
   ```bash
   git clone https://github.com/openjdk/jdk21u
   cd jdk21u
   bash configure
   make java
   ```
4. Extract the class files:
   ```bash
   mkdir /tmp/govm-test
   cd build/*-server-release/jdk/modules
   cp -r ./*/* /tmp/govm-test/ 2>/dev/null
   ```
5. Write and compile a java test file (can just have the main function), or copy the [existing test file](https://github.com/LiterMC/wasm-jdk/blob/main/jcls/testdata/Test.class) to `/tmp/govm-test/`
6. Run gova in the directory (where you placed the java classfiles):
   ```bash
   go install github.com/LiterMC/wasm-jdk/cmd/gova@latest
   gova Test
   ```
