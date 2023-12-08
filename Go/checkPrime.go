package main

import (
	"fmt"
	"strconv"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func all(vals []bool) bool {
	for _, v := range vals {
		if !v {
			return false
		}
	}
	return true
}

func main() {
	numStr := ""
	fmt.Print("Enter a number: ")
	fmt.Scanln(&numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	if num == 1 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 2 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 3 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 4 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 5 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 6 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 7 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 8 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 9 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 10 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 11 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 12 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 13 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 14 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 15 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 16 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 17 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 18 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 19 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 20 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 21 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 22 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 23 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 24 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 25 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 26 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 27 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 28 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 29 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 30 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 31 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 32 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 33 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 34 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 35 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 36 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 37 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 38 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 39 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 40 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 41 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 42 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 43 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 44 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 45 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 46 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 47 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 48 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 49 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 50 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 51 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 52 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 53 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 54 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 55 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 56 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 57 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 58 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 59 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 60 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 61 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 62 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 63 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 64 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 65 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 66 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 67 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 68 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 69 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 70 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 71 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 72 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 73 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 74 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 75 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 76 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 77 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 78 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 79 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 80 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 81 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 82 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 83 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 84 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 85 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 86 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 87 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 88 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 89 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 90 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 91 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 92 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 93 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 94 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 95 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 96 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 97 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 98 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 99 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 100 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 101 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 102 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 103 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 104 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 105 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 106 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 107 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 108 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 109 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 110 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 111 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 112 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 113 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 114 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 115 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 116 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 117 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 118 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 119 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 120 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 121 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 122 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 123 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 124 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 125 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 126 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 127 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 128 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 129 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 130 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 131 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 132 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 133 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 134 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 135 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 136 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 137 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 138 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 139 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 140 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 141 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 142 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 143 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 144 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 145 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 146 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 147 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 148 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 149 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 150 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 151 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 152 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 153 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 154 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 155 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 156 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 157 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 158 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 159 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 160 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 161 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 162 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 163 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 164 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 165 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 166 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 167 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 168 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 169 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 170 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 171 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 172 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 173 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 174 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 175 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 176 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 177 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 178 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 179 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 180 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 181 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 182 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 183 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 184 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 185 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 186 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 187 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 188 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 189 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 190 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 191 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 192 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 193 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 194 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 195 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 196 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 197 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 198 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 199 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 200 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 201 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 202 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 203 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 204 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 205 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 206 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 207 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 208 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 209 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 210 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 211 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 212 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 213 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 214 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 215 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 216 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 217 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 218 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 219 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 220 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 221 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 222 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 223 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 224 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 225 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 226 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 227 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 228 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 229 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 230 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 231 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 232 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 233 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 234 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 235 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 236 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 237 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 238 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 239 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 240 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 241 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 242 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 243 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 244 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 245 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 246 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 247 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 248 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 249 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 250 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 251 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 252 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 253 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 254 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 255 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 256 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 257 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 258 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 259 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 260 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 261 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 262 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 263 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 264 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 265 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 266 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 267 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 268 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 269 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 270 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 271 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 272 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 273 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 274 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 275 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 276 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 277 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 278 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 279 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 280 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 281 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 282 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 283 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 284 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 285 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 286 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 287 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 288 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 289 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 290 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 291 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 292 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 293 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 294 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 295 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 296 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 297 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 298 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 299 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 300 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 301 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 302 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 303 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 304 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 305 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 306 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 307 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 308 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 309 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 310 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 311 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 312 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 313 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 314 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 315 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 316 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 317 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 318 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 319 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 320 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 321 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 322 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 323 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 324 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 325 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 326 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 327 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 328 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 329 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 330 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 331 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 332 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 333 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 334 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 335 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 336 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 337 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 338 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 339 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 340 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 341 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 342 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 343 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 344 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 345 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 346 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 347 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 348 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 349 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 350 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 351 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 352 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 353 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 354 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 355 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 356 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 357 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 358 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 359 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 360 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 361 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 362 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 363 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 364 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 365 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 366 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 367 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 368 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 369 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 370 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 371 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 372 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 373 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 374 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 375 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 376 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 377 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 378 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 379 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 380 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 381 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 382 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 383 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 384 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 385 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 386 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 387 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 388 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 389 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 390 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 391 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 392 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 393 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 394 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 395 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 396 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 397 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 398 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 399 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 400 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 401 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 402 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 403 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 404 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 405 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 406 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 407 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 408 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 409 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 410 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 411 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 412 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 413 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 414 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 415 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 416 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 417 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 418 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 419 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 420 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 421 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 422 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 423 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 424 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 425 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 426 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 427 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 428 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 429 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 430 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 431 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 432 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 433 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 434 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 435 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 436 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 437 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 438 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 439 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 440 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 441 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 442 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 443 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 444 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 445 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 446 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 447 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 448 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 449 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 450 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 451 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 452 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 453 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 454 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 455 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 456 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 457 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 458 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 459 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 460 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 461 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 462 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 463 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 464 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 465 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 466 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 467 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 468 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 469 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 470 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 471 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 472 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 473 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 474 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 475 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 476 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 477 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 478 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 479 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 480 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 481 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 482 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 483 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 484 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 485 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 486 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 487 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 488 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 489 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 490 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 491 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 492 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 493 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 494 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 495 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 496 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 497 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 498 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 499 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 500 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 501 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 502 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 503 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 504 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 505 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 506 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 507 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 508 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 509 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 510 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 511 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 512 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 513 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 514 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 515 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 516 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 517 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 518 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 519 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 520 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 521 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 522 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 523 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 524 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 525 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 526 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 527 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 528 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 529 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 530 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 531 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 532 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 533 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 534 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 535 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 536 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 537 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 538 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 539 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 540 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 541 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 542 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 543 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 544 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 545 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 546 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 547 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 548 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 549 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 550 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 551 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 552 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 553 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 554 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 555 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 556 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 557 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 558 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 559 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 560 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 561 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 562 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 563 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 564 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 565 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 566 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 567 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 568 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 569 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 570 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 571 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 572 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 573 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 574 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 575 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 576 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 577 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 578 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 579 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 580 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 581 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 582 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 583 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 584 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 585 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 586 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 587 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 588 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 589 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 590 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 591 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 592 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 593 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 594 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 595 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 596 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 597 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 598 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 599 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 600 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 601 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 602 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 603 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 604 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 605 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 606 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 607 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 608 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 609 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 610 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 611 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 612 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 613 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 614 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 615 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 616 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 617 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 618 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 619 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 620 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 621 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 622 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 623 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 624 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 625 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 626 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 627 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 628 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 629 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 630 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 631 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 632 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 633 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 634 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 635 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 636 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 637 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 638 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 639 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 640 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 641 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 642 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 643 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 644 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 645 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 646 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 647 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 648 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 649 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 650 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 651 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 652 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 653 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 654 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 655 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 656 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 657 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 658 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 659 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 660 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 661 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 662 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 663 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 664 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 665 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 666 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 667 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 668 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 669 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 670 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 671 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 672 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 673 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 674 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 675 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 676 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 677 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 678 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 679 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 680 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 681 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 682 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 683 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 684 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 685 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 686 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 687 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 688 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 689 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 690 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 691 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 692 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 693 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 694 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 695 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 696 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 697 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 698 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 699 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 700 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 701 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 702 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 703 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 704 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 705 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 706 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 707 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 708 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 709 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 710 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 711 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 712 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 713 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 714 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 715 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 716 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 717 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 718 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 719 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 720 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 721 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 722 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 723 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 724 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 725 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 726 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 727 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 728 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 729 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 730 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 731 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 732 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 733 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 734 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 735 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 736 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 737 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 738 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 739 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 740 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 741 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 742 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 743 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 744 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 745 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 746 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 747 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 748 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 749 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 750 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 751 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 752 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 753 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 754 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 755 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 756 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 757 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 758 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 759 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 760 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 761 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 762 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 763 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 764 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 765 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 766 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 767 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 768 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 769 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 770 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 771 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 772 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 773 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 774 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 775 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 776 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 777 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 778 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 779 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 780 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 781 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 782 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 783 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 784 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 785 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 786 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 787 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 788 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 789 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 790 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 791 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 792 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 793 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 794 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 795 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 796 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 797 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 798 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 799 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 800 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 801 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 802 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 803 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 804 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 805 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 806 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 807 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 808 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 809 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 810 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 811 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 812 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 813 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 814 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 815 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 816 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 817 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 818 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 819 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 820 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 821 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 822 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 823 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 824 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 825 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 826 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 827 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 828 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 829 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 830 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 831 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 832 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 833 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 834 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 835 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 836 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 837 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 838 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 839 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 840 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 841 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 842 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 843 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 844 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 845 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 846 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 847 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 848 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 849 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 850 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 851 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 852 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 853 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 854 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 855 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 856 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 857 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 858 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 859 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 860 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 861 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 862 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 863 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 864 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 865 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 866 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 867 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 868 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 869 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 870 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 871 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 872 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 873 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 874 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 875 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 876 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 877 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 878 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 879 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 880 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 881 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 882 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 883 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 884 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 885 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 886 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 887 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 888 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 889 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 890 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 891 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 892 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 893 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 894 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 895 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 896 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 897 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 898 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 899 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 900 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 901 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 902 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 903 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 904 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 905 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 906 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 907 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 908 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 909 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 910 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 911 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 912 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 913 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 914 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 915 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 916 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 917 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 918 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 919 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 920 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 921 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 922 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 923 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 924 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 925 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 926 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 927 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 928 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 929 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 930 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 931 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 932 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 933 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 934 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 935 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 936 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 937 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 938 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 939 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 940 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 941 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 942 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 943 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 944 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 945 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 946 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 947 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 948 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 949 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 950 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 951 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 952 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 953 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 954 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 955 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 956 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 957 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 958 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 959 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 960 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 961 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 962 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 963 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 964 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 965 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 966 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 967 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 968 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 969 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 970 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 971 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 972 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 973 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 974 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 975 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 976 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 977 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 978 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 979 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 980 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 981 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 982 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 983 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 984 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 985 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 986 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 987 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 988 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 989 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 990 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 991 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 992 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 993 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 994 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 995 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 996 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 997 {
		fmt.Printf("%d is a prime number.\n", num)
	}
	if num == 998 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 999 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	if num == 1000 {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}
