pub trait BlockchainStore {
    fn save_block(&self, block: &Block) -> Result;
    fn get_block(&self, hash: &Hash) -> Result<Option, StorageError>;
    fn save_transaction(&self, transaction: &Transaction) -> Result;
    fn get_transaction(&self, hash: &Hash) -> Result<Option, StorageError>;
    fn save_chain_state(&self, state: &ChainState) -> Result;
    fn get_chain_state(&self) -> Result<Option, StorageError>;
}

pub struct RocksDbStore {
    db: DB,
}

impl BlockchainStore for RocksDbStore {
    // Implementation using RocksDB
}