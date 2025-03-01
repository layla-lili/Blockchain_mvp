pub trait Serializable {
    fn serialize(&self) -> Result<Vec, SerializationError>;
    fn deserialize(bytes: &[u8]) -> Result where Self: Sized;
}

impl Serializable for Block {
    // Implementation using serde
}

impl Serializable for Transaction {
    // Implementation using serde
}