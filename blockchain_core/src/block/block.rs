pub struct Block {
    header: BlockHeader,
    transactions: Vec,
    hash: Hash,
}

impl Block {
    pub fn new(header: BlockHeader, transactions: Vec) -> Self { ... }
    pub fn genesis() -> Self { ... }
    pub fn validate(&self, previous_block: &Block) -> Result { ... }
    pub fn hash(&self) -> Hash { ... }
}