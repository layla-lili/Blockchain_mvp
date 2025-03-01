pub enum TransactionValidationError {
    InvalidSignature,
    InvalidHash,
    InsufficientFunds,
    DoubleSpend,
    // ...other validation errors
}