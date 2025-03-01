pub enum Message {
    NewBlock(Block),
    NewTransaction(Transaction),
    GetBlocks(Vec),
    GetData(Vec),
    // Other message types
}