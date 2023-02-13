package main

import (
    "fmt"
    "os"

    core "github.com/iden3/go-iden3-core"
    "github.com/iden3/go-iden3-crypto/keccak256"
)

func main() {

  	schemaBytes, _ := os.ReadFile("./claim-schema/proof-of-dao-membership.json-ld")

    var sHash core.SchemaHash
    h := keccak256.Hash(schemaBytes, []byte("ProofOfDaoMembership"))

    copy(sHash[:], h[len(h)-16:])

    sHashHex, _ := sHash.MarshalText()

    fmt.Println(string(sHashHex))
    // c68f819206a2dd4401841e76b9d51e14
}