// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/.. -lfilecoin
#cgo pkg-config: ${SRCDIR}/../filecoin.pc
#include "../filecoin.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// FilBLSSignature as declared in filecoin-ffi/filecoin.h:49
type FilBLSSignature struct {
	Inner          [96]byte
	refa2ac09ba    *C.fil_BLSSignature
	allocsa2ac09ba interface{}
}

// FilAggregateResponse as declared in filecoin-ffi/filecoin.h:56
type FilAggregateResponse struct {
	Signature      FilBLSSignature
	refb3efa36d    *C.fil_AggregateResponse
	allocsb3efa36d interface{}
}

// FilClearCacheResponse as declared in filecoin-ffi/filecoin.h:61
type FilClearCacheResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	refa9a80400    *C.fil_ClearCacheResponse
	allocsa9a80400 interface{}
}

// FilFinalizeTicketResponse as declared in filecoin-ffi/filecoin.h:67
type FilFinalizeTicketResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	Ticket         [32]byte
	refb370fa86    *C.fil_FinalizeTicketResponse
	allocsb370fa86 interface{}
}

// FilCandidate as declared in filecoin-ffi/filecoin.h:74
type FilCandidate struct {
	SectorId             uint64
	PartialTicket        [32]byte
	Ticket               [32]byte
	SectorChallengeIndex uint64
	refaa02d4eb          *C.fil_Candidate
	allocsaa02d4eb       interface{}
}

// FilGenerateCandidatesResponse as declared in filecoin-ffi/filecoin.h:81
type FilGenerateCandidatesResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	CandidatesPtr  []FilCandidate
	CandidatesLen  uint
	ref1b9cd8d8    *C.fil_GenerateCandidatesResponse
	allocs1b9cd8d8 interface{}
}

// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filecoin.h:87
type FilGenerateDataCommitmentResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	CommD          [32]byte
	ref87da7dd9    *C.fil_GenerateDataCommitmentResponse
	allocs87da7dd9 interface{}
}

// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filecoin.h:98
type FilGeneratePieceCommitmentResponse struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	CommP           [32]byte
	NumBytesAligned uint64
	ref4b00fda4     *C.fil_GeneratePieceCommitmentResponse
	allocs4b00fda4  interface{}
}

// FilPoStProof as declared in filecoin-ffi/filecoin.h:104
type FilPoStProof struct {
	RegisteredProof FilRegisteredPoStProof
	ProofLen        uint
	ProofPtr        string
	ref3451bfa      *C.fil_PoStProof
	allocs3451bfa   interface{}
}

// FilGeneratePoStResponse as declared in filecoin-ffi/filecoin.h:111
type FilGeneratePoStResponse struct {
	ErrorMsg       string
	ProofsLen      uint
	ProofsPtr      []FilPoStProof
	StatusCode     FCPResponseStatus
	ref1a2ff84c    *C.fil_GeneratePoStResponse
	allocs1a2ff84c interface{}
}

// FilGpuDeviceResponse as declared in filecoin-ffi/filecoin.h:118
type FilGpuDeviceResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	DevicesLen     uint
	DevicesPtr     []string
	ref58f92915    *C.fil_GpuDeviceResponse
	allocs58f92915 interface{}
}

// FilBLSDigest as declared in filecoin-ffi/filecoin.h:122
type FilBLSDigest struct {
	Inner          [96]byte
	ref215fc78c    *C.fil_BLSDigest
	allocs215fc78c interface{}
}

// FilHashResponse as declared in filecoin-ffi/filecoin.h:129
type FilHashResponse struct {
	Digest         FilBLSDigest
	refc52a22ef    *C.fil_HashResponse
	allocsc52a22ef interface{}
}

// FilBLSPrivateKey as declared in filecoin-ffi/filecoin.h:133
type FilBLSPrivateKey struct {
	Inner          [32]byte
	ref2f77fe3a    *C.fil_BLSPrivateKey
	allocs2f77fe3a interface{}
}

// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filecoin.h:140
type FilPrivateKeyGenerateResponse struct {
	PrivateKey    FilBLSPrivateKey
	ref2dba09f    *C.fil_PrivateKeyGenerateResponse
	allocs2dba09f interface{}
}

// FilBLSPublicKey as declared in filecoin-ffi/filecoin.h:144
type FilBLSPublicKey struct {
	Inner          [48]byte
	ref6d0cab13    *C.fil_BLSPublicKey
	allocs6d0cab13 interface{}
}

// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filecoin.h:151
type FilPrivateKeyPublicKeyResponse struct {
	PublicKey      FilBLSPublicKey
	refee14e59d    *C.fil_PrivateKeyPublicKeyResponse
	allocsee14e59d interface{}
}

