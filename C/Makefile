TARGETS = primes checkPrimes
CC_LINUX = gcc
CC_WINDOWS = x86_64-w64-mingw32-gcc

SRCS = primes.c checkPrimes.c
OBJS = $(SRCS:.c=.o)

# Detect the operating system
ifeq ($(OS),Windows_NT)
    OS_TARGET = windows
else
    OS_TARGET = linux
endif

# Set the compiler based on the detected OS
ifeq ($(OS_TARGET),windows)
    CC = $(CC_WINDOWS)
    EXECUTABLE_EXT = .exe
else
    CC = $(CC_LINUX)
    EXECUTABLE_EXT =
endif

# Compilation flags
CFLAGS = -Wall -Wextra

all: $(TARGETS)

$(TARGETS): $(OBJS)
	$(CC) $(CFLAGS) -o $@$(EXECUTABLE_EXT) $@.o

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

clean:
	rm -f $(TARGETS) $(OBJS)
