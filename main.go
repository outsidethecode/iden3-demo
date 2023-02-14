package main

import (
    "context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
    "io/ioutil"

	"github.com/iden3/go-circuits"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
	merkletree "github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-merkletree-sql/v2/db/memory"
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

    // Three identity trees example ----------------------------------------------------------------------

    // Create empty Claims tree
    clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Create empty Revocation tree
    ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Create empty Roots tree
    rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Get the Index and the Value of the authClaim
    hIndex, hValue, _ := authClaim.HiHv()

    // add auth claim to claims tree with value hValue at index hIndex
    clt.Add(ctx, hIndex, hValue)

    // print the roots
    fmt.Println(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())

    // Identity state example ----------------------------------------------------------------------

    // calculate Identity State as a hash of the three roots
    state, _ := merkletree.HashElems(
        clt.Root().BigInt(),
        ret.Root().BigInt(),
        rot.Root().BigInt())

    fmt.Println("Identity State:", state)

    identityId, _ := core.IdGenesisFromIdenState(core.TypeDefault, state.BigInt())

    fmt.Println("ID:", identityId)

    // Signature Claim Issuance ----------------------------------------------------------------------

    // claimIndex, claimValue := claim.RawSlots()
    // indexHash2, _ := poseidon.Hash(core.ElemBytesToInts(claimIndex[:]))
    // valueHash2, _ := poseidon.Hash(core.ElemBytesToInts(claimValue[:]))

    // // Poseidon Hash the indexHash and the valueHash together to get the claimHash
    // claimHash, _ := merkletree.HashElems(indexHash2, valueHash2)

    // // Sign the claimHash with the private key of the issuer
    // claimSignature := babyJubjubPrivKey.SignPoseidon(claimHash.BigInt())
    // fmt.Println("Claim Signature:", claimSignature)

    // Signature via Mercle Tree ----------------------------------------------------------------------

    // GENESIS STATE:

    // 1. Generate Merkle Tree Proof for authClaim at Genesis State
    authMTPProof, _, _ := clt.GenerateProof(ctx, hIndex, nil)
    authMTPProofToMarshal, _ := json.Marshal(authMTPProof)
    fmt.Println("------ authMTPProof ------ ")
    fmt.Println(string(authMTPProofToMarshal))


    
    // 2. Generate the Non-Revocation Merkle tree proof for the authClaim at Genesis State
    authNonRevMTPProof, _, _ := ret.GenerateProof(ctx, new(big.Int).SetUint64(revNonce), nil)

    // Snapshot of the Genesis State
    genesisTreeState := circuits.TreeState{
        State:          state,
        ClaimsRoot:     clt.Root(),
        RevocationRoot: ret.Root(),
        RootOfRoots:    rot.Root(),
    }
    // STATE 1:

    // Before updating the claims tree, add the claims tree root at Genesis state to the Roots tree.
    rot.Add(ctx, clt.Root().BigInt(), big.NewInt(0))

    // Create a new random claim
    schemaHex := hex.EncodeToString([]byte("myAge_test_claim"))
    schema, _ := core.NewSchemaHashFromHex(schemaHex)

    code := big.NewInt(51)

    newClaim, _ := core.NewClaim(schema, core.WithIndexDataInts(code, nil))

    // Get hash Index and hash Value of the new claim
    hi, hv, _ := newClaim.HiHv()

    // Add claim to the Claims tree
    clt.Add(ctx, hi, hv)

    // Fetch the new Identity State
    newState, _ := merkletree.HashElems(
        clt.Root().BigInt(),
        ret.Root().BigInt(),
        rot.Root().BigInt())

    newTreeState := circuits.TreeState{
        State:          newState,
        ClaimsRoot:     clt.Root(),
        RevocationRoot: ret.Root(),
        RootOfRoots:    rot.Root(),
    }
    

    authNewStateMTPProof, _, _ := clt.GenerateProof(ctx, hi, nil)

    // Sign a message (hash of the genesis state + the new state) using your private key
    hashOldAndNewStates, _ := poseidon.Hash([]*big.Int{state.BigInt(), newState.BigInt()})

    signature := babyJubjubPrivKey.SignPoseidon(hashOldAndNewStates)

    // Generate state transition inputs
    stateTransitionInputs := circuits.StateTransitionInputs{
        ID:                identityId,
        OldTreeState:      genesisTreeState,
        NewTreeState:      newTreeState,
        IsOldStateGenesis: true,
        AuthClaim: authClaim,
        AuthClaimIncMtp: authMTPProof,
        AuthClaimNonRevMtp: authNonRevMTPProof,
        AuthClaimNewStateIncMtp: authNewStateMTPProof,
        Signature: signature,
    }

    // Perform marshalling of the state transition inputs
    inputBytes, _ := stateTransitionInputs.InputsMarshal()

    fmt.Println("string(inputBytes)")
    fmt.Println(string(inputBytes))
    fmt.Printf("%+v\n", stateTransitionInputs)

    file, _ := json.Marshal(string(inputBytes))
 
	_ = ioutil.WriteFile("./compiled-circuits/stateTransition/stateTransition_js/input.json", file, 0644)

}