// FilPrivateKeySignResponse as declared in filecoin-ffi/filecoin.h:158
type FilPrivateKeySignResponse struct {
	Signature      FilBLSSignature
	refcdf97b28    *C.fil_PrivateKeySignResponse
	allocscdf97b28 interface{}
}

// FilSealCommitPhase1Response as declared in filecoin-ffi/filecoin.h:165
type FilSealCommitPhase1Response struct {
	StatusCode                FCPResponseStatus
	ErrorMsg                  string
	SealCommitPhase1OutputPtr string
	SealCommitPhase1OutputLen uint
	ref61ed8561               *C.fil_SealCommitPhase1Response
	allocs61ed8561            interface{}
}

// FilSealCommitPhase2Response as declared in filecoin-ffi/filecoin.h:172
type FilSealCommitPhase2Response struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ProofPtr       string
	ProofLen       uint
	ref5860b9a4    *C.fil_SealCommitPhase2Response
	allocs5860b9a4 interface{}
}

// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filecoin.h:179
type FilSealPreCommitPhase1Response struct {
	ErrorMsg                     string
	StatusCode                   FCPResponseStatus
	SealPreCommitPhase1OutputPtr string
	SealPreCommitPhase1OutputLen uint
	ref132bbfd8                  *C.fil_SealPreCommitPhase1Response
	allocs132bbfd8               interface{}
}

// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filecoin.h:187
type FilSealPreCommitPhase2Response struct {
	ErrorMsg        string
	StatusCode      FCPResponseStatus
	RegisteredProof FilRegisteredSealProof
	CommD           [32]byte
	CommR           [32]byte
	ref2aa6831d     *C.fil_SealPreCommitPhase2Response
	allocs2aa6831d  interface{}
}

// FilStringResponse as declared in filecoin-ffi/filecoin.h:196
type FilStringResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	StringVal      string
	ref4f413043    *C.fil_StringResponse
	allocs4f413043 interface{}
}

// FilUnsealRangeResponse as declared in filecoin-ffi/filecoin.h:201
type FilUnsealRangeResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref61e219c9    *C.fil_UnsealRangeResponse
	allocs61e219c9 interface{}
}

// FilUnsealResponse as declared in filecoin-ffi/filecoin.h:206
type FilUnsealResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	refdb3aa0f1    *C.fil_UnsealResponse
	allocsdb3aa0f1 interface{}
}

// FilVerifyPoStResponse as declared in filecoin-ffi/filecoin.h:212
type FilVerifyPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	ref3a164861    *C.fil_VerifyPoStResponse
	allocs3a164861 interface{}
}

// FilVerifySealResponse as declared in filecoin-ffi/filecoin.h:218
type FilVerifySealResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refd4397079    *C.fil_VerifySealResponse
	allocsd4397079 interface{}
}

// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filecoin.h:226
type FilWriteWithAlignmentResponse struct {
	CommP                 [32]byte
	ErrorMsg              string
	LeftAlignmentUnpadded uint64
	StatusCode            FCPResponseStatus
	TotalWriteUnpadded    uint64
	refa330e79            *C.fil_WriteWithAlignmentResponse
	allocsa330e79         interface{}
}

// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filecoin.h:233
type FilWriteWithoutAlignmentResponse struct {
	CommP              [32]byte
	ErrorMsg           string
	StatusCode         FCPResponseStatus
	TotalWriteUnpadded uint64
	refc8e1ed8         *C.fil_WriteWithoutAlignmentResponse
	allocsc8e1ed8      interface{}
}

// Fil32ByteArray as declared in filecoin-ffi/filecoin.h:237
type Fil32ByteArray struct {
	Inner          [32]byte
	ref373ec61a    *C.fil_32ByteArray
	allocs373ec61a interface{}
}

// FilPrivateReplicaInfo as declared in filecoin-ffi/filecoin.h:245
type FilPrivateReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CacheDirPath    string
	CommR           [32]byte
	ReplicaPath     string
	SectorId        uint64
	ref81a31e9b     *C.fil_PrivateReplicaInfo
	allocs81a31e9b  interface{}
}

// FilPublicPieceInfo as declared in filecoin-ffi/filecoin.h:250
type FilPublicPieceInfo struct {
	NumBytes       uint64
	CommP          [32]byte
	refd00025ac    *C.fil_PublicPieceInfo
	allocsd00025ac interface{}
}

// FilPublicReplicaInfo as declared in filecoin-ffi/filecoin.h:256
type FilPublicReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CommR           [32]byte
	SectorId        uint64
	ref81b617c2     *C.fil_PublicReplicaInfo
	allocs81b617c2  interface{}
}
