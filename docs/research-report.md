# Toy Blockchain Research Report

## 1. Introduction

This report presents experiments performed on the custom Go-based toy blockchain implementation.

The purpose is to evaluate:
- blockchain tamper detection
- proof-of-work performance
- design decisions and limitations

---

## 2. Tamper Evidence Experiment

### Method

A valid blockchain was created and mined. After mining, a transaction inside an existing block was modified.

### Before Tampering

Validation result:

true


### After Tampering

Validation result:

block 1 hash mismatch


### Observation

Changing transaction data modifies the block hash. Since the stored hash no longer matches the recalculated hash, the validation process detects the modification.

---

## 3. Proof-of-Work Difficulty Experiment

| Difficulty | Attempts | Duration |
|------------|----------|----------|
| 1 | 14 | 0s |
| 2 | 47 | 0s |
| 3 | 11031 | 17.47ms |
| 4 | 51179 | 41.36ms |

### Observation

Higher difficulty requires more computational effort. The number of attempts increases unpredictably because the correct nonce is found through trial and error.

---

## 4. Hashing Design

The block hash is generated using SHA-256.

Fields included:

1. Block index
2. Timestamp
3. Transaction sender
4. Transaction recipient
5. Transaction amount
6. Previous block hash
7. Nonce

The hash field itself is excluded to avoid circular dependency.

---

## 5. Validation Design

Blockchain validation checks:

- Stored hash matches recalculated hash
- Previous hash links are correct
- Proof-of-work requirement is satisfied
- Block ordering is valid
- Timestamp ordering is consistent

---

## 6. Limitations Compared to Production Blockchains

This implementation does not include:

1. Peer-to-peer networking
2. Digital signatures
3. Merkle trees

---

## 7. Future Improvements

Possible extensions:

- Add digital signatures for transaction authentication
- Add Merkle tree transaction summaries
- Implement peer-to-peer consensus