pub type Hash = [u8; 32];

pub fn hash_block(block: &Block) -> Hash { ... }
pub fn hash_transaction(transaction: &Transaction) -> Hash { ... }
pub fn merkle_root(transactions: &[Transaction]) -> Hash { ... }