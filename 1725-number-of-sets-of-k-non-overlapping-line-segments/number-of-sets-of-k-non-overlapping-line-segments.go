package main

const MOD = 1000000007

// Calculates (base^exp) % MOD using binary exponentiation
func power(base, exp int) int {
	res := 1
	base = base % MOD
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

// Calculates modular multiplicative inverse using Fermat's Little Theorem
func modInverse(n int) int {
	return power(n, MOD-2)
}

func numberOfSets(n int, k int) int {
	totalPoints := n + k - 1
	choosePoints := 2 * k

	// If we don't have enough points, it's impossible
	if totalPoints < choosePoints {
		return 0
	}

	// Compute nCr % MOD optimization: nCr = n! / (r! * (n-r)!)
	num := 1
	den := 1

	for i := 0; i < choosePoints; i++ {
		num = (num * (totalPoints - i)) % MOD
		den = (den * (i + 1)) % MOD
	}

	return (num * modInverse(den)) % MOD
}