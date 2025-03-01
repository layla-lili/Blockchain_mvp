pub struct Transaction {
    inputs: Vec,
    outputs: Vec,
    timestamp: u64,
    signature: Option,
    hash: Hash,
}

pub struct TransactionInput {
    previous_output: OutPoint,
    script_sig: ScriptSig,
    sequence: u32,
}

pub struct TransactionOutput {
    value: u64,
    script_pubkey: ScriptPubKey,
}