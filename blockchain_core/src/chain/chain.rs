pub struct Blockchain {
    blocks: Vec,
    state: ChainState,
}

impl Blockchain {
    pub fn new() -> Self { ... }
    pub fn add_block(&mut self, block: Block) -> Result { ... }
    pub fn get_latest_block(&self) -> &Block { ... }
    pub fn get_block_by_hash(&self, hash: &Hash) -> Option { ... }
    pub fn get_block_by_height(&self, height: u64) -> Option { ... }
}