<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rpc/block.proto

namespace GPBMetadata\Rpc;

class Block
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
rpc/block.protoKotFlowBlock"l
VerifyRequest
Height (
	BlockHash (
Hash (
Time (
Nonce (
Miner ("�
TransactionRequest
Value (
BaseFee (

To (
Random (
TxHash (
Time (
Nonce (
Sign ("
HashRequest
Hash ("!
AddressRequest
Address ("
	ReplyBool

Ok ("%

ReplyValue

In (
Out ("
ReplyBalance
Balance ("
HeightRequest
Height ("w

ReplyBlock
Height ($
Header (2.KotFlowBlock.Header3
Transaction (2.KotFlowBlock.ReplyTransaction"+
ReplyLatest
Height (
Hash ("V
Header

ParentHash (
Hash (
Time (
Nonce (
Miner ("�
ReplyTransaction
Value (
BaseFee (

To (
Random (
TxHash (
Time (
Nonce (
Sign (
State	 (2�
GreeterF
BlockHeight.KotFlowBlock.HeightRequest.KotFlowBlock.ReplyBlock" B
Latest.KotFlowBlock.HeightRequest.KotFlowBlock.ReplyLatest" C
CheckBroken.KotFlowBlock.HashRequest.KotFlowBlock.ReplyBool" H
AddressValue.KotFlowBlock.AddressRequest.KotFlowBlock.ReplyValue" L
AddressBalance.KotFlowBlock.AddressRequest.KotFlowBlock.ReplyBalance" M
GetTransaction.KotFlowBlock.HashRequest.KotFlowBlock.ReplyTransaction" @
Verify.KotFlowBlock.VerifyRequest.KotFlowBlock.ReplyBool" J
Transaction .KotFlowBlock.TransactionRequest.KotFlowBlock.ReplyBool" BZ
golang/kotbproto3'
        , true);

        static::$is_initialized = true;
    }
}

