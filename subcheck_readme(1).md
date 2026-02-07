# subcheck

**subcheck** runs your C/C++ project locally in a CAEN-accurate environment, so you can catch build and valgrind failures *before* wasting an autograder submission.

It removes the need to SSH into CAEN and deal with Duo just to check whether your code will compile.

---

## Requirements

- **Docker Desktop** (must be running) [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
- **Go 1.16+** (only if building from source)

---

## Quick Start

### Project setup

Your project must contain a `Makefile` with:

```makefile
EXECUTABLE = myprogram

release: build
	gcc -o $(EXECUTABLE) main.c
```

This matches what the autograder runs.

---

### Next steps

After setting up your Makefile, run the environment setup once:

```bash
subcheck env
```

This downloads and caches the CAEN environment locally.

Then, depending on what you want to check:

- Use `build` to verify your project compiles
- Use `valgrind` to run memory checks

---

### Commands

```bash
subcheck env          # Prepare or verify the CAEN environment (one-time)
subcheck build        # Run make release in the CAEN environment
subcheck valgrind     # Run valgrind on your executable
subcheck doctor       # Check Docker and system readiness
```

### Flags

```bash
subcheck --verbose build   # Show full Docker and compiler output
```

Output is quiet by default and focused on errors.

---

## Why subcheck exists

The CAEN autograder runs:

- Rocky Linux 9
- GCC 11.3.0
- glibc toolchain
- valgrind

Most student laptops run macOS with Apple Clang. With `-Werror` enabled, code that works locally can fail on CAEN and vice versa.

subcheck eliminates this mismatch by using the same OS and compiler as the autograder.

---

## How it works

subcheck embeds a Dockerfile configured to match CAEN.

On first run, it:

1. Hashes the Dockerfile
2. Builds the image if it does not already exist
3. Caches it locally

Subsequent runs are fast and require no network or SSH access.

---

## Common issues

**No Makefile found**\
Your project directory must contain a `Makefile` with a `release:` target.

**Missing EXECUTABLE variable**\
Add `EXECUTABLE = your_binary_name` to your Makefile.

**Platform mismatch warning**\
Normal on Apple Silicon. Docker handles x86\_64 emulation automatically.

**Docker daemon not running**\
Start Docker Desktop and try again.

---

##

