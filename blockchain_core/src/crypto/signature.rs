pub struct Signature {
    // Signature data
}

pub fn sign(private_key: &PrivateKey, data: &[u8]) -> Signature { ... }
pub fn verify(public_key: &PublicKey, data: &[u8], signature: &Signature) -> bool { ... }