pub struct BlockHeader {
    version: u32,
    previous_hash: Hash,
    merkle_root: Hash,
    timestamp: u64,
    difficulty: u32,
    nonce: u64,
}