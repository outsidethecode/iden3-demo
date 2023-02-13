package main

import (
    "fmt"
    "github.com/iden3/go-iden3-crypto/babyjub"

    "context"
    "math/big"

    merkletree "github.com/iden3/go-merkletree-sql"
    "github.com/iden3/go-merkletree-sql/db/memory"

    "time"
    "encoding/json"
    core "github.com/iden3/go-iden3-core"

)

// BabyJubJub key
func main() {

    // Identity Keypair -------------------------------------------------------------------------------
    // generate babyJubjub private key randomly
    babyJubjubPrivKey := babyjub.NewRandPrivKey()

    // generate public key from private key
    babyJubjubPubKey := babyJubjubPrivKey.Public()

    // print public key
    fmt.Println(babyJubjubPubKey)

    // Proof of membership example ----------------------------------------------------------------------
    ctx := context.Background()

    // Tree storage
    store := memory.NewMemoryStorage()

    // Generate a new MerkleTree with 32 levels
    mt, _ := merkletree.NewMerkleTree(ctx, store, 32)

    // Add a leaf to the tree with index 1 and value 10
    index1 := big.NewInt(1)
    value1 := big.NewInt(10)
    mt.Add(ctx, index1, value1)

    // Add another leaf to the tree
    index2 := big.NewInt(2)
    value2 := big.NewInt(15)
    mt.Add(ctx, index2, value2)

    // Proof of membership of a leaf with index 1
    proofExist, value, _ := mt.GenerateProof(ctx, index1, mt.Root())

    fmt.Println("Proof of membership:", proofExist.Existence)
    fmt.Println("Value corresponding to the queried index:", value)

    // Proof of non-membership of a leaf with index 4
    proofNotExist, _, _ := mt.GenerateProof(ctx, big.NewInt(4), mt.Root())

    fmt.Println("Proof of membership:", proofNotExist.Existence)

    // Generic claim example ----------------------------------------------------------------------

    // set claim expriation date to 2361-03-22T00:44:48+05:30
    t := time.Date(2361, 3, 22, 0, 44, 48, 0, time.UTC)

    // set schema
    ageSchema, _ := core.NewSchemaHashFromHex ("2e2d1c11ad3e500de68d7ce16a0a559e")  

    // define data slots
    birthday := big.NewInt(19960424)
    documentType := big.NewInt(1)   

    // set revocation nonce 
    revocationNonce := uint64(1909830690)

    // set ID of the claim subject
    id, _ := core.IDFromString("113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2")

    // create claim 
    claim, _ := core.NewClaim(ageSchema, core.WithExpirationDate(t), core.WithRevocationNonce(revocationNonce), core.WithIndexID(id), core.WithIndexDataInts(birthday, documentType))

    // transform claim from bytes array to json 
    claimToMarshal, _ := json.Marshal(claim)

    fmt.Println(string(claimToMarshal))


    // Auth claim example ----------------------------------------------------------------------

    authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")

    // Add revocation nonce. Used to invalidate the claim. This may be a random number in the real implementation.
    revNonce := uint64(1)

    // Create auth Claim 
    authClaim, _ := core.NewClaim(authSchemaHash,
    core.WithIndexDataInts(babyJubjubPubKey.X, babyJubjubPubKey.Y),
    core.WithRevocationNonce(revNonce))

    authClaimToMarshal, _ := json.Marshal(authClaim)

    fmt.Println(string(authClaimToMarshal))

}