pub struct RpcServer {
    blockchain: Arc<Mutex>,
    // Other server state
}

impl RpcServer {
    pub fn new(blockchain: Arc<Mutex>) -> Self { ... }
    pub fn start(&self, address: &str) -> Result { ... }
}