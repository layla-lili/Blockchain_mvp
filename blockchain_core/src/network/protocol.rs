pub struct Protocol {
    peers: HashMap,
    // Other protocol state
}

impl Protocol {
    pub fn new() -> Self { ... }
    pub fn broadcast_block(&self, block: Block) { ... }
    pub fn broadcast_transaction(&self, transaction: Transaction) { ... }
    pub fn handle_message(&mut self, peer_id: PeerId, message: Message) { ... }
}