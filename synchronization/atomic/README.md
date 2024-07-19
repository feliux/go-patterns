# Atomic

Atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

These functions require great care to be used correctly. Except for special, low-level applications, synchronization is better done with channels or the facilities of the sync package. Share memory by communicating; don't communicate by sharing memory.

Examples [here](../locks/spinLock.go) and [here](../../concurrency/barrier-pattern/memBarriers.go)
