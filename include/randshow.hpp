#pragma once

#include <chrono>
#include <cmath>
#include <cstdint>

namespace randshow {
class RNG {
public:
    // Next random 32-bit unsigned int.
    [[nodiscard]] virtual uint32_t next32() const = 0;

    // Next random 32-bit unsigned int from uniform distribution in [0, n) range.
    [[nodiscard]] virtual uint32_t next32_range(uint32_t n) const {
        return next32() % n;
    }

    // Next random 32-bit unsigned int from uniform distribution in [a, b) range.
    [[nodiscard]] virtual uint32_t next32_range(uint32_t a, uint32_t b) const {
        return b > a ? next32() % (b - a) + a : a;
    }

    // Next random 64-bit unsigned int.
    [[nodiscard]] virtual uint64_t next64() const = 0;

    // Next random 64-bit unsigned int from uniform distribution in [0, n) range.
    [[nodiscard]] virtual uint64_t next64_range(uint64_t n) const {
        return next64() % n;
    }

    // Next random 64-bit unsigned int from uniform distribution in [a, b) range.
    [[nodiscard]] virtual uint64_t next64_range(uint64_t a, uint64_t b) const {
        return b > a ? next64() % (b - a) + a : a;
    }

    // Next floating point number from standard uniform distribution i.e. (0, 1)
    // range.
    [[nodiscard]] virtual double_t nextfp() const {
        return static_cast<double_t>(next32()) / UINT32_MAX;
    }

    // Floating point uniform distribution in [a, b) range.
    [[nodiscard]] virtual double_t nextfp_range(double_t a, double_t b) const {
        return nextfp() * (b - a) + a;
    }

    // Balanced coin flip with 50% chance of heads and tails.
    [[nodiscard]] virtual bool coin_flip() const { return nextfp() < 0.5; }

    // Weighted coin flip. Parameter weight must be in (0, 1) range.
    // Weight smaller or equal than 0 will always yield false and
    // bigger or equal to 1 always yielding true.
    [[nodiscard]] virtual bool coin_flip(double_t weight) const {
        return nextfp() < weight;
    }
};

namespace lcg {
#define MUL_LCG 6458928179451363983ULL
#define INC_LCG 0UL
#define MOD_LCG ((1ULL << 63ULL) - 25ULL)

// LCG or Linear Congruential Generator has very low footprint and is very fast.
// It has m_modulo limit of generating pseudorandom numbers. LCGs are suitable
// for games or other trivial use-cases, but shouldn't be used for work
// requiring true random numbers.
//
// This LCG implementation requires 32 bytes of memory per instance.
class LCG : public RNG {
private:
    mutable uint64_t state = std::chrono::high_resolution_clock::now().time_since_epoch().count();
    uint64_t a_multiplier = MUL_LCG;
    uint64_t c_increment = INC_LCG;
    uint64_t m_modulo = MOD_LCG;

public:
    uint32_t next32() const override { return static_cast<uint32_t>(next64()); }

    uint64_t next64() const override {
        auto x = state;
        state = (a_multiplier * state + c_increment) % m_modulo;
        return x;
    }

    // Creates a new LCG engine with a, c, m parameters equal to the default
    // engine. Seed is current time.
    LCG() = default;

    // Creates a new LCG engine with a, c, m parameters equal to the default
    // engine with custom seed value. Meant for reproducibility.
    explicit LCG(uint64_t seed) : state(seed) {}

    // Creates a new LCG engine with custom a, c, m parameters and seed equal to
    // current time.
    LCG(uint64_t multiplier, uint64_t increment, uint64_t modulo)
            : a_multiplier(multiplier), c_increment(increment), m_modulo(modulo) {}

    // Creates a new LCG engine.
    LCG(uint64_t seed, uint64_t multiplier, uint64_t increment, uint64_t modulo)
            : state(seed), a_multiplier(multiplier), c_increment(increment), m_modulo(modulo) {}

    // Getter for state value.
    [[nodiscard]] uint64_t get_state() const { return state; }

    // Getter for state multiplier.
    [[nodiscard]] uint64_t get_multiplier() const { return a_multiplier; }

    // Getter for multiplied state increment.
    [[nodiscard]] uint64_t get_increment() const { return c_increment; }

    // Getter for incremented state modulo.
    [[nodiscard]] uint64_t get_modulo() const { return m_modulo; }
};

static const LCG default_engine = LCG();
} // namespace lcg

namespace pcg {
#define MUL_PCG 6364136223846793005UL
#define INC_PCG 1442695040888963407UL

// (high_bits XOR low_bits) of x.
inline static uint64_t xorshift32(uint64_t x) {
    auto high_bits = static_cast<uint32_t>(x >> 32U);
    auto low_bits = static_cast<uint32_t>(x);
    return high_bits ^ low_bits;
}

inline static uint32_t rotr32(uint64_t x, uint32_t rot) {
    return (x >> rot) | (x << ((-rot) & 31U));
}

class PCG : public RNG {
private:
    mutable uint64_t state = std::chrono::high_resolution_clock::now().time_since_epoch().count();
    const uint64_t a_multiplier = MUL_PCG;
    const uint64_t c_increment = INC_PCG;

    uint64_t advance() const {
        auto x = state;
        state = a_multiplier * state + c_increment;
        return x;
    }

public:
    // PCG-XSH-RR
    uint32_t next32() const override {
        uint64_t x = advance();
        auto rotation = static_cast<uint32_t>(x >> 59U);

        x = xorshift32(x);
        x = rotr32(x, rotation);
        return x;
    }

    // PCG-XSH-RR-RR
    uint64_t next64() const override {
        uint64_t x = advance();
        auto rotation = static_cast<uint32_t>(x >> 59U);

        // XSH
        auto high_bits = static_cast<uint32_t>(x >> 32U);
        auto low_bits = static_cast<uint32_t>(x);
        x = low_bits ^ high_bits;

        // RR-RR
        auto x_low = rotr32(x, rotation);
        auto x_high = rotr32(high_bits, x & 31U);
        x = (static_cast<uint64_t>(x_high) << 32U) | x_low;
        return x;
    }

    PCG() = default;

    explicit PCG(uint64_t seed) : state(seed) {}

    PCG(uint64_t multiplier, uint64_t increment)
            : a_multiplier(multiplier), c_increment(increment) {}

    PCG(uint64_t seed, uint64_t multiplier, uint64_t increment)
            : state(seed), a_multiplier(multiplier), c_increment(increment) {}

    // Getter for state value.
    [[nodiscard]] uint64_t get_state() const { return state; }

    // Getter for state multiplier.
    [[nodiscard]] uint64_t get_multiplier() const { return a_multiplier; }

    // Getter for multiplied state increment.
    [[nodiscard]] uint64_t get_increment() const { return c_increment; }
};

static const PCG default_engine = PCG();
} // namespace pcg
} // namespace randshow