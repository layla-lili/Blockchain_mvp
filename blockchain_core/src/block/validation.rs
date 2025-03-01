pub enum BlockValidationError {
    InvalidHash,
    InvalidPreviousHash,
    InvalidMerkleRoot,
    InvalidProofOfWork,
    // ...other validation errors
}